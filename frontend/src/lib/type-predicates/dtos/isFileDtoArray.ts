import type { FileDto } from 'src/client/generated';
import { isFileDto } from './isFileDto';

export function isFileDtoArray(input: unknown): input is FileDto[] {
  if (!input || !Array.isArray(input)) {
    return false;
  }

  if (!input.length) {
    return true;
  }

  return input.every(isFileDto);
}
