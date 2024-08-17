import { IFileTreeDto } from '$models/dtos/fileTreeDto';
import { ImageCard } from '$sharedComponents/card/ImageCard';

interface IFileListSectionProps {
  selectedFolder: IFileTreeDto;
}

export function FileListSection({ selectedFolder }: IFileListSectionProps) {
  return (
    <section role="presentation">
      {selectedFolder.files.length === 0 && (
        <p>This folder contains no files</p>
      )}
      {selectedFolder.files.length > 0 && (
        <ul className="grid gap-4 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
          {selectedFolder.files.map((file) => (
            <li key={file.id}>
              <ImageCard title={file.name} imageUrl="/example.jpg" />
            </li>
          ))}
        </ul>
      )}
    </section>
  );
}
