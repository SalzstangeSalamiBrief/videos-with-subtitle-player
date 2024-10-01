import { IFileDto } from './dtos/fileDtos';

export interface IFileTree {
  name: string;
  id: string;
  thumbnailId: string | undefined;
  children: IFileTree[];
  continuousFiles: IFileNode[];
  images: IFileNode[];
}

export interface IFileNode extends IFileDto {
  subtitleFileId?: string;
}
