import { getImageUrlForId } from '$lib/utilities/getImageUrl';
import type { IFileNode } from '$models/fileTree';
import { ImageSlider } from '$sharedComponents/imageSlider/ImageSlider';
import { XMarkIcon } from '@heroicons/react/24/outline';
import { useRef, useState } from 'react';
import styles from './Lightbox.module.css';

interface ILightboxContainerProps {
  images: IFileNode[];
}

export function Lightbox({ images }: ILightboxContainerProps) {
  const dialogRef = useRef<HTMLDialogElement>(null);
  const [activeImage, setActiveImage] = useState<Maybe<IFileNode>>(undefined);

  return (
    <>
      <ImageSlider
        images={images}
        onImageClick={(selectedImage: IFileNode) => {
          setActiveImage(selectedImage);
          dialogRef.current?.showModal();
        }}
      />
      <dialog ref={dialogRef}>
        <div className={styles.lightbox}>
          {!activeImage && <p>Please select an image</p>}
          {activeImage && (
            <figure>
              <img
                src={getImageUrlForId(activeImage.id)}
                alt={activeImage.name}
              />
            </figure>
          )}
          <button
            className="absolute right-0 top-0 w-10 bg-fuchsia-800 p-2 text-fuchsia-100 hover:bg-fuchsia-700"
            onClick={() => dialogRef.current?.close()}
            aria-label="Close the light box"
          >
            <XMarkIcon />
          </button>
        </div>
      </dialog>
    </>
  );
}
