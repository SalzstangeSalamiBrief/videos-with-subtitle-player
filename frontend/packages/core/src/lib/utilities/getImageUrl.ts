// TODO DOES THIS WORK WITH THE MONO  SETUP AND PACKAGES?
const baseUrl = import.meta.env.VITE_BASE_URL;

export function getImageUrlForId(id: string | undefined): string | undefined {
  if (!id) {
    return undefined;
  }

  // TODO OUTSOURCE PATHS INTO CONSTANTS OR GENERATE CODE?
  return `${baseUrl}/api/file/discrete/${id}`;
}
