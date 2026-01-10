import { FolderList } from '$sharedComponents/folderList/FolderList';
import { createFileRoute } from '@tanstack/react-router';
import { Route as RootLayoutRoute } from './__root';

export const Route = createFileRoute('/')({
  component: LandingPage,
});

function LandingPage() {
  document.title = import.meta.env.VITE_APP_TITLE;
  const { fileTrees } = RootLayoutRoute.useLoaderData();

  // TODO HJANDLE NO ITEMS FOUND
  // TODO CHECK LOADING STATE
  // TODO REFACTOR COMPONENTS
  if (!fileTrees.length) {
    return <p>The app does not contain any folders and files.</p>;
  }

  return <FolderList folders={fileTrees} />;
}
