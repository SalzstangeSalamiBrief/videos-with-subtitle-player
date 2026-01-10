import { Player } from './player/Player';
import { Playlist } from './playlist/Playlist';
import type { IFileNode } from '$models/fileTree/fileTree';

interface IPlayerCompoundProps {
  currentFile: IFileNode;
  siblings: IFileNode[];
}

export function PlayerCompound({
  currentFile,
  siblings,
}: IPlayerCompoundProps) {
  return (
    <div className="flex gap-4">
      <Player
        key={currentFile.id}
        audioId={currentFile.id}
        subtitleId={currentFile.subtitleFileId}
        fileType={currentFile.fileType}
      />

      <Playlist siblings={siblings} />
    </div>
  );
}
