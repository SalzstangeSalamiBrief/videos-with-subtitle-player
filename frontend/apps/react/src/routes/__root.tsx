import { ErrorComponent } from '$sharedComponents/errorComponent/ErrorComponent';
import type { QueryClient } from '@tanstack/react-query';
import { createRootRouteWithContext, Outlet } from '@tanstack/react-router';
import type { IAuthenticator } from '@videos-with-subtitle-player/authentication';
import {
  getFileTreeQuery,
  getFileTreeSelect,
} from '@videos-with-subtitle-player/core';
import { DefaultPendingComponent } from '../App';

interface IRouterContext {
  queryClient: QueryClient;
  authenticator: IAuthenticator;
}

const baseUrl = import.meta.env.VITE_BASE_URL || '';
export const Route = createRootRouteWithContext<IRouterContext>()({
  component: Root,
  loader: async ({ context: { authenticator } }) => {
    if (!authenticator.getIsSignedIn()) {
      await authenticator.signIn();
    }

    const responseData = await getFileTreeQuery(baseUrl);
    const result = getFileTreeSelect(responseData);
    return result;
  },
  shouldReload: false,
  errorComponent: ErrorComponent,
  pendingComponent: DefaultPendingComponent,
  wrapInSuspense: true,
});

function Root() {
  return (
    <main className="h-full max-h-lvh overflow-y-auto p-4">
      <Outlet />
    </main>
  );
}
