import { createFileRoute, useParams } from '@tanstack/react-router';
import { FolderListSection } from '$features/folderListSection/FolderListSection';
import { ErrorMessage } from '$sharedComponents/errorMessage/ErrorMessage';
import { FileListSection } from '$features/fileListSection/FileListSection';
import { ITab, Tabs } from '$sharedComponents/tabs/Tabs';
import { Route as RootLayoutRoute } from '../../../__root';
import { IFileTree } from '$models/fileTree';
import { ImageListSection } from '$features/imageListSection/ImageListSection';

export const Route = createFileRoute('/folders/_folderLayout/$folderId/')({
  component: AudioFilePage,
  // TODO ADD META  => SET TITLE AS TITLE OF THE FOLDER => MAYBE USE ROUTER CONTEXT
});

function AudioFilePage() {
  const { fileTrees } = RootLayoutRoute.useLoaderData();
  const { folderId } = useParams({ strict: false });
  const baka = Route.useLoaderData();
  console.log('RR', baka);

  console.log('TT', folderId, fileTrees);
  const selectedFolder = getFolderFromFileTree(fileTrees, folderId);
  if (!selectedFolder) {
    const message = `Could not find folder with id '${folderId}'`;
    return (
      <ErrorMessage
        error={message}
        message={message}
        description="Please try again"
      />
    );
  }

  const tabs: ITab[] = [
    {
      label: `Subfolders (${selectedFolder.children?.length})`,
      content: <FolderListSection folders={selectedFolder.children} />,
    },
    // TODO INCONSISTENT USAGE OF TABS => IMAGE TAB AS REFERENCE
    {
      label: `Video and audio files (${selectedFolder.continuousFiles.length})`,
      content: (
        <FileListSection
          folderId={selectedFolder.id}
          files={selectedFolder.continuousFiles}
        />
      ),
    },
    {
      label: `Image (${selectedFolder.images.length})`,
      content: <ImageListSection images={selectedFolder.images} />,
    },
  ];

  return <Tabs tabs={tabs} label="Content" />;
}

function getFolderFromFileTree(
  fileTrees: IFileTree[],
  folderId: string | undefined,
): Maybe<IFileTree> {
  for (let i = 0; i < fileTrees.length; i += 1) {
    const currentTree = fileTrees[i];
    if (currentTree.id === folderId) {
      return currentTree;
    }

    if (!currentTree.children.length) {
      continue;
    }

    const matchingFolderFromChild = getFolderFromFileTree(
      currentTree.children,
      folderId,
    );
    if (matchingFolderFromChild) {
      return matchingFolderFromChild;
    }
  }
}
