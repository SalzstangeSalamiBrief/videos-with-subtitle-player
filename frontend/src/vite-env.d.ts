/// <reference types="vite/client" />

type Maybe<T> = T | undefined;

interface ImportMetaEnv {
  readonly VITE_BASE_URL: string;
  readonly VITE_APP_TITLE: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
