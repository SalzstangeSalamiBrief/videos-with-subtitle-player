package fileTreeManager

import (
	"backend/pkg/models"
	"errors"
	"log"
)

type SubFileTree []models.FileNode

type FileTreeManager struct {
	rootPath  string
	fileNodes *[]models.FileNode
	tree      *models.FolderNode
}

func NewFileTreeManager(rootPath string) *FileTreeManager {
	return &FileTreeManager{
		rootPath:  rootPath,
		fileNodes: &[]models.FileNode{},
	}
}

func (fileTreeManager *FileTreeManager) InitializeFiles() *FileTreeManager {
	log.Default().Println("Start file tree initialization")
	files := fileTreeManager.scanForFilesInPath(fileTreeManager.rootPath)
	fileTreeManager.fileNodes = &files
	log.Default().Println("Finish file tree initialization")
	return fileTreeManager
}

func (fileTreeManager *FileTreeManager) InitializeFileNodesTree() *FileTreeManager {
	log.Default().Println("Start file tree initialization")
	fullTree := fileTreeManager.convertFileNodesToTree()
	fileTreeManager.tree = &fullTree
	log.Default().Println("Finish file tree initialization")
	return fileTreeManager
}

func (fileTreeManager *FileTreeManager) GetFiles() []models.FileNode {
	if fileTreeManager.fileNodes == nil {
		log.Fatal(errors.New("file nodes are not initialized"))
	}

	return *fileTreeManager.fileNodes
}

func (fileTreeManager *FileTreeManager) GetRootTree() models.FolderNode {
	if fileTreeManager.tree == nil {
		log.Fatal(errors.New("file tree is not initialized"))
	}

	return *fileTreeManager.tree
}

func (fileTreeManager *FileTreeManager) GetSubTrees() []models.FolderNode {
	if fileTreeManager.tree == nil {
		log.Fatal(errors.New("file tree is not initialized"))
	}

	return fileTreeManager.tree.ChildFolders
}
