import { Outlet, createRootRoute } from '@tanstack/react-router';
import { Navigation } from '$sharedComponents/navigation/Navigation';
import { FileTreeContextWrapper } from '$contexts/FileTreeContextWrapper';

export const Route = createRootRoute({
  component: Root,
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
