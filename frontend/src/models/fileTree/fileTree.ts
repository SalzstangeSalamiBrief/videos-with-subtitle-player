import type { FileDto } from 'src/client/generated/fileDto';

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

export interface IFileNode extends FileDto {
  subtitleFileId?: string;
}
