import {
  isAudioFile,
  isImageFile,
  isSubtitleFile,
  isVideoFile,
} from '../../lib/type-predicates/file-type-predicates';
import type {
  IFileDto,
  ISubtitleFileDto,
} from '../../models/fileTree/dtos/fileDtos';
import type { IFileTreeDto } from '../../models/fileTree/dtos/fileTreeDto';
import type { IFileNode, IFileTree } from '../../models/fileTree/fileTree';

export interface IGetFileTreeSelectReturn {
  fileTrees: IFileTree[];
  fileGroups: IFileNode[][];
}

export function getFileTreeSelect(
  input: IFileTreeDto[],
): IGetFileTreeSelectReturn {
  const fileTrees = transformDtoTreeToFileTree(input);
  const fileGroups = getFlatFilesGroups(fileTrees);
  return { fileGroups, fileTrees };
}

function getFlatFilesGroups(fileTrees: IFileTree[]) {
  const fileGroups: IFileNode[][] = [];

  fileTrees.forEach((fileTree) => {
    if (fileTree.videos.length) {
      fileGroups.push(fileTree.videos);
    }

    if (fileTree.audios.length) {
      fileGroups.push(fileTree.audios);
    }

    if (fileTree.images.length) {
      fileGroups.push(fileTree.images);
    }

    if (fileTree.children?.length) {
      const flatGroup = getFlatFilesGroups(fileTree.children);
      fileGroups.push(...flatGroup);
    }
  });

  return fileGroups;
}

function transformDtoTreeToFileTree(
  dtoTree: IFileTreeDto[] | undefined,
): IFileTree[] {
  if (!dtoTree?.length) {
    return [];
  }

  const fileTrees: IFileTree[] = dtoTree.map<IFileTree>((fileTree) => {
    const subtitleFiles = fileTree.files?.filter(isSubtitleFile);
    const images: IFileNode[] = fileTree.files
      ?.filter(isImageFile)
      .map((f) => transformFileDtoToFile(f, subtitleFiles));

    const videos: IFileNode[] = fileTree.files
      ?.filter(isVideoFile)
      .map((f) => transformFileDtoToFile(f, subtitleFiles));

    const audios: IFileNode[] = fileTree.files
      ?.filter(isAudioFile)
      .map((f) => transformFileDtoToFile(f, subtitleFiles));
    const children = transformDtoTreeToFileTree(fileTree.children);

    const result: IFileTree = {
      id: fileTree.id,
      name: fileTree.name,
      thumbnailId: fileTree.thumbnailId || undefined,
      images,
      children,
      audios,
      videos,
    };

    return result;
  });

  return fileTrees;
}

function transformFileDtoToFile(
  dto: IFileDto,
  subtitleFiles: ISubtitleFileDto[],
): IFileNode {
  const result: IFileNode = {
    id: dto.id,
    name: dto.name,
    fileType: dto.fileType,
  };

  if (isAudioFile(dto)) {
    const matchingSubtitleFile = subtitleFiles.find(
      (f) => f.audioFileId === dto.id,
    );
    if (matchingSubtitleFile) {
      result.subtitleFileId = matchingSubtitleFile.id;
    }
  }

  return result;
}
