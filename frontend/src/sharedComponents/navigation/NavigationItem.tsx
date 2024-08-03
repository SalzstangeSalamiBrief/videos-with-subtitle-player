import { IFileTreeDto } from '$models/dtos/fileTreeDto';
import { ChevronDownIcon } from '@heroicons/react/24/outline';
import { Link as TanStackLink } from '@tanstack/react-router';

interface INavigationItemProps {
  item: IFileTreeDto;
  isActive: boolean;
  hasChildren: boolean;
}

export function NavigationItem({
  item,
  isActive,
  hasChildren,
}: INavigationItemProps) {
  return (
    <TanStackLink
      className={getNavigationItemStyles(isActive)}
      to="/folders/$folderId"
      title={item.name}
      params={{ folderId: item.id }}
    >
      <ItemStateIndicator
        hasChildren={hasChildren}
        isActive={isActive}
        itemName={item.name}
      />
      {item.name}
    </TanStackLink>
  );
}

interface IItemStateIndicatorProps {
  isActive: boolean;
  hasChildren: boolean;
  itemName: string;
}
function ItemStateIndicator({
  hasChildren,
  isActive,
  itemName,
}: IItemStateIndicatorProps) {
  if (!hasChildren) {
    return null;
  }

  return (
    <ChevronDownIcon
      width={20}
      className={`shrink-0 transition-rotate duration-200 ease-linear ${isActive ? 'rotate-180' : 'rotate-0'}`}
      aria-label={`Close ${itemName}`}
    />
  );
}

function getNavigationItemStyles(isActive: boolean): string {
  const baseClasses = 'block px-2 py-2 rounded flex gap-4';

  if (!isActive) {
    return `${baseClasses} hover:bg-slate-700`;
  }

  return `${baseClasses} text-fuchsia-50 bg-fuchsia-800 hover:bg-fuchsia-700`;
}
