package fileTreeManager

import (
	"backend/pkg/models"
	"log"
)

type SubFileTree []models.FileNode

type FileTreeManager struct {
	rootPath  string
	fileNodes []models.FileNode
	tree      models.FolderNode
}

func NewFileTreeManager(rootPath string) *FileTreeManager {
	return &FileTreeManager{
		rootPath:  rootPath,
		fileNodes: []models.FileNode{},
	}
}

func (fileTreeManager *FileTreeManager) InitializeFiles() *FileTreeManager {
	log.Default().Println("Start file tree initialization")
	fullTree := fileTreeManager.scanForFilesInPath(fileTreeManager.rootPath)
	fileTreeManager.fileNodes = fullTree
	log.Default().Println("Finish file tree initialization")
	return fileTreeManager
}

func (fileTreeManager *FileTreeManager) ConvertFileNodesToTree() *FileTreeManager {
	log.Default().Println("Start file tree initialization")
	fullTree := fileTreeManager.convertFileNodesToTree()
	fileTreeManager.tree = fullTree
	log.Default().Println("Finish file tree initialization")
	return fileTreeManager
}

func (fileTreeManager *FileTreeManager) GetFiles() []models.FileNode {
	return fileTreeManager.fileNodes
}

func (fileTreeManager *FileTreeManager) GetTree() models.FolderNode {
	return fileTreeManager.tree
}
