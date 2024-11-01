import { baseLinkStyles } from '$lib/styles/baseLinkStyles';
import {
  LinkOptions,
  Link as TanStackRouterLink,
} from '@tanstack/react-router';
import styles from './ImageCard.module.css';
interface ICardProps {
  title: string;
  children?: JSX.Element;
  imageUrl?: string;
  linkOptions: LinkOptions;
}

export function ImageCard({ title, imageUrl, linkOptions }: ICardProps) {
  return (
    <article
      className={`${styles.imageCard} grid h-64 gap-4 rounded-md bg-cover bg-no-repeat ${imageUrl ? undefined : 'bg-fuchsia-800'}`}
      style={{
        backgroundImage: imageUrl ? `url(${imageUrl})` : undefined,
      }}
    >
      <div
        role="presentation"
        className="flex min-w-0 items-end rounded-md bg-gradient-to-t from-slate-800 from-[5ch] to-[20ch] p-4"
      >
        <header className="h-fit overflow-hidden rounded-b-md">
          <h1
            className="line-clamp-3 max-w-text overflow-hidden text-ellipsis whitespace-normal font-bold"
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
    </article>
  );
}
