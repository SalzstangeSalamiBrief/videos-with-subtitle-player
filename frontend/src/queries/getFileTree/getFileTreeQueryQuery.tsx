import { IFileTreeDto } from '$models/dtos/fileTreeDto';

const baseUrl = import.meta.env.VITE_BASE_URL || '';
const path = '/api/file-tree';

const url = baseUrl + path;

export async function getFileTreeQuery(): Promise<IFileTreeDto[]> {
  const response = await fetch(url);
  const json: IFileTreeDto[] = await response.json();
  return json;
}
