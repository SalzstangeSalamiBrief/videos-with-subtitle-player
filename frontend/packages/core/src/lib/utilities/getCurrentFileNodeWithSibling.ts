import type { IFileNode } from '../../models/fileTree/fileTree';
import type { Maybe } from '../../models/maybe';

export function getCurrentNodeWithSiblings(
  fileGroups: IFileNode[][],
  fileId: string,
): [siblings: IFileNode[], currentFile: Maybe<IFileNode>] {
  const siblings = fileGroups.find((group) =>
    group.some((file) => file.id === fileId),
  );
  if (!siblings) {
    return [[], undefined];
  }

  const currentNode = siblings.find((file) => file.id === fileId);
  if (!currentNode) {
    return [[], undefined];
  }

  return [siblings, currentNode];
}
