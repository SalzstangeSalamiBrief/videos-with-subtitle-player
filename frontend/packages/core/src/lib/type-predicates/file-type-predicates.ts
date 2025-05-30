import { FileType } from '../../enums/FileType';
import type {
  IFileDto,
  ISubtitleFileDto,
} from '../../models/fileTree/dtos/fileDtos';
import type { PossibleFilesDto } from '../../models/fileTree/dtos/fileTreeDto';
import type { IFileNode } from '../../models/fileTree/fileTree';

export function isSubtitleFile(
  file: PossibleFilesDto | IFileNode,
): file is ISubtitleFileDto {
  return file.fileType === FileType.SUBTITLE;
}

export function isAudioFile(
  file: PossibleFilesDto | IFileNode,
): file is IFileDto {
  return file.fileType === FileType.AUDIO;
}

export function isVideoFile(
  file: PossibleFilesDto | IFileNode,
): file is IFileDto {
  return file.fileType === FileType.VIDEO;
}

export function isImageFile(
  file: PossibleFilesDto | IFileNode,
): file is IFileDto {
  return file.fileType === FileType.IMAGE;
}
