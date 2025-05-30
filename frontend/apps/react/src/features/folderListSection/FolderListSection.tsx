import { FolderList } from '$sharedComponents/folderList/FolderList';
import type { IFileTree } from '@videos-with-subtitle-player/core';

interface IFolderListSectionProps {
  folders: IFileTree[];
}

export function FolderListSection({ folders }: IFolderListSectionProps) {
  if (!folders.length) {
    return <p>This folder contains no subfolders</p>;
  }

  return <FolderList folders={folders} />;
}
