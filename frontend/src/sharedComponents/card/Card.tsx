interface ICardProps {
  title: string;
  children?: JSX.Element;
  imageUrl: string;
}

export function Card({ title, imageUrl }: ICardProps) {
  return (
    <article
      className="grid gap-4 h-64 rounded-md bg-cover bg-no-repeat"
      style={{
        backgroundImage: `url(${imageUrl})`,
      }}
    >
      <div
        role="presentation"
        className="rounded-md min-w-0 p-4 flex items-end bg-gradient-to-t from-slate-800 from-[5ch] to-[20ch]"
      >
        <header className="overflow-hidden h-fit rounded-b-md">
          <h1
            className="line-clamp-3 whitespace-normal text-ellipsis overflow-hidden font-bold"
            title={title}
          >
            {title}
          </h1>
        </header>
      </div>
    </article>
  );
}
