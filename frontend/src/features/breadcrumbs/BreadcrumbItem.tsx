import { getFoldersInActiveTree } from '$lib/utilities/getFoldersInActiveTree';
import { useParams } from '@tanstack/react-router';
import { Link as TanStackRouterLink } from '@tanstack/react-router';
import { IFileTreeDto } from '$models/dtos/fileTreeDto';
import { baseLinkStyles } from '$lib/styles/baseLinkStyles';
import { Route as RootLayoutRoute } from '../../routes/__root';

interface IBreadcrumbItemProps {
  item: IFileTreeDto;
  isLastItem: boolean;
}

export function BreadcrumbItem({ isLastItem, item }: IBreadcrumbItemProps) {
  const { folderId } = useParams({ strict: false });
  const { fileTrees } = RootLayoutRoute.useLoaderData();
  const items = getFoldersInActiveTree(fileTrees, folderId);

  if (!items.length) {
    return null;
  }

  if (isLastItem) {
    return (
      <li className={baseClasses} title={item.name}>
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
