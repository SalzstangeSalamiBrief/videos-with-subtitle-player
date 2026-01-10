import type { IFileDto } from './dtos/fileDtos';

export interface IFileTree {
  name: string;
  id: string;
  thumbnailId: Maybe<string>;
  lowQualityThumbnailId: Maybe<string>;
  children: IFileTree[];
  images: IFileNode[];
  audios: IFileNode[];
  videos: IFileNode[];
}

export interface IFileNode extends IFileDto {
  subtitleFileId?: string;
}
