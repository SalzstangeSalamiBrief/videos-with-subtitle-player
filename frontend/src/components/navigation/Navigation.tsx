import { Menu, MenuProps } from 'antd';
import { IFileTreeDto } from '../../models/dtos/fileTreeDto';
import { useContext, useMemo } from 'react';
import { Link as TanStackLink, useParams } from '@tanstack/react-router';
import { FileTreeContext } from '../../contexts/FileTreeContextWrapper';

export function Navigation() {
  const { fileId } = useParams({ strict: false });
  const { fileTrees } = useContext(FileTreeContext);
  const menuItems = useMemo(
    () =>
      fileTrees.sort((a, b) => a.name.localeCompare(b.name)).map(getMenuTree),
    [fileTrees],
  );

  return (
    <nav style={{ height: '100%' }}>
      <Menu items={menuItems} mode="inline" selectedKeys={[fileId ?? '']} />
    </nav>
  );
}

type MenuItem = Required<MenuProps>['items'][number];
function getMenuTree(fileTree: IFileTreeDto): MenuItem {
  let children: MenuItem[] = [];
  if (fileTree.children?.length) {
    children = [...children, ...fileTree.children.map(getMenuTree)];
  }

  if (fileTree.files?.length) {
    children = [
      ...children,
      ...fileTree.files.map<MenuItem>((file) => {
        return {
          key: file.id,
          label: (
            <TanStackLink
              to="files/$fileId"
              title={file.name}
              params={{ fileId: file.id }}
            >
              {file.name}
            </TanStackLink>
          ),
          type: 'item',
        };
      }),
    ];
  }

  const menuItem: MenuItem = {
    key: fileTree.id,
    label: <span title={fileTree.name}>{fileTree.name}</span>,
    children,
  };

  return menuItem;
}
