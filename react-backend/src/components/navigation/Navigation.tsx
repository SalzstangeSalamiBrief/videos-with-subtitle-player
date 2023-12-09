import { Menu, MenuProps } from "antd";
import { IFileTreeDto } from "../../models/fileTreeDto";
import { useEffect } from "react";
import { useGetFileTree } from "../../hooks/useGetFileTree";

export function Navigation() {
  const { getFileTree, fileTrees } = useGetFileTree(); // todo error isLoading => MAYBE IN NAVIGATIOPN?
  useEffect(() => {
    getFileTree();
  }, []);
  const menuItems = fileTrees.map(getMenuTree);

  //   todo collapsible
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
      ...fileTree.audioFiles.map((audioFile) => ({
        key: `${fileTree.id}-${audioFile.name}`,
        label: audioFile.name,
        type: "item",
        onClick: () => console.log(audioFile),
      })),
    ];
  }

  const menuItem: MenuItem = {
    key: fileTree.id,
    label: fileTree.name,
    children,
  };

  return menuItem;
};
