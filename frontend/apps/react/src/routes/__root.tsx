import { getFileTreeQuery } from '$queries/getFileTree/getFileTreeQueryQuery';
import { getFileTreeSelect } from '$queries/getFileTree/getFileTreeSelect';
import { ErrorComponent } from '$sharedComponents/errorComponent/ErrorComponent';
import { createRootRoute, Outlet } from '@tanstack/react-router';
import { DefaultPendingComponent } from '../App';

export const Route = createRootRoute({
  component: Root,
  loader: async () => {
    const responseData = await getFileTreeQuery();
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
    <main className="h-full max-h-[100lvh] overflow-y-auto p-4">
      <Outlet />
    </main>
  );
}
