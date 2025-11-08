// css
import './index.css';

// <---- enums ---->
export { FileType } from './enums/FileType';

// <---- models---->
export {
  isFileDto,
  isFileDtoArray,
  isSubtitleFileDto,
  isSubtitleFileDtoArray,
  type IFileDto,
  type ISubtitleFileDto,
} from './models/fileTree/dtos/fileDtos';
export {
  isFileTreeDto,
  isFileTreeDtoArray,
  type IFileTreeDto,
  type PossibleFilesDto,
} from './models/fileTree/dtos/fileTreeDto';
export { type IFileNode, type IFileTree } from './models/fileTree/fileTree';
export type { Maybe } from './models/maybe';

export { ApiError, type IApiError } from './models/ApiError';

// <---- lib ---->
export {
  isAudioFile,
  isImageFile,
  isSubtitleFile,
  isVideoFile,
} from './lib/type-predicates/file-type-predicates';
export { getCurrentNodeWithSiblings } from './lib/utilities/getCurrentFileNodeWithSibling';
export { getFolderFromFileTree } from './lib/utilities/getFoldersFromFileTree';
export { getFoldersInActiveTree } from './lib/utilities/getFoldersInActiveTree';
export { ImageGetter, type IImageGetter } from './lib/utilities/imageGetter';

// <---- queries ---->
export {
  getFileTreeSelect,
  type IGetFileTreeSelectReturn,
} from './queries/getFileTree/getFileTree.select';
export { getFileTreeQuery } from './queries/getFileTree/getFileTreeQuery.query';
