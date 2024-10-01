import { FileType } from '$enums/FileType';
import { ISubtitleFileDto, IFileDto } from '$models/dtos/fileDtos';
import { PossibleFilesDto } from '$models/dtos/fileTreeDto';
import { IFileNode } from '$models/fileTree';

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
