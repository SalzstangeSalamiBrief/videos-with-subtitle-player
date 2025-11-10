import { baseLinkStyles } from '$lib/styles/baseLinkStyles';
import {
  Link as TanStackRouterLink,
  type LinkOptions,
} from '@tanstack/react-router';
import type { JSX } from 'react';
import {
  ProgressiveImage,
  type IProgressiveImageProps,
} from '../progressiveImage/ProgressiveImage';
import styles from './ImageCard.module.css';

interface ICardProps {
  title: string;
  children?: JSX.Element;
  imageUrls?: Omit<IProgressiveImageProps, 'alt'>;
  linkOptions: LinkOptions;
}

export function ImageCard({ title, imageUrls, linkOptions }: ICardProps) {
  // TODO DOES NOT USE THE WHOLE HEIGHT
  return (
    <>
      <article className={`card image-full ${styles.card}`}>
        {imageUrls && (
          <ProgressiveImage
            {...imageUrls}
            alt={`Cover image of the item ${title}`}
          />
        )}
        <div
          role="presentation"
          className={`card-body ${styles['card-body']} bg-linear-to-t from-slate-800 from-[5ch] to-[20ch] p-2`}
        >
          <header className="card-title">
            <h1
              title={title}
              className="max-w-text line-clamp-3 overflow-hidden font-bold text-ellipsis whitespace-normal"
            >
              <TanStackRouterLink
                {...linkOptions}
                className={`${baseLinkStyles} block h-full w-full`}
              >
                {title}
              </TanStackRouterLink>
            </h1>
          </header>
        </div>
        {/* <div className="card-body">
          <h2 className="card-title">{title}</h2>
          <p>
            A card component has a figure, a body part, and inside body there
            are title and actions parts
          </p>
          <div className="card-actions justify-end">
            <button className="btn btn-primary">Buy Now</button>
          </div>
        </div> */}
      </article>
      {/* <article
        className={`${styles.card} ${imageUrls ? '' : 'bg-fuchsia-800'}`}
      >
        {imageUrls && (
          <ProgressiveImage
            {...imageUrls}
            alt={`Cover image of the item ${title}`}
          />
        )}
        <div role="presentation">
          <div
            role="presentation"
            className="flex h-full items-end bg-linear-to-t from-slate-800 from-[5ch] to-[20ch]"
          >
            <header className="h-fit overflow-hidden rounded-b-md p-4">
              <h1
                className="max-w-text line-clamp-3 overflow-hidden font-bold text-ellipsis whitespace-normal"
                title={title}
              >
                <TanStackRouterLink
                  {...linkOptions}
                  className={`${baseLinkStyles} block h-full w-full`}
                >
                  {title}
                </TanStackRouterLink>
              </h1>
            </header>
          </div>
        </div>
      </article> */}
    </>
  );
}
