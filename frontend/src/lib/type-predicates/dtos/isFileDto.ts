import type { FileDto } from 'src/client/generated';

export function isFileDto(input: unknown): input is FileDto {
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
