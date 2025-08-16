package imageHandler

type ImageHandler interface {
	ReduceImageQuality(sourceImagePath string) (lowQualityImagePath string, err error)
	IsLowQualityFile(sourcePath string) bool
}
