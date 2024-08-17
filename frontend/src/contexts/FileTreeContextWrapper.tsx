import { IFileTreeDto } from '$models/dtos/fileTreeDto';
import { IFileNode } from '$models/fileTree';
import { createContext } from 'react';

interface IFileTreeContext {
  fileTrees: IFileTreeDto[];
  fileGroups: IFileNode[][];
}

interface IFileTreeContextWrapperProps {
  input: IFileTreeContext;
  children: React.ReactNode;
}

export const FileTreeContext = createContext<IFileTreeContext>({
  fileTrees: [],
  fileGroups: [],
});

export function FileTreeContextWrapper({
  input,
  children,
}: IFileTreeContextWrapperProps) {
  return (
    <FileTreeContext.Provider value={input}>
      {children}
    </FileTreeContext.Provider>
  );
}
