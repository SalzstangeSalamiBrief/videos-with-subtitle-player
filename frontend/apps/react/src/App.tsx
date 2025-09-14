import ErrorBoundary from '$sharedComponents/errorBoundary/ErrorBoundary';
import { LoadingSpinner } from '$sharedComponents/loadingSpinner/LoadingSpinner';
import { NotFoundPage } from '$sharedComponents/notFoundPage/NotFoundPage';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { RouterProvider, createRouter } from '@tanstack/react-router';
import {
  AuthenticatorFactory,
  type AuthenticatorType,
} from '@videos-with-subtitle-player/authentication';
import { routeTree } from './routeTree.gen';

const authenticator = AuthenticatorFactory.createAuthenticator({
  authenticatorType: import.meta.env
    .VITE_AUTHENTICATOR_PROVIDER as AuthenticatorType,
  config: {
    keycloak: {
      adminRoleName: import.meta.env.VITE_ADMIN_ROLE_NAME,
      clientId: import.meta.env.VITE_KEYCLOAK_CLIENT_ID,
      realm: import.meta.env.VITE_KEYCLOAK_REALM,
      url: import.meta.env.VITE_KEYCLOAK_URL,
    },
  },
});

// eslint-disable-next-line react-refresh/only-export-components
export const queryClient = new QueryClient();
const router = createRouter({
  routeTree,
  defaultNotFoundComponent: NotFoundPage,
  context: {
    queryClient,
    authenticator,
  },
});

export function App() {
  return (
    <ErrorBoundary>
      <QueryClientProvider client={queryClient}>
        <RouterProvider router={router} />
      </QueryClientProvider>
    </ErrorBoundary>
  );
}

export function DefaultPendingComponent() {
  return (
    <div style={{ paddingTop: '1.5rem' }}>
      <LoadingSpinner text="Loading data..." />
    </div>
  );
}
// Register things for typesafety
declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router;
  }
}
