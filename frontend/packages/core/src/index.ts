// css
import './index.css';

// <---- enums ---->
export { FileType } from './enums/FileType';

// <---- models---->
export {
  type IFileDto,
  type ISubtitleFileDto,
} from './models/fileTree/dtos/fileDtos';
export {
  type IFileTreeDto,
  type PossibleFilesDto,
} from './models/fileTree/dtos/fileTreeDto';
export { type IFileNode, type IFileTree } from './models/fileTree/fileTree';
export type { Maybe } from './models/maybe';

// <---- lib ---->
export {
  isAudioFile,
  isImageFile,
  isSubtitleFile,
  isVideoFile,
} from './lib/type-predicates/file-type-predicates';
export { getFolderFromFileTree } from './lib/utilities/getFoldersFromFileTree';
export { getFoldersInActiveTree } from './lib/utilities/getFoldersInActiveTree';
export { ImageGetter, type IImageGetter } from './lib/utilities/imageGetter';

// <---- queries ---->
export {
  getFileTreeSelect,
  type IGetFileTreeSelectReturn,
} from './queries/getFileTree/getFileTree.select';
export { getFileTreeQuery } from './queries/getFileTree/getFileTreeQuery.query';
