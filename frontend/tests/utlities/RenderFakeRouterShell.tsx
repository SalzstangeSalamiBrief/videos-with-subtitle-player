import {
  createRootRoute,
  createRoute,
  createRouter,
  RouterProvider,
} from '@tanstack/react-router';
import { render } from '@testing-library/react';

// reference: https://github.com/TanStack/router/blob/main/packages/react-router/tests/routeContext.test.tsx
export function RenderFakeRouterShell(component: () => React.JSX.Element) {
  const rootRoute = createRootRoute();

  const indexRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: '/folders/$folderId/files/$fileId',
    component,
  });

  const routeTree = rootRoute.addChildren([indexRoute]);
  const router = createRouter({ routeTree }) as any;

  render(<RouterProvider router={router} />);
}
