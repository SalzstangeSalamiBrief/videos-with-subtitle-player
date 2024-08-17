import { Outlet, createRootRoute } from '@tanstack/react-router';
import { Navigation } from '$features/navigation/Navigation';
import { FileTreeContextWrapper } from '$contexts/FileTreeContextWrapper';

export interface RootSearchParams {
  activeTab: number | undefined;
}

export const Route = createRootRoute({
  component: Root,
  validateSearch(input: Record<string, unknown>): RootSearchParams {
    const result: RootSearchParams = { activeTab: undefined };
    if (!input.activeTab) {
      return result;
    }

    const activeTab = Number(input.activeTab);
    if (!Number.isNaN(activeTab)) {
      result.activeTab = activeTab;
    }

    return result;
  },
});

function Root() {
  return (
    <FileTreeContextWrapper>
      <div className="flex">
        <div className="grid min-h-[100lvh] max-h-[100lvh] overflow-y-auto">
          <Navigation />
        </div>
        <div className="grid gap-4">
          <main
            style={{
              maxHeight: '100lvh',
              padding: '1rem',
              overflowY: 'auto',
            }}
          >
            <Outlet />
          </main>
        </div>
      </div>
    </FileTreeContextWrapper>
  );
}
