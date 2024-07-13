import { RouterProvider, createRouter } from '@tanstack/react-router';
import { routeTree } from './routeTree.gen';
import ErrorBoundary from '$sharedComponents/errorBoundary/ErrorBoundary';
import { NotFoundPage } from '$sharedComponents/notFoundPage/NotFoundPage';

const router = createRouter({
  routeTree,
  defaultNotFoundComponent: NotFoundPage,
});

export function App() {
  return (
    <ErrorBoundary>
      <RouterProvider router={router} />
    </ErrorBoundary>
  );
}
