import type { Maybe } from '@videos-with-subtitle-player/core';

export interface IUserData {
  userName: Maybe<string>;
  email: Maybe<string>;
}

export interface IAuthenticator {
  signOut: () => Promise<void>;
  signIn: () => Promise<void>;
  getIsSignedIn: () => boolean;
  getIsAdmin: () => boolean;
  getAuthHeader: () => '' | `Bearer ${string}`;
  getUserData: () => Maybe<IUserData>;
}

export interface IAuthenticatorBaseInput {
  url: string;
  realm: string;
  clientId: string;
  adminRoleName: string;
}

export enum AuthenticatorType {
  Keycloak = 'keycloak',
}
