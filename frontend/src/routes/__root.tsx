import { Outlet, createRootRoute } from '@tanstack/react-router';
import { FileTreeContextWrapper } from '$contexts/FileTreeContextWrapper';
import { getFileTreeQueryOptions } from '$queries/getFileTree/getFileTreeQueryOptions';
import { queryClient } from '../App';
import { useSuspenseQuery } from '@tanstack/react-query';
import { LoadingSpinner } from '$sharedComponents/loadingSpinner/LoadingSpinner';

export const Route = createRootRoute({
  component: Root,
  meta: getPageMetadata,
  loader: () => queryClient.ensureQueryData(getFileTreeQueryOptions),
  // TODO ERROR COMPONENT
});

function Root() {
  const {
    data: { fileGroups, fileTrees },
    isLoading,
  } = useSuspenseQuery(getFileTreeQueryOptions);

  if (isLoading) {
    return (
      <div style={{ paddingTop: '1.5rem' }}>
        <LoadingSpinner text="Loading audio files..." />
      </div>
    );
  }

  return (
    <FileTreeContextWrapper input={{ fileTrees, fileGroups }}>
      <div className="grid gap-4">
        <main className="p-4 overflow-y-auto max-h-[100lvh]">
          <Outlet />
        </main>
      </div>
    </FileTreeContextWrapper>
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
