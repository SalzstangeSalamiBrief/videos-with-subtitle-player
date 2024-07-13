import {useState} from 'react';
import {IFileTreeDto, PossibleFilesDto} from '../models/dtos/fileTreeDto';
import {IFileNode, IFileTree} from '../models/fileTree';
import {IFileDto, ISubtitleFileDto} from '../models/dtos/fileDtos';
import {FileType} from '../enums/FileType';

const baseUrl = import.meta.env.VITE_BASE_URL || '';
const path = '/api/file-tree';

const url = baseUrl + path;

export function useGetFileTree() {
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<any>(); // TODO TYPING
  const [fileTrees, setFileTrees] = useState<IFileTreeDto[]>([]);
  const [fileGroups, setFileGroups] = useState<IFileNode[][]>([]);

  async function getFileTree() {
    try {
      setIsLoading(true);
      const response = await fetch(url);
      const json: IFileTreeDto[] = await response.json();
      const transformedTree = transformDtoTreeToFileTree(json);
      setFileTrees(transformedTree);
      const flatAudioFiles = getFlatFilesGroups(transformedTree);
      setFileGroups(flatAudioFiles);
    } catch (error) {
      setError(error);
    } finally {
      setIsLoading(false);
    }
  }

  return {isLoading, error, fileTrees, getFileTree, fileGroups};
}

function getFlatFilesGroups(fileTrees: IFileTreeDto[]) {
  const fileGroups: IFileNode[][] = [];

  fileTrees.forEach((fileTree) => {
    if (fileTree.files?.length) {
      fileGroups.push(fileTree.files);
      // return;
    }

    if (fileTree.children?.length) {
      const flatGroup = getFlatFilesGroups(fileTree.children);
      fileGroups.push(...flatGroup);
    }
  });

  return fileGroups;
}

function transformDtoTreeToFileTree(dtoTree: IFileTreeDto[]): IFileTree[] {
  const fileTrees: IFileTree[] = dtoTree.map<IFileTree>((fileTree) => {
    if (fileTree.files?.length) {
      fileTree.files = replaceDtosWithFiles(fileTree.files);
    }

    if (fileTree.children?.length) {
      fileTree.children = transformDtoTreeToFileTree(fileTree.children);
    }

    return fileTree;
  });

  return fileTrees;
}

function replaceDtosWithFiles(files: PossibleFilesDto[]): IFileNode[] {
  const nodes: IFileNode[] = [];
  let remainingFiles = structuredClone(files);

  while (remainingFiles.length) {
    const currentItem = remainingFiles.shift();
    if (!currentItem) {
      throw new Error('No files are remaining to be processed.');
    }

    const isAudio = isAudioFile(currentItem);

    if (isAudio) {
      const subtitle = remainingFiles.find((file) => isSubtitleFile(file) && file.audioFileId === currentItem.id);

      const item: IFileNode = {
        id: currentItem.id,
        name: currentItem.name,
        fileType: currentItem.fileType,
        subtitleFileId: subtitle?.id,
      };

      remainingFiles = remainingFiles.filter((file) => file.id !== subtitle?.id);
      nodes.push(item);
      continue;
    }

    const isSubtitle = isSubtitleFile(currentItem);
    if (isSubtitle) {
      const audio = remainingFiles.find((file) => isAudioFile(file) && file.id === (currentItem as ISubtitleFileDto).audioFileId);
      if (!audio) {
        throw new Error(`Subtitle file '${(currentItem as ISubtitleFileDto).id}' does not have a corresponding audio file.`);
      }

      const item: IFileNode = {
        id: audio.id,
        name: audio.name,
        fileType: audio.fileType,
        subtitleFileId: (currentItem as IFileDto).id,
      };

      remainingFiles = remainingFiles.filter((file) => file.id !== item.id);
      nodes.push(item);
      continue;
    }

    nodes.push(currentItem);
  }

  return nodes;
}

function isSubtitleFile(file: PossibleFilesDto): file is ISubtitleFileDto {
  return file.fileType === FileType.SUBTITLE;
}

function isAudioFile(file: PossibleFilesDto): file is IFileDto {
  return file.fileType === FileType.AUDIO;
}
