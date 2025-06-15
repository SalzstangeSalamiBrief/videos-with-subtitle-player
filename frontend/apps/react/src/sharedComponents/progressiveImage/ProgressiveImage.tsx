import type { Maybe } from '@videos-with-subtitle-player/core';
import { useState } from 'react';
import styles from './ProgressiveImage.module.css';

export interface IProgressiveImageProps {
  highQualityImageUrl: Maybe<string>;
  lowQualityImageUrl: Maybe<string>;
  alt: string;
}

export function ProgressiveImage({
  highQualityImageUrl,
  lowQualityImageUrl,
  alt,
}: IProgressiveImageProps) {
  // TODO INTERSECTION OBSERVER for performance => one observer for all
  const [isHighQualityImageLoaded, setIsHighQualityImageLoaded] =
    useState<boolean>(false);

  if (!highQualityImageUrl) {
    return null;
  }

  return (
    <div className={styles['progressive-image-container']}>
      {lowQualityImageUrl && (
        <img
          loading="eager"
          src={lowQualityImageUrl}
          className={`${styles['progressive-image']} ${styles['low-quality']} ${isHighQualityImageLoaded ? styles['hidden'] : ''}`}
          alt={alt}
        />
      )}
      {highQualityImageUrl && (
        <img
          loading="lazy"
          src={highQualityImageUrl}
          className={`${styles['progressive-image']} ${styles['high-quality']} ${isHighQualityImageLoaded ? '' : styles['hidden']}`}
          alt={alt}
          onLoad={() => setIsHighQualityImageLoaded(true)}
        />
      )}
    </div>
  );
}
