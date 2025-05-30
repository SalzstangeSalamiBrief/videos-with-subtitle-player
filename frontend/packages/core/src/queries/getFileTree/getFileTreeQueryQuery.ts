import type { IFileTreeDto } from '../../models/fileTree/dtos/fileTreeDto';

// TODO DOES THIS WORK WITH PACKAGES?
const baseUrl = import.meta.env.VITE_BASE_URL || '';
// TODO OUTSOURCE PATHS INTO CONSTANTS OR GENERATE CODE?

const path = '/api/file-tree';

const url = baseUrl + path;

export async function getFileTreeQuery(): Promise<IFileTreeDto[]> {
  const response = await fetch(url);
  const json: IFileTreeDto[] = await response.json();
  return json;
}
