import { Lightbox } from '$sharedComponents/lightbox/Lightbox';
import type { IFileNode } from '@videos-with-subtitle-player/core';

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
