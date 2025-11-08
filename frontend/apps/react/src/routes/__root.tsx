import { ErrorComponent } from '$sharedComponents/errorComponent/ErrorComponent';
import { createRootRoute, Outlet } from '@tanstack/react-router';
import {
  getFileTreeQuery,
  getFileTreeSelect,
} from '@videos-with-subtitle-player/core';
import { DefaultPendingComponent } from '../App';

const baseUrl = import.meta.env.VITE_BASE_URL || '';
export const Route = createRootRoute({
  component: Root,
  loader: async () => {
    const responseData = await getFileTreeQuery(baseUrl);
    if (responseData.error) {
      throw responseData.error;
    }

    const result = getFileTreeSelect(responseData.data!);
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
