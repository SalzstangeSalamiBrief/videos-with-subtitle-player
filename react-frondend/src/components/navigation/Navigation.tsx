import { Menu, MenuProps } from "antd";
import { IFileTreeDto } from "../../models/fileTreeDto";
import { useContext } from "react";
import { generatePath } from "react-router-dom";
import { Link as ReactRouterLink } from "react-router-dom";
import { FileTreeContext } from "../../contexts/FileTreeContextWrapper";

export function Navigation() {
  const fileTreeContext = useContext(FileTreeContext);
  const menuItems = fileTreeContext.fileTrees.map(getMenuTree);

  return <Menu items={menuItems} mode="inline" style={{ height: "100%" }} />;
}

type MenuItem = Required<MenuProps>["items"][number];
const getMenuTree = (fileTree: IFileTreeDto): MenuItem => {
  let children: MenuItem[] = [];
  if (fileTree.children?.length) {
    children = [...children, ...fileTree.children.map(getMenuTree)];
  }

  if (fileTree.audioFiles?.length) {
    children = [
      ...children,
      ...fileTree.audioFiles.map<MenuItem>((audioFile) => {
        const targetUrl = generatePath("/audio/:audioId", {
          audioId: audioFile.audioFile.id,
        });
        return {
          key: `${fileTree.id}-${audioFile.name}`,
          label: (
            <ReactRouterLink to={targetUrl} title={audioFile.name}>
              {audioFile.name}
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
};
