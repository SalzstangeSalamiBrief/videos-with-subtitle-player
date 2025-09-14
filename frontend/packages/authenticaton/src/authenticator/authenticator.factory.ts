import type { Maybe } from '@videos-with-subtitle-player/core';
import { AuthenticatorType, type IAuthenticator } from './authenticator.model';
import {
  type IKeycloakAuthenticatorParams,
  KeycloakAuthenticator,
} from './keycloak.authenticator';

export interface IAuthenticatorFactoryInput {
  config: {
    keycloak?: Maybe<IKeycloakAuthenticatorParams>;
  };
  authenticatorType: Maybe<AuthenticatorType>;
}

interface IAuthenticatorFactory {
  createAuthenticator: (input: IAuthenticatorFactoryInput) => IAuthenticator;
}

function createAuthenticator(
  input: IAuthenticatorFactoryInput,
): IAuthenticator {
  const { config, authenticatorType } = input;

  switch (authenticatorType) {
    case AuthenticatorType.Keycloak:
      return KeycloakAuthenticator(config.keycloak);
    default:
      throw new Error(`Unsupported authenticator type: ${authenticatorType}`);
  }
}

export const AuthenticatorFactory: IAuthenticatorFactory = {
  createAuthenticator,
};
