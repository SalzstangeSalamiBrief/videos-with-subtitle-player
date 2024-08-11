import { FileTreeContext } from '$contexts/FileTreeContextWrapper';
import { getFoldersInActiveTree } from '$lib/getFoldersInActiveTree';
import { useParams } from '@tanstack/react-router';
import { useContext } from 'react';
import { Link as TanStackRouterLink } from '@tanstack/react-router';

export function FolderBreadcrumbs() {
  const { folderId } = useParams({ strict: false });
  const { fileTrees } = useContext(FileTreeContext);
  const activeFolders = getFoldersInActiveTree(fileTrees, folderId);

  if (!activeFolders.length) {
    return null;
  }

  return (
    <nav>
      <menu className="flex gap-2">
        {activeFolders.map((activeFolder) => (
          <TanStackRouterLink
            to="/folders/$folderId"
            params={{ folderId: activeFolder.id }}
            key={activeFolder.id}
            className="block max-w-[30ch] overflow-x-hidden text-ellipsis whitespace-nowrap"
            title={activeFolder.name}
          >
            / {activeFolder.name}
          </TanStackRouterLink>
        ))}
      </menu>
    </nav>
  );
}
