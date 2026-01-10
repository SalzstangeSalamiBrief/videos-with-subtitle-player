import { ImageGetter } from '$lib/utilities/imageGetter';

const baseUrl = import.meta.env.VITE_BASE_URL || '';

if (!baseUrl) {
  throw new Error(
    'The environment variable VITE_BASE_URL is missing. Please provide it',
  );
}

export const imageHandler = ImageGetter(baseUrl);
