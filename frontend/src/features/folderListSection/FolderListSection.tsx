import type { IFileTree } from '$models/fileTree/fileTree';
import { FolderList } from '$sharedComponents/folderList/FolderList';

interface IFolderListSectionProps {
  folders: IFileTree[];
}

export function FolderListSection({ folders }: IFolderListSectionProps) {
  if (!folders.length) {
    return <p>This folder contains no subfolders</p>;
  }

  return <FolderList folders={folders} />;
}
