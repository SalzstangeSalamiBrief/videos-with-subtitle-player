import { Outlet, createRootRoute } from '@tanstack/react-router';
import { FileTreeContextWrapper } from '$contexts/FileTreeContextWrapper';

export interface RootSearchParams {
  activeTab: number | undefined;
}

export const Route = createRootRoute({
  // TODO META TO ADD TAGS
  component: Root,
  validateSearch: searchParamValidator,
  meta: getPageMetadata,
});

function Root() {
  return (
    <FileTreeContextWrapper>
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

function searchParamValidator(
  input: Record<string, unknown>,
): RootSearchParams {
  const result: RootSearchParams = { activeTab: undefined };
  if (!input.activeTab) {
    return result;
  }

  const activeTab = Number(input.activeTab);
  if (!Number.isNaN(activeTab)) {
    result.activeTab = activeTab;
  }

  return result;
}
