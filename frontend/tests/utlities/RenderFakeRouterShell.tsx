import {
  createRootRoute,
  createRoute,
  createRouter,
  RouterProvider,
} from '@tanstack/react-router';
import { render } from '@testing-library/react';

export function RenderFakeRouterShell(component: () => React.JSX.Element) {
  const rootRoute = createRootRoute({
    component,
  });

  const indexRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: '/folders/$folderId/files/$fileId',
    component,
  });

  const routeTree = rootRoute.addChildren([indexRoute]);
  const router = createRouter({ routeTree }) as any;

  render(<RouterProvider router={router} />);
}
