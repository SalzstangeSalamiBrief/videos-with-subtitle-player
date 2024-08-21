import { Outlet, createRootRoute } from '@tanstack/react-router';
import { LoadingSpinner } from '$sharedComponents/loadingSpinner/LoadingSpinner';
import { ErrorComponent } from '$sharedComponents/errorComponent/ErrorComponent';
import {
  getFileTreeSelect,
  IGetFileTreeSelectReturn,
} from '$queries/getFileTree/getFileTreeSelect';
import { getFileTreeQuery } from '$queries/getFileTree/getFileTreeQueryQuery';

export const Route = createRootRoute({
  component: Root,
  meta: getPageMetadata,
  loader: async () => {
    // TODO DOES NOT SHOW PENDING COMPONENT
    return await new Promise<IGetFileTreeSelectReturn>((resolve) => {
      setTimeout(async () => {
        const responseData = await getFileTreeQuery();
        const result = getFileTreeSelect(responseData);
        resolve(result);
        // return result;
      }, 1500);
    });
  },
  shouldReload: false,
  errorComponent: ErrorComponent,
  pendingComponent: RootLoader,
});

function Root() {
  return (
    <div className="grid gap-4">
      <main className="p-4 overflow-y-auto max-h-[100lvh]">
        <Outlet />
      </main>
    </div>
  );
}

function RootLoader() {
  return (
    <div style={{ paddingTop: '1.5rem' }}>
      <LoadingSpinner text="Loading data..." />
    </div>
  );
}
function getPageMetadata() {
  return [
    {
      name: 'viewport',
      content: 'width=device-width, initial-scale=1',
    },
    { title: 'Video with subtitle player' },
    {
      name: 'Description',
      content:
        'Stream audio and video files with subtitle support, offering seamless playback for your media content.',
    },
    { lang: 'en' },
    {
      charSet: 'utf-8',
    },
  ];
}
