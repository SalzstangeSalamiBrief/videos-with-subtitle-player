import { baseLinkStyles } from '$lib/styles/baseLinkStyles';
import { IFileTreeDto } from '$models/dtos/fileTreeDto';
import { ImageCard } from '$sharedComponents/card/ImageCard';
import { Link as TanStackRouterLink } from '@tanstack/react-router';

interface IFolderListSectionProps {
  folders: IFileTreeDto[];
}

export function FolderListSection({ folders }: IFolderListSectionProps) {
  if (!folders.length) {
    return <p>This folder contains no subfolders</p>;
  }

  return (
    <ul className="grid gap-4 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
      {folders.map((child) => (
        <li key={child.id}>
          <TanStackRouterLink
            to="/folders/$folderId"
            params={{ folderId: child.id }}
            className={baseLinkStyles}
          >
            <ImageCard title={child.name} imageUrl="/example.avif" />
          </TanStackRouterLink>
        </li>
      ))}
    </ul>
  );
}
