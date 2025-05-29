import { getImageUrlForId } from '$lib/utilities/getImageUrl';
import type { IFileNode } from '$models/fileTree';
import styles from './ImageSlider.module.css';

interface IFolderListSectionProps {
  images: IFileNode[];
  onImageClick?: (image: IFileNode) => void;
}

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
            <img src={getImageUrlForId(image.id)} alt={image.name} />
          </figure>
        ))}
      </div>
    </div>
  );
}
