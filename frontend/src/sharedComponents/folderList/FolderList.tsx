import type { IFileTree } from '$models/fileTree/fileTree';
import type { LinkOptions } from '@tanstack/react-router';
import { imageHandler } from '$lib/imageHandler';
import { ImageCard } from '$sharedComponents/card/ImageCard';

interface IFolderListSectionProps {
  folders: IFileTree[];
}

export function FolderList({ folders }: IFolderListSectionProps) {
  return (
    <ol className="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
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
              imageUrls={{
                highQualityImageUrl: imageHandler.getImageUrlForId(
                  child.thumbnailId,
                ),
                lowQualityImageUrl: imageHandler.getImageUrlForId(
                  child.lowQualityThumbnailId,
                ),
              }}
            />
          </li>
        );
      })}
    </ol>
  );
}
