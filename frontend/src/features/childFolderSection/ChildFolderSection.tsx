import { IFileTreeDto } from '$models/dtos/fileTreeDto';
import { Card } from '$sharedComponents/card/Card';

interface IChildFolderSectionProps {
  selectedFolder: IFileTreeDto;
}

export function ChildFolderSection({
  selectedFolder,
}: IChildFolderSectionProps) {
  // TODO MAYBE CREATE SECTION WRAPPER WITH NAME AND CHILDREN PROPERTY
  return (
    <section>
      <h1>Nested folder</h1>
      {selectedFolder.children.length === 0 && (
        <p>This folder contains no nested folder</p>
      )}
      {selectedFolder.children.length > 0 && (
        <ul className="grid gap-4 grid-cols-2">
          {selectedFolder.children.map((child) => (
            <li key={child.id}>
              <Card title={child.name} imageUrl="/example.jpg"></Card>
            </li>
          ))}
        </ul>
      )}
    </section>
  );
}
