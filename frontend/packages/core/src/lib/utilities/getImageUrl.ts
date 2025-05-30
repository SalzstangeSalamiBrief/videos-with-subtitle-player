export function getImageUrlForId(
  baseUrl: string,
  id: string | undefined,
): string | undefined {
  if (!id) {
    return undefined;
  }

  // TODO OUTSOURCE PATHS INTO CONSTANTS OR GENERATE CODE?
  return `${baseUrl}/api/file/discrete/${id}`;
}
