import { Maybe } from '@videos-with-subtitle-player/core';

export interface IProgressiveImageProps {
  highQualityImageUrl: Maybe<string>;
  lowQualityImageUrl: Maybe<string>;
  alt: string;
}
