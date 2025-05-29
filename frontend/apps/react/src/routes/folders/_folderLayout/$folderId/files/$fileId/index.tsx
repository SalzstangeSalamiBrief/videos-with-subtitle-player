import { PlayerCompound } from '$features/playerCompound/PlayerCompound';
import type { IFileNode } from '$models/fileTree';
import { ErrorMessage } from '$sharedComponents/errorMessage/ErrorMessage';
import { createFileRoute } from '@tanstack/react-router';
import { Route as RootLayoutRoute } from '../../../../../__root';
export const Route = createFileRoute(
  '/folders/_folderLayout/$folderId/files/$fileId/',
)({
  component: FilePage,
  // TODO ADD META  => SET TITLE AS TITLE OF THE FILE=> MAYBE USE ROUTER CONTEXT
});

function FilePage() {
  const { fileGroups } = RootLayoutRoute.useLoaderData();
  const { fileId } = Route.useParams();
  const [siblings, currentFile] = getCurrentNodeWithSiblings(
    fileGroups,
    fileId,
  );

  if (!currentFile) {
    return (
      <ErrorMessage
        error="Could not find file."
        message="Could not find file."
      />
    );
  }

  return (
    <div className="grid">
      <h1 className="m-0 text-lg font-bold">{currentFile.name}</h1>
      <PlayerCompound currentFile={currentFile} siblings={siblings} />
    </div>
  );
}

function getCurrentNodeWithSiblings(
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
