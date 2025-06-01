import type { IFileNode } from '@videos-with-subtitle-player/core';
import { useId } from 'react';
import { Route } from '../../../routes/folders/_folderLayout/$folderId/files/$fileId';
import { PlaylistItem } from './PlaylistItem';

interface IPlaylistProps {
  siblings: IFileNode[];
}

export function Playlist({ siblings }: IPlaylistProps) {
  const labelId = useId();
  const { folderId, fileId } = Route.useParams();

  return (
    <section aria-labelledby={labelId}>
      <h2 id={labelId} className="mb-2 text-base font-bold">
        Playlist
      </h2>
      <ol className="flex h-80 flex-col gap-1 overflow-x-hidden overflow-y-auto">
        {siblings.map((item) => (
          <PlaylistItem
            key={item.id}
            isActive={item.id === fileId}
            item={item}
            folderId={folderId}
          />
        ))}
      </ol>
    </section>
  );
}
