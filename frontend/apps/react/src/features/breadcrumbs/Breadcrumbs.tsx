import { getFoldersInActiveTree } from '$lib/utilities/getFoldersInActiveTree';
import { useParams } from '@tanstack/react-router';
import { BreadcrumbItem } from './BreadcrumbItem';
import { Link as TanStackRouterLink } from '@tanstack/react-router';
import { baseLinkStyles } from '$lib/styles/baseLinkStyles';
import { Route as RootLayoutRoute } from '../../routes/__root';

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
