// import { Menu, MenuProps } from 'antd';
import { useContext } from 'react';
import { FileTreeContext } from '$contexts/FileTreeContextWrapper';
import { Menu } from '$sharedComponents/menu/Menu';
import { IFileTreeDto } from '$models/dtos/fileTreeDto';
import './Navigation.css';
import { NavigationItem } from './NavigationItem';
import { useParams } from '@tanstack/react-router';

export function Navigation() {
  const { folderId } = useParams({ strict: false });
  const { fileTrees } = useContext(FileTreeContext);
  const activeFileIds = getActiveFolderIds(fileTrees, folderId);

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
              hasChildren={Boolean(item.children?.length)}
            />
          );
        }}
      />
    </nav>
  );
}

function getActiveFolderIds(
  nodes: Maybe<IFileTreeDto[]>,
  folderId: Maybe<string>,
): string[] {
  if (!folderId) {
    return [];
  }

  if (!nodes) {
    return [];
  }

  let activeFileIds: string[] = [];
  for (let i = 0; i < nodes.length; i += 1) {
    const currentNode = nodes[i];
    const [hasMatch, matchingIds] = getActiveChildNodes(currentNode, folderId);
    if (hasMatch) {
      activeFileIds = matchingIds;
      break;
    }
  }

  return activeFileIds;
}

function getActiveChildNodes(
  currentNode: IFileTreeDto,
  folderId: Maybe<string>,
): [hasMatch: boolean, childIds: string[]] {
  if (!folderId) {
    return [false, []];
  }

  const result = [currentNode.id];
  if (currentNode.id === folderId) {
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
      folderId,
    );

    if (hasChildMatch) {
      result.push(...matchingChildIds);
      hasMatch = true;
      break;
    }
  }

  return [hasMatch, result];
}
