import { baseLinkStyles } from '$lib/styles/baseLinkStyles';
import { Link as TanStackRouterLink, useParams } from '@tanstack/react-router';
import { getFoldersInActiveTree } from '@videos-with-subtitle-player/core';
import { Route as RootLayoutRoute } from '../../routes/__root';
import { BreadcrumbItem } from './BreadcrumbItem';

export function Breadcrumbs() {
  const { folderId, fileId } = useParams({ strict: false });
  const { fileTrees } = RootLayoutRoute.useLoaderData();
  const activeFolders = getFoldersInActiveTree(fileTrees, folderId);

  if (!activeFolders.length) {
    return null;
  }

  return (
    <nav className="h-fit">
      <menu className="flex gap-2">
        <li>
          <TanStackRouterLink to="/" className={baseLinkStyles}>
            Home
          </TanStackRouterLink>
        </li>
        {activeFolders.map((activeFolder, index) => {
          const isLastItem = index === activeFolders.length - 1;
          const isFileSelected = fileId !== undefined;
          return (
            <BreadcrumbItem
              key={activeFolder.id}
              item={activeFolder}
              isLink={!isLastItem || isFileSelected}
            />
          );
        })}
      </menu>
    </nav>
  );
}
