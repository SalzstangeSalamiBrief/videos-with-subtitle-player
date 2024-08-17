import { Breadcrumbs } from '$features/breadcrumbs/Breadcrumbs';
import { createFileRoute, Outlet } from '@tanstack/react-router';

export interface RootSearchParams {
  activeTab?: number;
}

export const Route = createFileRoute('/folders/_folderLayout')({
  component: FolderLayoutComponent,
  validateSearch: searchParamValidator,
});

function FolderLayoutComponent() {
  return (
    <div className="grid gap-4">
      <Breadcrumbs />
      <Outlet />
    </div>
  );
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
