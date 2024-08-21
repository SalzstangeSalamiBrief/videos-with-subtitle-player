import { FileType } from '$enums/FileType';
import { ISubtitleFileDto, IFileDto } from '$models/dtos/fileDtos';
import { IFileTreeDto, PossibleFilesDto } from '$models/dtos/fileTreeDto';
import { IFileNode, IFileTree } from '$models/fileTree';

interface IGetFileTreeSelectReturn {
  fileTrees: IFileTreeDto[];
  fileGroups: IFileNode[][];
}

export function getFileTreeSelect(
  input: IFileTreeDto[],
): IGetFileTreeSelectReturn {
  const fileTrees = transformDtoTreeToFileTree(input);
  const fileGroups = getFlatFilesGroups(fileTrees);
  console.log('select', fileTrees);
  return { fileGroups, fileTrees };
}

function getFlatFilesGroups(fileTrees: IFileTreeDto[]) {
  const fileGroups: IFileNode[][] = [];

  fileTrees.forEach((fileTree) => {
    if (fileTree.files?.length) {
      fileGroups.push(fileTree.files);
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

function replaceDtosWithFiles(input: PossibleFilesDto[]): IFileNode[] {
  const nodes: IFileNode[] = [];
  const subtitleFiles = input.filter((file) => isSubtitleFile(file));
  const mediaFiles = input.filter((file) => !isSubtitleFile(file));

  while (mediaFiles.length) {
    const currentItem = mediaFiles.shift();
    if (!currentItem) {
      throw new Error('No files are remaining to be processed.');
    }

    const isAudio = isAudioFile(currentItem);
    if (!isAudio) {
      nodes.push(currentItem);
      continue;
    }

    const matchingSubtitleFile = subtitleFiles.find(
      (file) => file.audioFileId === currentItem.id,
    );

    const item: IFileNode = {
      id: currentItem.id,
      name: currentItem.name,
      fileType: currentItem.fileType,
      subtitleFileId: matchingSubtitleFile?.id,
    };

    nodes.push(item);
  }

  return nodes;
}

function isSubtitleFile(file: PossibleFilesDto): file is ISubtitleFileDto {
  return file.fileType === FileType.SUBTITLE;
}

function isAudioFile(file: PossibleFilesDto): file is IFileDto {
  return file.fileType === FileType.AUDIO;
}
