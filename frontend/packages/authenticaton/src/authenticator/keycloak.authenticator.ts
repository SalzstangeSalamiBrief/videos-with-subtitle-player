import type { Maybe } from '@videos-with-subtitle-player/core';
import Keycloak from 'keycloak-js';
import type { IAuthenticator, IUserData } from './authenticator.model';

export interface IKeycloakAuthenticatorParams {
  url: Maybe<string>;
  realm: Maybe<string>;
  clientId: Maybe<string>;
  adminRoleName: Maybe<string>;
}

export function KeycloakAuthenticator(
  input: Maybe<IKeycloakAuthenticatorParams>,
): IAuthenticator {
  if (!input) {
    throw new Error('Keycloak authenticator configuration is required');
  }

  const { clientId, realm, url, adminRoleName } = input;

  if (!url) {
    throw new Error("The Keycloak configuration is missing the property 'url'");
  }

  if (!realm) {
    throw new Error(
      "The Keycloak configuration is missing the property 'realm'",
    );
  }

  if (!clientId) {
    throw new Error(
      "The Keycloak configuration is missing the property 'clientId'",
    );
  }

  if (!adminRoleName) {
    throw new Error(
      "The Keycloak configuration is missing the property 'adminRoleName'",
    );
  }

  const keycloak = new Keycloak({
    url,
    realm,
    clientId,
  });

  async function signIn(): Promise<void> {
    if (keycloak.didInitialize) {
      return;
    }

    const authenticated = await keycloak.init({
      onLoad: 'login-required',
      checkLoginIframe: false,
      flow: 'standard',
      enableLogging: import.meta.env.DEV,
      pkceMethod: 'S256',
      // TODO CHECK REFRESH
    });

    if (!authenticated) {
      throw new Error('Could not authenticate');
    }
  }

  async function signOut(): Promise<void> {
    if (!keycloak?.didInitialize) {
      return;
    }

    await keycloak.logout();
  }

  function getIsSignedIn(): boolean {
    return keycloak.authenticated ?? false;
  }

  function getAuthHeader(): '' | `Bearer ${string}` {
    if (!keycloak?.didInitialize) {
      return '';
    }

    return `Bearer ${keycloak.token}`;
  }

  function getIsAdmin(): boolean {
    if (!adminRoleName) {
      return false;
    }

    return keycloak.hasRealmRole(adminRoleName);
  }

  function getUserData(): IUserData {
    return {
      userName: keycloak.tokenParsed?.name,
      email: keycloak.tokenParsed?.email,
    };
  }

  return {
    signOut,
    signIn,
    getIsSignedIn,
    getAuthHeader,
    getIsAdmin,
    getUserData,
  };
}
