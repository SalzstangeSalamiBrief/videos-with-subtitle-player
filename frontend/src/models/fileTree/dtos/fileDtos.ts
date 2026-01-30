import type { FileDto } from 'src/client/generated/fileDto';

export interface ISubtitleFileDto extends FileDto {
  audioFileId: string;
}
