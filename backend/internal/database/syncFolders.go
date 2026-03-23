package database

import (
	"backend/pkg/models"
	"context"
	"log"
	"slices"

	"gorm.io/gorm"
)

func syncFolders(databaseConnection *gorm.DB, ctx context.Context, folderTreesFromDisk []models.FolderNode, flatHierarchyOfFoldersFromDatabase []models.FolderNode) error {
	flatHierarchyOfFoldersFromDisk := models.NodesToFlatHierarchy(folderTreesFromDisk)
	foldersToCreate := getDistinctFolders(flatHierarchyOfFoldersFromDisk, flatHierarchyOfFoldersFromDatabase)
	foldersToDelete := getDistinctFolders(flatHierarchyOfFoldersFromDatabase, flatHierarchyOfFoldersFromDisk)
	foldersToUpdate := getFoldersToUpdate(flatHierarchyOfFoldersFromDisk, foldersToCreate, foldersToDelete)

	deleteFoldersError := deleteFoldersBatchFromDb(databaseConnection, ctx, foldersToDelete)
	if deleteFoldersError != nil {
		return deleteFoldersError
	}

	createFoldersError := createFoldersBatch(databaseConnection, ctx, foldersToCreate)
	if createFoldersError != nil {
		return createFoldersError
	}
	// TODO THIS DOES NOT WORK PROPERLY
	// 1. USE THE DEFAULT DIRECTORY with n unique entries
	// 2. USE ANOTHER DIRECTORY with m unique entries
	// EXPECTED: n entries delete, m entries added
	// CURRENT: m-n entries added
	syncFoldersRecursivelyError := syncFoldersRecursively(databaseConnection, ctx, foldersToUpdate)
	if syncFoldersRecursivelyError != nil {
		return syncFoldersRecursivelyError
	}

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

func getFoldersToUpdate(foldersFromDisk []models.FolderNode, createdFolders []models.FolderNode, deletedFolders []models.FolderNode) []models.FolderNode {
	manipulatedFolders := slices.Concat(createdFolders, deletedFolders)
	//manipulatedPaths := mapFolderSliceToSliceOfPaths(manipulatedFolders)
	unchangedFolders := getDistinctFolders(foldersFromDisk, manipulatedFolders)
	return unchangedFolders
}

func createFoldersBatch(databaseConnection *gorm.DB, ctx context.Context, foldersToCreate []models.FolderNode) error {
	result := gorm.WithResult()

	x := make([]string, len(foldersToCreate))
	for i, folder := range foldersToCreate {
		x[i] = folder.Path
	}

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

func syncFoldersRecursively(databaseConnection *gorm.DB, ctx context.Context, foldersFromDisk []models.FolderNode) error {
	for _, folderFromDisk := range foldersFromDisk {
		folderToSync, getFolderFromDbError := tryGetFolderNodeByPath(databaseConnection, ctx, folderFromDisk.Path)
		if getFolderFromDbError != nil {
			return getFolderFromDbError
		}

		if folderToSync == nil {
			log.Printf("Skipped syncing folder with path='%s' because there is no match in the database", folderFromDisk.Path)
			continue
		}

		syncFolderError := syncFolderRecursively(databaseConnection, ctx, *folderToSync, folderFromDisk)
		if syncFolderError != nil {
			return syncFolderError
		}
	}

	return nil
}

func syncFolderRecursively(databaseConnection *gorm.DB, ctx context.Context, folderFromDb models.FolderNode, folderFromDisk models.FolderNode) error {
	shouldSync := hasFolderChanges(folderFromDb, folderFromDisk)
	if !shouldSync {
		return nil
	}

	syncFolderError := updateFolder(databaseConnection, ctx, folderFromDb, folderFromDisk)
	if syncFolderError != nil {
		return syncFolderError
	}

	//// TODO CHECK IF FILES HAVE THE CORRECT PARENT_FOLDER_ID#
	filesInFolder := folderFromDisk.Files
	for i := range filesInFolder {
		filesInFolder[i].ParentFolderId = folderFromDb.FolderId
	}

	syncFilesError := syncFiles(databaseConnection, ctx, folderFromDb.Files, filesInFolder)
	if syncFilesError != nil {
		return syncFilesError
	}

	return nil
}

func hasFolderChanges(folderFromDb models.FolderNode, folderFromDisk models.FolderNode) bool {
	hasNameChanged := folderFromDb.Name != folderFromDisk.Name
	hasParentFolderChanged := folderFromDb.ParentFolderId != folderFromDisk.ParentFolderId
	hasThumbnailChanged := folderFromDb.ThumbnailId != folderFromDisk.ThumbnailId
	hasLowQualityThumbnailChanged := folderFromDb.LowQualityThumbnailId != folderFromDisk.LowQualityThumbnailId
	hasNumberOfFilesChanged := len(folderFromDb.Files) != len(folderFromDisk.Files)
	return hasNameChanged || hasParentFolderChanged || hasThumbnailChanged || hasLowQualityThumbnailChanged || hasNumberOfFilesChanged
}

func updateFolder(databaseConnection *gorm.DB, ctx context.Context, folderFromDb models.FolderNode, folderFromDisk models.FolderNode) error {
	folderFromDb.Name = folderFromDisk.Name
	folderFromDb.ParentFolderId = folderFromDisk.ParentFolderId
	folderFromDb.ThumbnailId = folderFromDisk.ThumbnailId
	folderFromDb.LowQualityThumbnailId = folderFromDisk.LowQualityThumbnailId
	_, err := gorm.G[models.FolderNode](databaseConnection).Where("id = ?", folderFromDb.ID).Updates(ctx, folderFromDb)
	return err
}

func tryGetFolderNodeByPath(databaseConnection *gorm.DB, ctx context.Context, path string) (*models.FolderNode, error) {
	matchingFile, err := gorm.G[*models.FolderNode](databaseConnection).Where("path = ?", path).First(ctx)
	return matchingFile, err
}
