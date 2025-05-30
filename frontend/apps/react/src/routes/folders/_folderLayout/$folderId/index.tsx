import { FileListSection } from '$features/fileListSection/FileListSection';
import { FolderListSection } from '$features/folderListSection/FolderListSection';
import { ImageListSection } from '$features/imageListSection/ImageListSection';
import { ErrorMessage } from '$sharedComponents/errorMessage/ErrorMessage';
import { Tabs, type ITab } from '$sharedComponents/tabs/Tabs';
import {
  createFileRoute,
  useNavigate,
  useParams,
  useSearch,
} from '@tanstack/react-router';
import type { IFileTree, Maybe } from '@videos-with-subtitle-player/core';
import { Route as RootLayoutRoute } from '../../../__root';
import type { IFolderLayoutSearchParams } from '../../_folderLayout';

export const Route = createFileRoute('/folders/_folderLayout/$folderId/')({
  component: AudioFilePage,
  // TODO ADD META  => SET TITLE AS TITLE OF THE FOLDER => MAYBE USE ROUTER CONTEXT
});

function AudioFilePage() {
  const { fileTrees } = RootLayoutRoute.useLoaderData();
  const navigate = useNavigate({ from: Route.fullPath });
  const searchParams: IFolderLayoutSearchParams = useSearch({ strict: false });
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
      content: <FolderListSection folders={selectedFolder.children} />,
    },
    {
      label: `Videos (${selectedFolder.videos?.length})`,
      content: (
        <FileListSection
          folderId={selectedFolder.id}
          files={selectedFolder.videos}
        />
      ),
    },
    {
      label: `Audio files (${selectedFolder.audios.length})`,
      content: (
        <FileListSection
          folderId={selectedFolder.id}
          files={selectedFolder.audios}
        />
      ),
    },
    {
      label: `Image (${selectedFolder.images.length})`,
      content: <ImageListSection images={selectedFolder.images} />,
    },
  ];

  const activeTabIndex = getActiveTabIndex(searchParams.activeTab, tabs.length);
  if (activeTabIndex !== searchParams.activeTab) {
    navigate({
      search: () => ({
        activeTab: activeTabIndex,
      }),
    });
  }

  return (
    <Tabs
      tabs={tabs}
      label="Content"
      activeTabIndex={activeTabIndex}
      onChangeTab={(newIndex: number) =>
        navigate({
          search: () => ({ activeTab: newIndex }),
        })
      }
    />
  );
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

function getActiveTabIndex(
  input: number | undefined,
  numberOfTabs: number,
): number {
  if (input === undefined) {
    return 0;
  }

  if (Number.isNaN(input)) {
    return 0;
  }

  if (input < 0) {
    return 0;
  }

  if (numberOfTabs < input) {
    return numberOfTabs - 1;
  }

  return input;
}
