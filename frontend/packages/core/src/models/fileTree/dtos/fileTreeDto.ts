import type { IFileDto, ISubtitleFileDto } from './fileDtos';

export type PossibleFilesDto = IFileDto | ISubtitleFileDto;

export interface IFileTreeDto {
  name: string;
  id: string;
  thumbnailId: string;
  lowQualityThumbnailId: string;
  files: PossibleFilesDto[];
  children: IFileTreeDto[];
}
