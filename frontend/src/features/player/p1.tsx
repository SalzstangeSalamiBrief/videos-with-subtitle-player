import { LeftOutlined, RightOutlined } from '@ant-design/icons';
import { createFileRoute } from '@tanstack/react-router';
import { Tooltip, Button } from 'antd';
import { useContext } from 'react';
import { Link as TansStackLink } from '@tanstack/react-router';
import { ErrorMessage } from '$sharedComponents/errorMessage/ErrorMessage';
import { Player } from '$sharedComponents/player/Player';
import { FileTreeContext } from '$contexts/FileTreeContextWrapper';
import { IFileTreeDto } from '$models/dtos/fileTreeDto';
import { IFileNode } from '$models/fileTree';

export const Route = createFileRoute('/folders/$folderId/')({
  component: AudioFilePage,
  // TODO VALIDATION OF THE PARAM
});

function AudioFilePage() {
  const { fileGroups, fileTrees } = useContext(FileTreeContext);
  const { folderId } = Route.useParams();
  const { nextId, previousId, currentFile } = getFileIds(fileGroups, folderId);

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
    <div className="grid gap-4">
      <h1 style={{ fontSize: '1.25rem', margin: 0 }}>
        {getParentName(fileTrees, folderId ?? '')}
      </h1>
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
      <div className="flex gap-2 items-center">
        <Tooltip title="Previous track">
          <TansStackLink
            to="/files/$fileId/"
            params={{ fileId: previousId ?? '' }}
            aria-label="previous track"
          >
            <Button
              disabled={!previousId}
              icon={<LeftOutlined />}
              aria-label="previous track"
            />
          </TansStackLink>
        </Tooltip>

        <Player
          key={currentFile.id}
          audioId={currentFile.id}
          subtitleId={currentFile.subtitleFileId}
          fileType={currentFile.fileType}
        />

        <Tooltip title="Next track">
          <TansStackLink
            to="/files/$fileId/"
            params={{ fileId: nextId ?? '' }}
            aria-label="next track"
          >
            <Button
              disabled={!nextId}
              icon={<RightOutlined />}
              aria-label="next track"
            />
          </TansStackLink>
        </Tooltip>
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

const getParentName = (fileTrees: IFileTreeDto[], folderId: string): string => {
  let parentName = '';

  if (fileTrees.length === 0 || !folderId) {
    return parentName;
  }

  fileTrees.forEach((fileTree) => {
    if (isPartOfSubTree(fileTree, folderId)) {
      parentName = fileTree.name;
      return;
    }
  });

  return parentName;
};

const isPartOfSubTree = (fileTree: IFileTreeDto, audioId: string): boolean => {
  if (fileTree.files?.find((audioFile) => audioFile.id === audioId)) {
    return true;
  }

  if (fileTree.children?.length) {
    return fileTree.children.some((child) => isPartOfSubTree(child, audioId));
  }

  return false;
};
