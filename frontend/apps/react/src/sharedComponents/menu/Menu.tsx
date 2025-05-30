import type { Maybe } from '@videos-with-subtitle-player/core';
import type { JSX, Key } from 'react';

export interface IMenuItemBase<T> {
  id: string;
  children: Maybe<T[]>;
}
interface IMenuProps<T extends IMenuItemBase<T>> {
  items: T[];
  itemKey: keyof T;
  onRenderMenuItem: (item: T) => JSX.Element;
  activeItemIds: string[];
}

// TODO POSSIBLE IMPROVEMENTS: CONTEXT, TOGGLE CHILDREN ON CLICK
export function Menu<T extends IMenuItemBase<T>>(props: IMenuProps<T>) {
  const { items, itemKey, onRenderMenuItem, activeItemIds } = props;
  return (
    <menu>
      {items?.map((item) => {
        return (
          <li key={item[itemKey] as Key}>
            {onRenderMenuItem(item)}
            {/* TODO ANIMATE HEIGHT */}
            {item.children &&
              item.children?.length > 0 &&
              activeItemIds.includes(item.id) && (
                <Menu<T>
                  itemKey={itemKey}
                  items={item.children}
                  onRenderMenuItem={onRenderMenuItem}
                  activeItemIds={activeItemIds}
                />
              )}
          </li>
        );
      })}
    </menu>
  );
}
