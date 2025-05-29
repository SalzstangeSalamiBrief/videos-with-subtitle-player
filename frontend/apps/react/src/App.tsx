import ErrorBoundary from '$sharedComponents/errorBoundary/ErrorBoundary';
import { LoadingSpinner } from '$sharedComponents/loadingSpinner/LoadingSpinner';
import { NotFoundPage } from '$sharedComponents/notFoundPage/NotFoundPage';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { RouterProvider, createRouter } from '@tanstack/react-router';
import { routeTree } from './routeTree.gen';

export const queryClient = new QueryClient();
const router = createRouter({
  routeTree,
  defaultNotFoundComponent: NotFoundPage,
  context: {
    queryClient,
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
