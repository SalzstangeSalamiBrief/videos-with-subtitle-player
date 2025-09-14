/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_BASE_URL: string;
  readonly VITE_APP_TITLE: string;
  readonly VITE_KEYCLOAK_URL: Maybe<string>;
  readonly VITE_KEYCLOAK_REALM: Maybe<string>;
  readonly VITE_KEYCLOAK_CLIENT_ID: Maybe<string>;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
