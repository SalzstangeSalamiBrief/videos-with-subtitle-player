import { Outlet, createRootRoute } from '@tanstack/react-router';
import { LoadingSpinner } from '$sharedComponents/loadingSpinner/LoadingSpinner';
import { ErrorComponent } from '$sharedComponents/errorComponent/ErrorComponent';
import { getFileTreeQuery } from '$queries/getFileTree/getFileTreeQuery';
import { getFileTreeSelect } from '$queries/getFileTree/getFileTreeSelect';

export const Route = createRootRoute({
  component: Root,
  meta: getPageMetadata,
  loader: async () => {
    const responseData = await getFileTreeQuery();
    const result = getFileTreeSelect(responseData);
    return result;
  },
  errorComponent: ErrorComponent,
  pendingComponent: () => (
    <div style={{ paddingTop: '1.5rem' }}>
      <LoadingSpinner text="Loading data..." />
    </div>
  ),
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
