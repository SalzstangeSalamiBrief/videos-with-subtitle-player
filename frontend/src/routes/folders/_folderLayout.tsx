import { Breadcrumbs } from '$features/breadcrumbs/Breadcrumbs';
import { createFileRoute, Outlet } from '@tanstack/react-router';

export const Route = createFileRoute('/folders/_folderLayout')({
  component: FolderLayoutComponent,
});

function FolderLayoutComponent() {
  return (
    <div className="grid gap-4">
      <Breadcrumbs />
      <Outlet />
    </div>
  );
}
