import type { ISubtitleFileDto } from '$models/fileTree/dtos/fileDtos';
import { isFileDto } from './isFileDto';

export function isSubtitleFileDto(input: unknown): input is ISubtitleFileDto {
  if (!input || typeof input !== 'object') {
    return false;
  }

  const isFile = isFileDto(input);
  const hasAudioFileId =
    'audioFileId' in input &&
    typeof input.audioFileId === 'string' &&
    Boolean(input.audioFileId);

  return isFile && hasAudioFileId;
}
