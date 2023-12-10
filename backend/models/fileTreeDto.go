package models

// TODO maybe swagger??
type FileTreeDto struct {
	Name       string         `json:"name"`
	Id         string         `json:"id"`
	AudioFiles []AudioFileDto `json:"audioFiles"`
	Children   []FileTreeDto  `json:"children"`
}

type AudioFileDto struct {
	Name         string      `json:"name"`
	SubtitleFile FileItemDto `json:"subtitleFile"`
	AudioFile    FileItemDto `json:"audioFile"`
}

type FileItemDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
