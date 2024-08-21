import { baseLinkStyles } from '$lib/styles/baseLinkStyles';
import { createFileRoute } from '@tanstack/react-router';
import { Link as TanStackRouterLink } from '@tanstack/react-router';
import { ImageCard } from '$sharedComponents/card/ImageCard';
import { Route as RootLayoutRoute } from './__root';

export const Route = createFileRoute('/')({ component: LandingPage });

function LandingPage() {
  const { fileTrees } = RootLayoutRoute.useLoaderData();

  // TODO HJANDLE NO ITEMS FOUND
  // TODO CHECK LOADING STATE
  // TODO REFACTOR COMPONENTS
  return (
    <ul className="grid gap-4 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
      {fileTrees.map((fileTree) => {
        // console.log(fileTree);
        return (
          <li key={fileTree.id}>
            <TanStackRouterLink
              to="/folders/$folderId"
              params={{ folderId: fileTree.id }}
              className={baseLinkStyles}
            >
              <ImageCard title={fileTree.name} imageUrl="/example.avif" />
            </TanStackRouterLink>
          </li>
        );
      })}
    </ul>
  );
}
