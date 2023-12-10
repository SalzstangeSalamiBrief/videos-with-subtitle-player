import { useState } from "react";
import { IFileTreeDto } from "../models/fileTreeDto";

const baseUrl = "http://localhost:3000"; // todo remove
const path = "/api/file-tree";

const url = baseUrl + path;

export function useGetFileTree() {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<any>(); // TODO TYPING
  const [fileTrees, setFileTrees] = useState<IFileTreeDto[]>([]);

  const getFileTree = async () => {
    try {
      setIsLoading(true);
      const response = await fetch(url);
      const json: IFileTreeDto[] = await response.json();
      setFileTrees(json);
    } catch (error) {
      setError(error);
    } finally {
      setIsLoading(false);
    }
  };

  return { isLoading, error, fileTrees, getFileTree };
}
