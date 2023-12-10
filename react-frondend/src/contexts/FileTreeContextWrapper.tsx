import { createContext, useEffect } from "react";
import { useGetFileTree } from "../hooks/useGetFileTree";
import { IFileTreeDto } from "../models/fileTreeDto";
import { LoadingSpinner } from "../components/loadingSpinner/LoadingSpinner";
import { ErrorMessage } from "../components/errorMessage/ErrorMessage";
import { IAudioFileDto } from "../models/audioFileDto";

interface IFileTreeContext {
  fileTrees: IFileTreeDto[];
  audioFileGroups: IAudioFileDto[][];
}

interface IFileTreeContextWrapperProps {
  children: React.ReactNode;
}

export const FileTreeContext = createContext<IFileTreeContext>({
  fileTrees: [],
  audioFileGroups: [],
});

export function FileTreeContextWrapper({
  children,
}: IFileTreeContextWrapperProps) {
  const { getFileTree, fileTrees, isLoading, error, audioFileGroups } =
    useGetFileTree();

  useEffect(() => {
    getFileTree();
  }, []);

  if (isLoading) {
    return (
      <div style={{ paddingTop: "1.5rem" }}>
        <LoadingSpinner text="Loading audio files..." />
      </div>
    );
  }

  if (error) {
    return (
      <ErrorMessage
        error={error}
        message="Something went wrong"
        description="Please try again later."
      />
    );
  }

  return (
    <FileTreeContext.Provider value={{ fileTrees, audioFileGroups }}>
      {children}
    </FileTreeContext.Provider>
  );
}
