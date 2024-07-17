import { FileType } from '$enums/FileType';

export interface IFileDto {
  id: string;
  name: string;
  fileType: FileType;
}

export interface ISubtitleFileDto extends IFileDto {
  audioFileId: string;
}
