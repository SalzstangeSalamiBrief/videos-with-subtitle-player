import { ImageCard } from '$sharedComponents/card/ImageCard';
import type { LinkOptions } from '@tanstack/react-router';
import type { IFileNode } from '@videos-with-subtitle-player/core';

interface IFileListSectionProps {
  folderId: string;
  files: IFileNode[];
}

export function FileListSection({ files, folderId }: IFileListSectionProps) {
  return (
    <section role="presentation">
      {files.length === 0 && <p>This folder contains no files</p>}
      {files.length > 0 && (
        <ol className="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
          {files.map((file) => {
            const linkOption: LinkOptions = {
              to: '/folders/$folderId/files/$fileId',
              params: { folderId, fileId: file.id },
            };
            return (
              <li key={file.id}>
                <ImageCard
                  title={file.name}
                  imageUrls={undefined}
                  linkOptions={linkOption}
                />
              </li>
            );
          })}
        </ol>
      )}
    </section>
  );
}
