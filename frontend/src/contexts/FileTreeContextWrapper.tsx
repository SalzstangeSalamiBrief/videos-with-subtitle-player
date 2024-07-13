import {createContext, useEffect} from 'react';
import {useGetFileTree} from '../hooks/useGetFileTree';
import {IFileTreeDto} from '../models/dtos/fileTreeDto';
import {LoadingSpinner} from '../components/loadingSpinner/LoadingSpinner';
import {ErrorMessage} from '../components/errorMessage/ErrorMessage';
import {IFileNode} from '../models/fileTree';

interface IFileTreeContext {
  fileTrees: IFileTreeDto[];
  fileGroups: IFileNode[][];
}

interface IFileTreeContextWrapperProps {
  children: React.ReactNode;
}

export const FileTreeContext = createContext<IFileTreeContext>({
  fileTrees: [],
  fileGroups: [],
});

export function FileTreeContextWrapper({children}: IFileTreeContextWrapperProps) {
  const {getFileTree, fileTrees, isLoading, error, fileGroups} = useGetFileTree();

  useEffect(() => {
    getFileTree();
  }, []);

  if (isLoading) {
    return (
      <div style={{paddingTop: '1.5rem'}}>
        <LoadingSpinner text="Loading audio files..." />
      </div>
    );
  }

  if (error) {
    return <ErrorMessage error={error} message="Something went wrong" description="Please try again later." />;
  }

  return <FileTreeContext.Provider value={{fileTrees, fileGroups}}>{children}</FileTreeContext.Provider>;
}
