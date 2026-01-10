import { Outlet, createRootRoute } from '@tanstack/react-router';
import { ErrorComponent } from '$sharedComponents/errorComponent/ErrorComponent';

import { ApiError } from '$models/ApiError';
import { getFileTreeSelect } from '$queries/getFileTree/getFileTree.select';
import { getFileTreeQuery } from '$queries/getFileTree/getFileTreeQuery.query';
import { LoadingSpinner } from '$sharedComponents/loadingSpinner/LoadingSpinner';

const baseUrl = import.meta.env.VITE_BASE_URL || '';
export const Route = createRootRoute({
  component: Root,
  loader: async () => {
    const responseData = await getFileTreeQuery(baseUrl);
    if (responseData.error) {
      throw responseData.error;
    }

    if (!responseData.data) {
      throw new ApiError({
        detail: 'No data received from server',
        status: 500,
        title: 'Data Error',
        type: 'about:blank',
      });
    }

    const result = getFileTreeSelect(responseData.data);
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

function DefaultPendingComponent() {
  return (
    <div style={{ paddingTop: '1.5rem' }}>
      <LoadingSpinner text="Loading data..." />
    </div>
  );
}
