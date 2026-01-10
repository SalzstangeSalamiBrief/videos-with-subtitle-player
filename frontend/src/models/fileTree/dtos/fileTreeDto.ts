import { isFileDto, isSubtitleFileDto } from './fileDtos';
import type { IFileDto, ISubtitleFileDto } from './fileDtos';

export type PossibleFilesDto = IFileDto | ISubtitleFileDto;

export interface IFileTreeDto {
  name: string;
  id: string;
  thumbnailId?: string;
  lowQualityThumbnailId?: string;
  files: PossibleFilesDto[];
  children: IFileTreeDto[];
}

export function isFileTreeDto(input: unknown): input is IFileTreeDto {
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

export function isFileTreeDtoArray(input: unknown): input is IFileTreeDto[] {
  if (!input || !Array.isArray(input)) {
    return false;
  }

  if (!input.length) {
    return true;
  }

  return input.every(isFileTreeDto);
}
