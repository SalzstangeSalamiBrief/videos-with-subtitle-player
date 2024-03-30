import { Menu, MenuProps } from "antd";
import { IFileTreeDto } from "../../models/dtos/fileTreeDto";
import { useContext, useMemo } from "react";
import { generatePath, useParams } from "react-router-dom";
import { Link as ReactRouterLink } from "react-router-dom";
import { FileTreeContext } from "../../contexts/FileTreeContextWrapper";

export function Navigation() {
  const { audioId } = useParams();
  const { fileTrees } = useContext(FileTreeContext);
  const menuItems = useMemo(
    () =>
      fileTrees.sort((a, b) => a.name.localeCompare(b.name)).map(getMenuTree),
    [fileTrees]
  );

  return (
    <nav style={{ height: "100%" }}>
      <Menu items={menuItems} mode="inline" selectedKeys={[audioId ?? ""]} />
    </nav>
  );
}

type MenuItem = Required<MenuProps>["items"][number];
function getMenuTree(fileTree: IFileTreeDto): MenuItem {
  let children: MenuItem[] = [];
  if (fileTree.children?.length) {
    children = [...children, ...fileTree.children.map(getMenuTree)];
  }

  if (fileTree.files?.length) {
    children = [
      ...children,
      ...fileTree.files.map<MenuItem>((file) => {
        const targetUrl = generatePath("/content/:fileId", {
          fileId: file.id,
        });
        return {
          key: file.id,
          label: (
            <ReactRouterLink to={targetUrl} title={file.name}>
              {file.name}
            </ReactRouterLink>
          ),
          type: "item",
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
