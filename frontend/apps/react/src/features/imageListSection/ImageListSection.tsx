import { Lightbox } from '$sharedComponents/lightbox/Lightbox';
import type { IFileNode } from '@videos-with-subtitle-player/core';

interface IImageListSectionProps {
  images: IFileNode[];
}

export function ImageListSection({ images }: IImageListSectionProps) {
  const highQualityImages = images.filter(
    (image) => !image.name.includes('_lowQuality'),
  );

  if (!highQualityImages.length) {
    return <p>This folder contains no images</p>;
  }

  // TODO MAYBE ADD EVENT_LISTENER FOR KEYBOARD USAGE
  return <Lightbox images={highQualityImages} />;
}
