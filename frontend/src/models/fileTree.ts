import { IFileDto } from './dtos/fileDtos';

export interface IFileTree {
  name: string;
  id: string;
  thumbnailId: string | undefined;
  files: IFileNode[];
  children: IFileTree[];
}

export interface IFileNode extends IFileDto {
  subtitleFileId?: string;
}
