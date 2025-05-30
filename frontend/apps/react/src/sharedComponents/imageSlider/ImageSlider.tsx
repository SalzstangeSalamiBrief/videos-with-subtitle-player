import {
  type IFileNode,
  getImageUrlForId,
} from '@videos-with-subtitle-player/core';
import styles from './ImageSlider.module.css';

interface IFolderListSectionProps {
  images: IFileNode[];
  onImageClick?: (image: IFileNode) => void;
}
const baseUrl = import.meta.env.VITE_BASE_URL;
export function ImageSlider({ images, onImageClick }: IFolderListSectionProps) {
  return (
    <div className={styles.slideContainer}>
      <div className={styles.slideShow}>
        {images.map((image) => (
          <figure
            className={`${styles.slide} ${onImageClick ? 'cursor-pointer' : undefined}`}
            key={image.id}
            onClick={() => onImageClick?.(image)}
          >
            <img src={getImageUrlForId(baseUrl, image.id)} alt={image.name} />
          </figure>
        ))}
      </div>
    </div>
  );
}
