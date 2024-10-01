import { baseLinkStyles } from '$lib/styles/baseLinkStyles';
import { IFileNode } from '$models/fileTree';
import { ImageCard } from '$sharedComponents/card/ImageCard';
import { Link as TanStackRouterLink } from '@tanstack/react-router';

interface IFileListSectionProps {
  folderId: string;
  files: IFileNode[];
}

export function FileListSection({ files, folderId }: IFileListSectionProps) {
  return (
    <section role="presentation">
      {files.length === 0 && <p>This folder contains no files</p>}
      {files.length > 0 && (
        <ul className="grid gap-4 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
          {files.map((file) => (
            <li key={file.id}>
              <TanStackRouterLink
                to="/folders/$folderId/files/$fileId"
                params={{ folderId, fileId: file.id }}
                className={baseLinkStyles}
              >
                <ImageCard title={file.name} imageUrl="/example.avif" />
              </TanStackRouterLink>
            </li>
          ))}
        </ul>
      )}
    </section>
  );
}
