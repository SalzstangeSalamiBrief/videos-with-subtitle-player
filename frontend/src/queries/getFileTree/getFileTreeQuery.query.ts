// TODO OUTSOURCE PATHS INTO CONSTANTS OR GENERATE CODE?

import type { IApiResponse } from '$models/ApiResponse';
import type {IFileTreeDto} from '$models/fileTree/dtos/fileTreeDto';
import type {IApiError} from '$models/ApiError';
import { ApiError  } from '$models/ApiError';
import {
  
  isFileTreeDtoArray
} from '$models/fileTree/dtos/fileTreeDto';

const path = '/api/file-tree';

export async function getFileTreeQuery(
  baseUrl: string,
): Promise<IApiResponse<IFileTreeDto[]>> {
  try {
    const url = baseUrl + path;

    const response = await fetch(url);

    const json: IFileTreeDto[] | IApiError = await response.json();
    if (!response.ok) {
      if (ApiError.isApiError(json)) {
        return { error: new ApiError(json) };
      }

      throw new Error('Unknown error occurred');
    }

    if (isFileTreeDtoArray(json)) {
      return { data: json };
    }

    throw new Error('Invalid data format received');
  } catch (error) {
    console.error('Error in getFileTreeQuery:', error);
    throw new Error('Failed to fetch file tree data');
  }
}
