import type { ISubtitleFileDto } from '$models/fileTree/dtos/fileDtos';
import { isSubtitleFileDto } from './isSubtitleFileDto';

export function isSubtitleFileDtoArray(
  input: unknown,
): input is ISubtitleFileDto[] {
  if (!input || !Array.isArray(input)) {
    return false;
  }

  if (!input.length) {
    return true;
  }

  return input.every(isSubtitleFileDto);
}
