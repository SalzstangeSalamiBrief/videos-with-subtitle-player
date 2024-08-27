import { createFileRoute } from '@tanstack/react-router';
import { Route as RootLayoutRoute } from './__root';
import { FolderListSection } from '$features/folderListSection/FolderListSection';

export const Route = createFileRoute('/')({ component: LandingPage });

function LandingPage() {
  const { fileTrees } = RootLayoutRoute.useLoaderData();

  // TODO HJANDLE NO ITEMS FOUND
  // TODO CHECK LOADING STATE
  // TODO REFACTOR COMPONENTS
  return <FolderListSection folders={fileTrees} />;
}
