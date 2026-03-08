import type { FileTreeDto } from 'src/client/generated';
import { isFileTreeDto } from './isFileTreeDto';

export function isFileTreeDtoArray(input: unknown): input is FileTreeDto[] {
  if (!input || !Array.isArray(input)) {
    return false;
  }

  if (!input.length) {
    return true;
  }

  return input.every(isFileTreeDto);
}
