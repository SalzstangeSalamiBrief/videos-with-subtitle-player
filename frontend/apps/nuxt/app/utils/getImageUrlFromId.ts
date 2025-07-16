import { ImageGetter } from '@videos-with-subtitle-player/core';

const baseUrl = import.meta.env.VITE_BASE_URL || '';

if (baseUrl === null) {
  throw new Error(
    'The environment variable VITE_BASE_URL is missing. Please provide it',
  );
}

export const imageHandler = ImageGetter(baseUrl);
