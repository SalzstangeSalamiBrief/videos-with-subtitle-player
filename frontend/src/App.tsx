import { RouterProvider, createRouter } from '@tanstack/react-router';
import { routeTree } from './routeTree.gen';
import ErrorBoundary from '$sharedComponents/errorBoundary/ErrorBoundary';
import { NotFoundPage } from '$sharedComponents/notFoundPage/NotFoundPage';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

// eslint-disable-next-line react-refresh/only-export-components
export const queryClient = new QueryClient();

const router = createRouter({
  routeTree,
  defaultNotFoundComponent: NotFoundPage,
  context: {
    queryClient,
  },
  // TODO MAYBE REMOVE BOOTH PROPERTIES
  defaultPreload: 'intent',
  defaultPreloadStaleTime: 0,
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
// Register things for typesafety
declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router;
  }
}
