import { IFileDto } from './dtos/fileDtos';

export interface IFileTree {
  name: string;
  id: string;
  thumbnailId: string | undefined;
  children: IFileTree[];
  images: IFileNode[];
  audios: IFileNode[];
  videos: IFileNode[];
}

export interface IFileNode extends IFileDto {
  subtitleFileId?: string;
}
