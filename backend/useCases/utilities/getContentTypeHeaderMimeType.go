package usecases

import (
	"backend/models"
	"mime"
	"path/filepath"
)

func GetContentTypeHeaderMimeType(file models.FileTreeItem) string {
	ext := filepath.Ext(file.Path)
	mimeType := mime.TypeByExtension(ext)
	if ext == ".vtt" {
		mimeType = "text/vtt"
	}

	return mimeType
}
