import { IAudioFileDto } from './audioFileDto';

export interface IFileTreeDto {
  name: string;
  id: string;
  audioFiles: IAudioFileDto[];
  children: IFileTreeDto[];
}
