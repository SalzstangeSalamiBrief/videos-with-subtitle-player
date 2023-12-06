package models

type DirectoryTreeItem struct {
	Path         string
	Name         string
	Children     []DirectoryTreeItem
	AudioFile    FileTreeItem
	SubtitleFile FileTreeItem
}
