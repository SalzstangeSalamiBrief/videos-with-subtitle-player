// import { Menu, MenuProps } from 'antd';
import { useContext } from 'react';
import { FileTreeContext } from '$contexts/FileTreeContextWrapper';
import { Menu } from '$sharedComponents/menu/Menu';
import { IFileTreeDto } from '$models/dtos/fileTreeDto';
import './Navigation.css';
import { NavigationItem } from './NavigationItem';
import { useParams } from '@tanstack/react-router';
export function Navigation() {
  const { fileId } = useParams({ strict: false });
  const { fileTrees } = useContext(FileTreeContext);
  const activeFileIds = getActiveFileIds(fileTrees, fileId);
  // const menuItems = useMemo(
  //   () =>
  //     fileTrees.sort((a, b) => a.name.localeCompare(b.name)).map(getMenuTree),
  //   [fileTrees],
  // );
  return (
    <nav className="navigation">
      <Menu<IFileTreeDto>
        itemKey="id"
        items={fileTrees}
        activeItemIds={activeFileIds}
        onRenderMenuItem={(item: IFileTreeDto) => {
          const isActive = activeFileIds.includes(item.id);
          return (
            <NavigationItem
              item={item}
              isActive={isActive}
              hasChildren={Boolean(item.children)}
            />
          );
        }}
      />
    </nav>
  );
}

function getActiveFileIds(
  nodes: Maybe<IFileTreeDto[]>,
  fileId: Maybe<string>,
): string[] {
  if (!fileId) {
    return [];
  }

  if (!nodes) {
    return [];
  }

  let activeFileIds: string[] = [];
  for (let i = 0; i < nodes.length; i += 1) {
    const currentNode = nodes[i];
    const [hasMatch, matchingIds] = getActiveChildNodes(currentNode, fileId);
    if (hasMatch) {
      activeFileIds = matchingIds;
      break;
    }
  }

  return activeFileIds;
}

function getActiveChildNodes(
  currentNode: IFileTreeDto,
  fileId: Maybe<string>,
): [hasMatch: boolean, childIds: string[]] {
  if (!fileId) {
    return [false, []];
  }

  const result = [currentNode.id];
  if (currentNode.id === fileId) {
    return [true, result];
  }

  if (!currentNode.children?.length) {
    return [false, []];
  }

  let hasMatch = false;
  for (let j = 0; j < currentNode.children.length; j += 1) {
    const currentChild = currentNode.children[j];
    const [hasChildMatch, matchingChildIds] = getActiveChildNodes(
      currentChild,
      fileId,
    );

    if (hasChildMatch) {
      result.push(...matchingChildIds);
      hasMatch = true;
      break;
    }
  }

  return [hasMatch, result];
}
