package database

import (
	"backend/pkg/models"
	"context"
	"log"
	"slices"

	"gorm.io/gorm"
)

// TODO SYNC TREE =>
// TODO CAN REUSE FOR FILES
// TODO HAVE TO ADD FOR FOLDERS
// TODO MIGRATION
// TODO ADJUST TABLE NAMES

func syncFolders(databaseConnection *gorm.DB, ctx context.Context, folderTreesFromDisk []models.FolderNode, flatHierarchyOfFoldersFromDatabase []models.FolderNode) error {
	flatHierarchyOfFoldersFromDisk := models.NodesToFlatHierarchy(folderTreesFromDisk)
	foldersToCreate := getDistinctFolders(flatHierarchyOfFoldersFromDisk, flatHierarchyOfFoldersFromDatabase)
	foldersToDelete := getDistinctFolders(flatHierarchyOfFoldersFromDatabase, flatHierarchyOfFoldersFromDisk)
	remainingFoldersToSyncRecursively := getRemainingFolders(flatHierarchyOfFoldersFromDisk, foldersToCreate, foldersToDelete)

	deleteFoldersError := deleteFoldersBatchFromDb(databaseConnection, ctx, foldersToDelete)
	if deleteFoldersError != nil {
		return deleteFoldersError
	}

	createFoldersError := createFoldersBatch(databaseConnection, ctx, foldersToCreate)
	if createFoldersError != nil {
		return createFoldersError
	}

	// TODO SYNC FOLDERS/FILES
	// TODO AFTER LEAVING THE FUNCTION THE DB CRASHES/IS FLOODED
	log.Println(remainingFoldersToSyncRecursively)
	return nil
}

func getDistinctFolders(left []models.FolderNode, right []models.FolderNode) []models.FolderNode {
	distinctFolders := make([]models.FolderNode, 0)
	for _, leftItem := range left {
		isItemInBoothSets := false
		for _, rightItem := range right {
			// TODO DOES not work with the nested structure => WOULD DELETE ALL CHILDREND OF A PARENT
			if rightItem.Path == leftItem.Path {
				isItemInBoothSets = true
				continue
			}
		}

		if !isItemInBoothSets {
			distinctFolders = append(distinctFolders, leftItem)
		}
	}

	return distinctFolders
}

func mapFolderSliceToSliceOfPaths(folders []models.FolderNode) []string {
	paths := make([]string, len(folders))
	for i, folder := range folders {
		paths[i] = folder.Path
	}

	return paths
}

func getRemainingFolders(foldersFromDisk []models.FolderNode, createdFolders []models.FolderNode, deletedFolders []models.FolderNode) []models.FolderNode {
	manipulatedFolders := slices.Concat(createdFolders, deletedFolders)
	//manipulatedPaths := mapFolderSliceToSliceOfPaths(manipulatedFolders)
	unchangedFolders := getDistinctFolders(foldersFromDisk, manipulatedFolders)
	return unchangedFolders
}

func createFoldersBatch(databaseConnection *gorm.DB, ctx context.Context, foldersToCreate []models.FolderNode) error {
	result := gorm.WithResult()
	createFoldersInBatchError := gorm.G[models.FolderNode](databaseConnection, result).CreateInBatches(ctx, &foldersToCreate, len(foldersToCreate))
	if createFoldersInBatchError != nil {
		log.Println("Error creating folders")
		return createFoldersInBatchError
	}

	return nil
}

func deleteFoldersBatchFromDb(databaseConnection *gorm.DB, ctx context.Context, foldersToDelete []models.FolderNode) error {
	for _, folderToDelete := range foldersToDelete {
		_, err := gorm.G[models.FolderNode](databaseConnection).Where("path = ?", folderToDelete.Path).Delete(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func tryGetFolderNodeByPath(databaseConnection *gorm.DB, ctx context.Context, path string) (*models.FileNode, error) {
	matchingFile, err := gorm.G[*models.FileNode](databaseConnection).Where("path = ?", path).First(ctx)
	return matchingFile, err
}
