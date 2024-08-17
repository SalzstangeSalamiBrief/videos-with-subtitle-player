import { baseLinkStyles } from '$lib/styles/baseLinkStyles';
import { IFileTreeDto } from '$models/dtos/fileTreeDto';
import { ImageCard } from '$sharedComponents/card/ImageCard';
import { Link as TanStackRouterLink } from '@tanstack/react-router';

interface IFolderListSectionProps {
  selectedFolder: IFileTreeDto;
}

export function FolderListSection({ selectedFolder }: IFolderListSectionProps) {
  return (
    <section role="presentation">
      {selectedFolder.children.length === 0 && (
        <p>This folder contains no subfolders</p>
      )}
      {selectedFolder.children.length > 0 && (
        <ul className="grid gap-4 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
          {selectedFolder.children.map((child) => (
            <li key={child.id}>
              <TanStackRouterLink
                to="/folders/$folderId"
                params={{ folderId: child.id }}
                className={baseLinkStyles}
              >
                <ImageCard title={child.name} imageUrl="/example.jpg" />
              </TanStackRouterLink>
            </li>
          ))}
        </ul>
      )}
    </section>
  );
}
