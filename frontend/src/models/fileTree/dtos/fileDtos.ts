import type { FileType } from '$enums/FileType';

export interface IFileDto {
  id: string;
  name: string;
  fileType: FileType;
}

export interface ISubtitleFileDto extends IFileDto {
  audioFileId: string;
}

export function isFileDto(input: unknown): input is IFileDto {
  if (!input || typeof input !== 'object') {
    return false;
  }

  const hasId =
    'id' in input && typeof input.id === 'string' && Boolean(input.id);
  const hasName =
    'name' in input && typeof input.name === 'string' && Boolean(input.name);
  const hasFileType =
    'fileType' in input &&
    typeof input.fileType === 'string' &&
    Boolean(input.fileType);
  return hasId && hasName && hasFileType;
}

export function isFileDtoArray(input: unknown): input is IFileDto[] {
  if (!input || !Array.isArray(input)) {
    return false;
  }

  if (!input.length) {
    return true;
  }

  return input.every(isFileDto);
}

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
