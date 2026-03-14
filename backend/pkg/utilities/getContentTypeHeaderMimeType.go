package utilities

import (
	"backend/pkg/models"
	"mime"
	"path/filepath"
)

func GetContentTypeHeaderMimeType(file models.FileNode) string {
	ext := filepath.Ext(file.Path)
	mimeType := mime.TypeByExtension(ext)
	if ext == ".vtt" {
		mimeType = "text/vtt"
	}

	return mimeType
}
