export interface IImageGetter {
  getImageUrlForId: (id: Maybe<string>) => Maybe<string>;
}

export function ImageGetter(baseUrl: string): IImageGetter {
  const backendUrl = `${baseUrl}/api/file/discrete`;

  function getImageUrlForId(id: Maybe<string>): Maybe<string> {
    if (!id) {
      return undefined;
    }

    return `${backendUrl}/${id}`;
  }

  return { getImageUrlForId };
}
