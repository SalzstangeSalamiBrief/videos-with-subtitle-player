import { getImageUrlForId } from '$lib/utilities/getImageUrl';
import type { IFileTree } from '$models/fileTree';
import { ImageCard } from '$sharedComponents/card/ImageCard';
import type { LinkOptions } from '@tanstack/react-router';

interface IFolderListSectionProps {
  folders: IFileTree[];
}

export function FolderList({ folders }: IFolderListSectionProps) {
  return (
    <ul className="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
      {folders.map((child) => {
        const linkOption: LinkOptions = {
          to: '/folders/$folderId',
          params: { folderId: child.id },
        };
        return (
          <li key={child.id}>
            <ImageCard
              linkOptions={linkOption}
              title={child.name}
              imageUrl={getImageUrlForId(child.thumbnailId)}
            />
          </li>
        );
      })}
    </ul>
  );
}
