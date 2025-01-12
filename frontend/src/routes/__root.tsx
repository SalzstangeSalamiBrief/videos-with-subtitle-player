import { getFileTreeQuery } from '$queries/getFileTree/getFileTreeQueryQuery';
import { getFileTreeSelect } from '$queries/getFileTree/getFileTreeSelect';
import { ErrorComponent } from '$sharedComponents/errorComponent/ErrorComponent';
import { Outlet, createRootRoute } from '@tanstack/react-router';
import { DefaultPendingComponent } from '../App';

export const Route = createRootRoute({
  component: Root,
  meta: getPageMetadata,
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
