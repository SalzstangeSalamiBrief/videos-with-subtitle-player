package models

import (
	"backend/pkg/enums"
	"mime"
	"path/filepath"
	"slices"
	"strings"
)

type FileTreeNode struct {
	Id                    string
	Path                  string
	Name                  string
	Type                  enums.FileType
	AssociatedAudioFileId string
}

func (file *FileTreeNode) GetPartsOfPath() []string {
	filePath, _ := filepath.Split(file.Path)
	allParts := strings.Split(filePath, string(filepath.Separator))
	var parts []string
	for _, part := range allParts {
		if part == "" {
			continue
		}

		parts = append(parts, part)
	}

	return parts
}

func (file *FileTreeNode) IsFileExtensionAllowed(allowedExtension ...string) bool {
	ext := filepath.Ext(file.Path)
	doesExtensionMatch := slices.Contains(allowedExtension, ext)
	return doesExtensionMatch
}

func (file *FileTreeNode) GetMimeType() string {
	ext := filepath.Ext(file.Path)
	mimeType := mime.TypeByExtension(ext)
	if ext == ".vtt" {
		mimeType = "text/vtt"
	}

	return mimeType
}
