import { IFileItemDto } from './fileItemDto';

export interface IFileTreeDto {
  audioFiles: { [key: string]: IFileItemDto[] };
  children: { [key: string]: IFileTreeDto };
}
