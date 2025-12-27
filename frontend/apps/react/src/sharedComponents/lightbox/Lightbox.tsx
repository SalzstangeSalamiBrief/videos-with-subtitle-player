import { imageHandler } from '$lib/styles/imageHandler';
import { ImageSlider } from '$sharedComponents/imageSlider/ImageSlider';
import { XMarkIcon } from '@heroicons/react/24/outline';
import { type IFileNode, type Maybe } from '@videos-with-subtitle-player/core';
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
      <dialog ref={dialogRef} className="modal">
        <div className="modal-backdrop" />
        <div className={`modal-box ${styles.lightbox}`}>
          {!activeImage && <p>Please select an image</p>}
          {activeImage && (
            <figure>
              <img
                src={imageHandler.getImageUrlForId(activeImage.id)}
                alt={activeImage.name}
              />
            </figure>
          )}
          <form method="dialog">
            <button
              className="btn btn-sm btn-circle btn-ghost absolute top-2 right-2 hover:bg-fuchsia-800"
              onClick={() => dialogRef.current?.close()}
              aria-label="Close the light box"
            >
              <XMarkIcon />
            </button>
          </form>
        </div>
      </dialog>
    </>
  );
}
