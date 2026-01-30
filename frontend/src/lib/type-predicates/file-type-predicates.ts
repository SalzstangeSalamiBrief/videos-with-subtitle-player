import type { PossibleFilesDto } from '$models/fileTree/dtos/fileTreeDto';
import type { IFileNode } from '$models/fileTree/fileTree';
import type { FileDto } from 'src/client/generated/fileDto';
import { FileType } from 'src/client/generated/fileType';
import type { ISubtitleFileDto } from '../../models/fileTree/dtos/fileDtos';

export function isSubtitleFile(
  file: PossibleFilesDto | IFileNode,
): file is ISubtitleFileDto {
  return file.fileType === FileType.Subtitle;
}

export function isAudioFile(
  file: PossibleFilesDto | IFileNode,
): file is FileDto {
  return file.fileType === FileType.Audio;
}

export function isVideoFile(
  file: PossibleFilesDto | IFileNode,
): file is FileDto {
  return file.fileType === FileType.Video;
}

export function isImageFile(
  file: PossibleFilesDto | IFileNode,
): file is FileDto {
  return file.fileType === FileType.Image;
}
