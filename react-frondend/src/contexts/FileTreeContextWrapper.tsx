import { createContext, useEffect } from "react";
import { useGetFileTree } from "../hooks/useGetFileTree";
import { IFileTreeDto } from "../models/fileTreeDto";
import { LoadingSpinner } from "../components/loadingSpinner/LoadingSpinner";
import { ErrorMessage } from "../components/errorMessage/ErrorMessage";

interface IFileTreeContext {
  fileTrees: IFileTreeDto[];
}

interface IFileTreeContextWrapperProps {
  children: React.ReactNode;
}

export const FileTreeContext = createContext<IFileTreeContext>({
  fileTrees: [],
});

export function FileTreeContextWrapper({
  children,
}: IFileTreeContextWrapperProps) {
  const { getFileTree, fileTrees, isLoading, error } = useGetFileTree(); // todo error isLoading => MAYBE IN NAVIGATIOPN?
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
    <FileTreeContext.Provider value={{ fileTrees }}>
      {children}
    </FileTreeContext.Provider>
  );
}
