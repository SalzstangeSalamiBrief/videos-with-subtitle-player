import { IFileNode } from '$models/fileTree';
import { Lightbox } from '$sharedComponents/lightbox/Lightbox';

interface IImageListSectionProps {
  images: IFileNode[];
}

export function ImageListSection({ images }: IImageListSectionProps) {
  if (!images.length) {
    return <p>This folder contains no images</p>;
  }

  // TODO MAYBE ADD EVENT_LISTENER FOR KEYBOARD USAGE
  return <Lightbox images={images} />;
}
