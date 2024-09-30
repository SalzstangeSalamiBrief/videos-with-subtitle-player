import { IFileDto, ISubtitleFileDto } from './fileDtos';

export type PossibleFilesDto = IFileDto | ISubtitleFileDto;

export interface IFileTreeDto {
  name: string;
  id: string;
  thumbnailId: string;
  files: PossibleFilesDto[];
  children: IFileTreeDto[];
}
