import { IFileNode } from '$models/fileTree';
import { ImageSlider } from '$sharedComponents/imageSlider/ImageSlider';

interface IImageListSectionProps {
  images: IFileNode[];
}

export function ImageListSection({ images }: IImageListSectionProps) {
  if (!images.length) {
    return <p>This folder contains no images</p>;
  }

  // TODO MAYBE ADD EVENT_LISTENER FOR KEYBOARD USAGE
  return <ImageSlider images={images} />;
}
