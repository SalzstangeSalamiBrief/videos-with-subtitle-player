import type { FileDto } from 'src/client/generated/fileDto';
import type { ISubtitleFileDto } from './fileDtos';

export type PossibleFilesDto = FileDto | ISubtitleFileDto;
