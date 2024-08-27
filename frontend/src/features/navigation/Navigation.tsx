// import { Menu, MenuProps } from 'antd';
import { useContext } from 'react';
import { Menu } from '$sharedComponents/menu/Menu';
import { IFileTreeDto } from '$models/dtos/fileTreeDto';
import styles from './Navigation.module.css';
import { NavigationItem } from './NavigationItem';
import { useParams } from '@tanstack/react-router';
import { getFoldersInActiveTree } from '$lib/utilities/getFoldersInActiveTree';
import { FileTreeContext } from '$contexts/FileTreeContextWrapper';

export function Navigation() {
  const { folderId } = useParams({ strict: false });
  const { fileTrees } = useContext(FileTreeContext);
  const activeFolders = getFoldersInActiveTree(fileTrees, folderId);

  return (
    <nav className={styles.navigation}>
      <Menu<IFileTreeDto>
        itemKey="id"
        items={fileTrees}
        activeItemIds={activeFolders.map((a) => a.id)}
        onRenderMenuItem={(item: IFileTreeDto) => {
          const isActive = activeFolders.some((a) => a.id === item.id);
          return (
            <NavigationItem
              item={item}
              isActive={isActive}
              hasChildren={Boolean(item.children?.length)}
            />
          );
        }}
      />
    </nav>
  );
}
