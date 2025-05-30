import type { IFileTreeDto } from '../../models/fileTree/dtos/fileTreeDto';

// TODO OUTSOURCE PATHS INTO CONSTANTS OR GENERATE CODE?

const path = '/api/file-tree';

export async function getFileTreeQuery(
  baseUrl: string,
): Promise<IFileTreeDto[]> {
  const url = baseUrl + path;

  const response = await fetch(url);
  const json: IFileTreeDto[] = await response.json();
  return json;
}
