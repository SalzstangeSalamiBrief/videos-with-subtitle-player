import { IFileItemDto } from './fileItemDto';

export interface IAudioFileDto {
  name: string;
  subtitleFile: IFileItemDto;
  audioFile: IFileItemDto;
}
