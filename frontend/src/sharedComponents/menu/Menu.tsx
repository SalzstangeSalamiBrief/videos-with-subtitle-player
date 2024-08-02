import { Key } from 'react';

export interface IMenuItemBase<T> {
  id: string;
  children: Maybe<T[]>;
}

interface IMenuContext<T> {
  itemKey: keyof T;
  onRenderMenuItem: React.FC<T>;
}

// TODO SELECTED KEY?
interface IMenuProps<T extends IMenuItemBase<T>> implements IMenuContext<T extends IMenuItemBase<T>> {
  items: T[];
  itemKey: keyof T;
  onRenderMenuItem: (item: T) => JSX.Element;
}
// TODO  MOVE TO NAVIGATION AND REMOVE MENU COMPONENT
// TODO CHECK IF THIS WORKS => IF NOT USE ANY
// const MenuContext = createContext<IMenuProps<any>>({} as IMenuProps<any>);

export function Menu<T extends IMenuItemBase<T>>(prosps: IMenuProps<T>) {
  const { items, itemKey, onRenderMenuItem } = props;
  // TODO ADD CONTEXT TO SHARE PROPS WITHOUT DRILLING
  return (
    // TODO CONTEXT?
    <menu>
      {items?.map((item) => {
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
      })}
    </menu>
  );
}
