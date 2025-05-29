const baseUrl = import.meta.env.VITE_BASE_URL;

export function getImageUrlForId(id: string | undefined): string | undefined {
  if (!id) {
    return undefined;
  }

  return `${baseUrl}/api/file/discrete/${id}`;
}
