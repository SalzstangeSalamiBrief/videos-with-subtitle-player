import type { FileTreeDto } from 'src/client/generated';
import { isFileDto } from './isFileDto';
import { isSubtitleFileDto } from './isSubtitleFileDto';

export function isFileTreeDto(input: unknown): input is FileTreeDto {
  if (!input || typeof input !== 'object') {
    return false;
  }

  const hasName =
    'name' in input && typeof input.name === 'string' && Boolean(input.name);
  const hasId =
    'id' in input && typeof input.id === 'string' && Boolean(input.id);
  const hasFiles =
    'files' in input &&
    Array.isArray(input.files) &&
    input.files.every((file) => isFileDto(file) || isSubtitleFileDto(file));
  const hasChildren =
    'children' in input &&
    Array.isArray(input.children) &&
    input.children.every(isFileTreeDto);

  return hasName && hasId && hasFiles && hasChildren;
}
