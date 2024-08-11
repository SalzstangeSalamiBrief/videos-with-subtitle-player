import { FileTreeContext } from '$contexts/FileTreeContextWrapper';
import { getFoldersInActiveTree } from '$lib/utilities/getFoldersInActiveTree';
import { useParams } from '@tanstack/react-router';
import { useContext } from 'react';
import { BreadcrumbItem } from './BreadcrumbItem';
import { Link as TanStackRouterLink } from '@tanstack/react-router';
import { baseLinkStyles } from '$lib/styles/baseLinkStyles';

export function Breadcrumbs() {
  const { folderId } = useParams({ strict: false });
  const { fileTrees } = useContext(FileTreeContext);
  const activeFolders = getFoldersInActiveTree(fileTrees, folderId);

  if (!activeFolders.length) {
    return null;
  }

  return (
    <nav>
      <menu className="flex gap-2">
        <li>
          <TanStackRouterLink to="/" className={baseLinkStyles}>
            Home
          </TanStackRouterLink>
        </li>
        {activeFolders.map((activeFolder, index) => (
          <BreadcrumbItem
            key={activeFolder.id}
            item={activeFolder}
            isLastItem={index === activeFolders.length - 1}
          />
        ))}
      </menu>
    </nav>
  );
}
