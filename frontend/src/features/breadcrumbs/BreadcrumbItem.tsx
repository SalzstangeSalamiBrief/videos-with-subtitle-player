import { baseLinkStyles } from '$lib/styles/baseLinkStyles';
import { getFoldersInActiveTree } from '$lib/utilities/getFoldersInActiveTree';
import type { IFileTree } from '$models/fileTree/fileTree';
import { Link as TanStackRouterLink, useParams } from '@tanstack/react-router';
import { Route as RootLayoutRoute } from '../../routes/__root';

interface IBreadcrumbItemProps {
  item: IFileTree;
  isLink: boolean;
}

export function BreadcrumbItem({ isLink, item }: IBreadcrumbItemProps) {
  const { folderId } = useParams({ strict: false });
  const { fileTrees } = RootLayoutRoute.useLoaderData();
  const items = getFoldersInActiveTree(fileTrees, folderId);

  if (!items.length) {
    return null;
  }

  if (!isLink) {
    return (
      <li className={baseClasses} title={item.name} aria-current="page">
        &gt; {item.name}
      </li>
    );
  }

  return (
    <li>
      <TanStackRouterLink
        to="/folders/$folderId"
        params={{ folderId: item.id }}
        search={{ activeTab: 0 }}
        className={linkClasses}
        title={item.name}
      >
        &gt; {item.name}
      </TanStackRouterLink>
    </li>
  );
}

const baseClasses =
  'block max-w-[30ch] overflow-x-hidden text-ellipsis whitespace-nowrap';
const linkClasses = `${baseClasses} ${baseLinkStyles}`;
