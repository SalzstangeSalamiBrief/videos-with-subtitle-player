import { getFoldersInActiveTree } from '$lib/utilities/getFoldersInActiveTree';
import { useParams } from '@tanstack/react-router';
import { BreadcrumbItem } from './BreadcrumbItem';
import { Link as TanStackRouterLink } from '@tanstack/react-router';
import { baseLinkStyles } from '$lib/styles/baseLinkStyles';
import { Route as RootLayoutRoute } from '../../routes/__root';

export function Breadcrumbs() {
  const { folderId } = useParams({ strict: false });
  const { fileTrees } = RootLayoutRoute.useLoaderData();
  // TODO: ON INITIAL LOAD THE CONTEXT IST EMPTY => MAYBE USE LOADER FUNCTIONS FROM THE ROUTER
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
