import { getImageUrlForId } from '$lib/utilities/getImageUrl';
import { IFileNode } from '$models/fileTree';
import styles from './ImageSlider.module.css';

interface IFolderListSectionProps {
  images: IFileNode[];
}

export function ImageSlider({ images }: IFolderListSectionProps) {
  return (
    <div className={styles.slideContainer}>
      <div className={styles.slideShow}>
        {images.map((image) => (
          <figure className={styles.slide} key={image.id}>
            <img src={getImageUrlForId(image.id)} alt={image.name} />
          </figure>
        ))}
      </div>
    </div>
  );
}
