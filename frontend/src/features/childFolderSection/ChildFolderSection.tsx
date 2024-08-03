import { IFileTreeDto } from '$models/dtos/fileTreeDto';

interface IChildFolderSectionProps {
  selectedFolder: IFileTreeDto;
}

export function ChildFolderSection({
  selectedFolder,
}: IChildFolderSectionProps) {
  // TODO MAYBE CREATE SECTION WRAPPER WITH NAME AND CHILDREN PROPERTY
  return (
    <section>
      <h2>Nested folder</h2>
      {selectedFolder.children.length === 0 && (
        <p>This folder contains no nested folder</p>
      )}
      {selectedFolder.children.length > 0 && (
        <ul>
          {selectedFolder.children.map((child) => (
            <li key={child.id}>{child.name}</li>
          ))}
        </ul>
      )}
    </section>
  );
}
