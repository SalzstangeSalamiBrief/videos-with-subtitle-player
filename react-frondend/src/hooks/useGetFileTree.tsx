import { useState } from "react";
import { IFileTreeDto } from "../models/fileTreeDto";
import { IAudioFileDto } from "../models/audioFileDto";

const baseUrl = "http://localhost:3000"; // todo remove
const path = "/api/file-tree";

const url = baseUrl + path;

export function useGetFileTree() {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<any>(); // TODO TYPING
  const [fileTrees, setFileTrees] = useState<IFileTreeDto[]>([]);
  const [audioFileGroups, SetAudioFileGroups] = useState<IAudioFileDto[][]>([]);

  const getFileTree = async () => {
    try {
      setIsLoading(true);
      const response = await fetch(url);
      const json: IFileTreeDto[] = await response.json();
      setFileTrees(json);
      const flatAudioFiles = getFlatAudioFiles(json);
      SetAudioFileGroups(flatAudioFiles);
    } catch (error) {
      setError(error);
    } finally {
      setIsLoading(false);
    }
  };

  return { isLoading, error, fileTrees, getFileTree, audioFileGroups };
}

const getFlatAudioFiles = (fileTrees: IFileTreeDto[]) => {
  const audioFileGroups: IAudioFileDto[][] = [];

  fileTrees.forEach((fileTree) => {
    if (fileTree.audioFiles?.length) {
      audioFileGroups.push(fileTree.audioFiles);
      return;
    }

    if (fileTree.children?.length) {
      audioFileGroups.push(...getFlatAudioFiles(fileTree.children));
    }
  });

  return audioFileGroups;
};
