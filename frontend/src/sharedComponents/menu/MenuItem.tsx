// import { IMenuItem } from "./Menu";

import { IMenuItemBase } from './Menu';

interface IMenuItemProps<T extends IMenuItemBase<T>> {
  item: T;
}

export function MenuItem<T extends IMenuItemBase<T>>({
  item,
}: IMenuItemProps<T>) {
  return (
    <li key={item[itemKey] as Key}>
      <a className="block hover:bg-fuchsia-700 px-2 py-2">
        {onRenderMenuItem(item)}
      </a>
      {item.children && item.children?.length > 0 && (
        <Menu<T>
          itemKey={itemKey}
          items={item.children}
          onRenderMenuItem={onRenderMenuItem}
        />
      )}
    </li>
  );
}
