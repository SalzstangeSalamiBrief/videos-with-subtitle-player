import { baseLinkStyles } from '$lib/styles/baseLinkStyles';
import { getImageUrlForId } from '$lib/utilities/getImageUrl';
import { IFileTree } from '$models/fileTree';
import { ImageCard } from '$sharedComponents/card/ImageCard';
import { Link as TanStackRouterLink } from '@tanstack/react-router';

interface IFolderListSectionProps {
  folders: IFileTree[];
}

export function FolderList({ folders }: IFolderListSectionProps) {
  return (
    <ul className="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
      {folders.map((child) => (
        <li key={child.id}>
          <TanStackRouterLink
            to="/folders/$folderId"
            params={{ folderId: child.id }}
            className={baseLinkStyles}
          >
            <ImageCard
              title={child.name}
              imageUrl={getImageUrlForId(child.thumbnailId)}
            />
          </TanStackRouterLink>
        </li>
      ))}
    </ul>
  );
}
