import { Outlet, createFileRoute } from '@tanstack/react-router';
import { Breadcrumbs } from '$features/breadcrumbs/Breadcrumbs';

export interface IFolderLayoutSearchParams {
  activeTab?: number;
}

export const Route = createFileRoute('/folders/_folderLayout')({
  component: FolderLayoutComponent,
  validateSearch: searchParamValidator,
});

function FolderLayoutComponent() {
  return (
    <>
      <div className="flex h-full flex-col gap-4">
        <Breadcrumbs />
        <Outlet />
      </div>
    </>
  );
}

function searchParamValidator(
  input: Record<string, unknown>,
): IFolderLayoutSearchParams {
  const result: IFolderLayoutSearchParams = { activeTab: undefined };
  if (!input.activeTab) {
    return result;
  }

  const activeTab = Number(input.activeTab);
  if (!Number.isNaN(activeTab)) {
    result.activeTab = activeTab;
  }

  return result;
}
