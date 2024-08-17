import { createFileRoute } from '@tanstack/react-router';
import { useContext } from 'react';
import { Link as TansStackLink } from '@tanstack/react-router';
import { ErrorMessage } from '$sharedComponents/errorMessage/ErrorMessage';
import { Player } from '$sharedComponents/player/Player';
import { FileTreeContext } from '$contexts/FileTreeContextWrapper';
import { IFileNode } from '$models/fileTree';
import { ChevronLeftIcon, ChevronRightIcon } from '@heroicons/react/24/outline';

export const Route = createFileRoute(
  '/folders/_folderLayout/$folderId/files/$fileId/',
)({
  component: FilePage,
  // TODO ADD META  => SET TITLE AS TITLE OF THE FILE
});

function FilePage() {
  const { fileGroups } = useContext(FileTreeContext);
  const { fileId, folderId } = Route.useParams();
  const { nextId, previousId, currentFile } = getFileIds(fileGroups, fileId);

  if (!currentFile) {
    return (
      <ErrorMessage
        error="Could not find file."
        message="Something went wrong"
        description="Please try again later."
      />
    );
  }

  return (
    <div className="grid">
      <h1 style={{ fontSize: '1.25rem', margin: 0 }}>{currentFile.name}</h1>
      <h2
        style={{
          fontWeight: 'normal',
          fontSize: '1rem',
          marginTop: 0,
          color: 'hsl(0, 0%, 10%)',
        }}
      >
        {currentFile.name}
      </h2>
      <div className="flex gap-4 items-center">
        <TansStackLink
          to="/folders/$folderId/files/$fileId/"
          params={{ fileId: previousId ?? '', folderId }}
          aria-label="previous track"
        >
          <button
            className="p-4 bg-slate-800 hover:bg-slate-700"
            disabled={!previousId}
            aria-label="Previous track"
            title="Previous track"
          >
            <ChevronLeftIcon />
            <span className="sr-only">Previous track</span>
          </button>
        </TansStackLink>

        <Player
          key={currentFile.id}
          audioId={currentFile.id}
          subtitleId={currentFile.subtitleFileId}
          fileType={currentFile.fileType}
        />

        <TansStackLink
          to="/folders/$folderId/files/$fileId/"
          params={{ fileId: nextId ?? '', folderId }}
          aria-label="Next track"
        >
          <button
            className="p-4 bg-slate-800 hover:bg-slate-700"
            disabled={!nextId}
            aria-label="next track"
            title="Next track"
          >
            <ChevronRightIcon />
            <span className="sr-only">Next track</span>
          </button>
        </TansStackLink>
      </div>
    </div>
  );
}

interface IGetFileFieldsReturn {
  previousId: string | undefined;
  nextId: string | undefined;
  currentFile: IFileNode | undefined;
}
const getFileIds = (
  fileGroups: IFileNode[][],
  fileId: string | undefined,
): IGetFileFieldsReturn => {
  const result: IGetFileFieldsReturn = {
    currentFile: undefined,
    previousId: undefined,
    nextId: undefined,
  };

  if (!fileId) {
    return result;
  }

  const matchingAudioFileGroup = fileGroups.find((audioFileGroup) => {
    const containsAudioFile = audioFileGroup.find(
      (audioFile) => audioFile.id === fileId,
    );
    return containsAudioFile;
  });

  if (!matchingAudioFileGroup) {
    console.warn(`Could not find audio file with id ${fileId}`);
    return result;
  }

  const matchingAudioFileIndex = matchingAudioFileGroup.findIndex(
    (audioFile) => audioFile.id === fileId,
  );

  if (matchingAudioFileIndex < 0) {
    console.warn(`Could not find audio file with id ${fileId}`);
    return result;
  }

  const previousAudioIndex =
    matchingAudioFileIndex > 0 ? matchingAudioFileIndex - 1 : -1;
  const nextAudioIndex =
    matchingAudioFileIndex < matchingAudioFileGroup.length - 1
      ? matchingAudioFileIndex + 1
      : -1;

  result.previousId = matchingAudioFileGroup[previousAudioIndex]?.id;
  result.nextId = matchingAudioFileGroup[nextAudioIndex]?.id;
  result.currentFile = matchingAudioFileGroup[matchingAudioFileIndex];

  return result;
};
