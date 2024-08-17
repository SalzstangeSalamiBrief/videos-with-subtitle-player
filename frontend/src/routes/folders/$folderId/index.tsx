import { createFileRoute, useParams } from '@tanstack/react-router';
import { FolderListSection } from '$features/folderListSection/FolderListSection';
import { FileTreeContext } from '$contexts/FileTreeContextWrapper';
import { IFileTreeDto } from '$models/dtos/fileTreeDto';
import { useContext } from 'react';
import { ErrorMessage } from '$sharedComponents/errorMessage/ErrorMessage';
import { FileListSection } from '$features/fileListSection/FileListSection';
import { Breadcrumbs } from '$features/breadcrumbs/Breadcrumbs';
import { ITab, Tabs } from '$sharedComponents/tabs/Tabs';

export const Route = createFileRoute('/folders/$folderId/')({
  component: AudioFilePage,
});

function AudioFilePage() {
  const { fileTrees } = useContext(FileTreeContext);
  const { folderId } = useParams({ strict: false });
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
      content: <FolderListSection selectedFolder={selectedFolder} />,
    },
    {
      label: `Video and audio files (${selectedFolder.files.length})`,
      content: <FileListSection selectedFolder={selectedFolder} />,
    },
  ];

  return (
    <div className="grid gap-4">
      <Breadcrumbs />
      <Tabs tabs={tabs} label="Content" />
    </div>
  );
}

function getFolderFromFileTree(
  fileTrees: IFileTreeDto[],
  folderId: string,
): Maybe<IFileTreeDto> {
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
