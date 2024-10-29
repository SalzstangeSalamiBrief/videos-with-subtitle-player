interface ICardProps {
  title: string;
  children?: JSX.Element;
  imageUrl?: string;
}

export function ImageCard({ title, imageUrl }: ICardProps) {
  // TODO ON HOVER rotate the card a bit on the z-axis
  return (
    <article
      className={`bg-no-repeate grid h-64 gap-4 rounded-md bg-cover ${imageUrl ? undefined : 'bg-fuchsia-800'}`}
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
            {title}
          </h1>
        </header>
      </div>
    </article>
  );
}
