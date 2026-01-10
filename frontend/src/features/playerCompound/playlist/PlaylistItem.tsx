import type { IFileNode } from '$models/fileTree/fileTree';
import { Link as TanStackLink } from '@tanstack/react-router';

interface IPlaylistItemProps {
  item: IFileNode;
  isActive: boolean;
  folderId: string;
}

export function PlaylistItem({ isActive, item, folderId }: IPlaylistItemProps) {
  return (
    <li
      className={`w-text h-16 shrink-0 rounded-md px-4 py-2 ${isActive ? 'bg-fuchsia-800 hover:bg-fuchsia-700' : 'bg-slate-800 hover:bg-slate-700'}`}
    >
      <TanStackLink
        className="clamp-container clamp-2"
        to="/folders/$folderId/files/$fileId"
        params={{ fileId: item.id, folderId }}
        aria-label={`Play ${item.name}`}
        title={item.name}
        aria-selected={isActive ? 'true' : 'false'}
      >
        {item.name}
      </TanStackLink>
    </li>
  );
}
