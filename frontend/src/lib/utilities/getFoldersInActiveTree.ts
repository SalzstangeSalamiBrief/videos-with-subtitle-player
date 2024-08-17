import { IFileTreeDto } from '$models/dtos/fileTreeDto';

export function getFoldersInActiveTree(
  nodes: Maybe<IFileTreeDto[]>,
  folderId: Maybe<string>,
): IFileTreeDto[] {
  if (!folderId) {
    return [];
  }

  if (!nodes) {
    return [];
  }

  let activeFileIds: IFileTreeDto[] = [];
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
): [hasMatch: boolean, childIds: IFileTreeDto[]] {
  if (!folderId) {
    return [false, []];
  }

  const result = [currentNode];
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