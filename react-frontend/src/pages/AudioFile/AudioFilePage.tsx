import { useParams } from "react-router-dom";
import { FileTreeContext } from "../../contexts/FileTreeContextWrapper";
import { useContext } from "react";
import { Button, Flex, Tooltip } from "antd";
import { LeftOutlined, RightOutlined } from "@ant-design/icons";
import { Link as ReactRouterLink } from "react-router-dom";
import { IFileTreeDto } from "../../models/dtos/fileTreeDto";
import { AudioPlayer } from "./components/AudioPlayer";
import { IFileNode } from "../../models/fileTree";
import { ErrorMessage } from "../../components/errorMessage/ErrorMessage";

export function AudioFilePage() {
  const { audioFileGroups, fileTrees } = useContext(FileTreeContext);
  const { audioId } = useParams();
  const fields = getFileIds(audioFileGroups, audioId ?? "");
  console.log(fields);
  const { nextAudioId, previousAudioId, currentFile } = fields;
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
    <Flex vertical>
      <h1 style={{ fontSize: "1.25rem", margin: 0 }}>
        {getParentName(fileTrees, audioId ?? "")}
      </h1>
      <h2
        style={{
          fontWeight: "normal",
          fontSize: "1rem",
          marginTop: 0,
          color: "hsl(0, 0%, 10%)",
        }}
      >
        {currentFile.name}
      </h2>
      <Flex gap="1rem" align="center">
        <Tooltip title="Previous track">
          <ReactRouterLink
            to={`/audio/${previousAudioId}`}
            aria-label="previous track"
          >
            <Button
              disabled={!previousAudioId}
              icon={<LeftOutlined />}
              aria-label="previous track"
            />
          </ReactRouterLink>
        </Tooltip>
        (
        <AudioPlayer
          key={currentFile.id}
          audioId={currentFile.id}
          subtitleId={currentFile.subtitleFileId}
          fileType={currentFile.fileType}
        />
        )
        <Tooltip title="Next track">
          <ReactRouterLink to={`/audio/${nextAudioId}`} aria-label="next track">
            <Button
              disabled={!nextAudioId}
              icon={<RightOutlined />}
              aria-label="next track"
            />
          </ReactRouterLink>
        </Tooltip>
      </Flex>
    </Flex>
  );
}

interface IGetFileFieldsReturn {
  previousAudioId: string | undefined;
  nextAudioId: string | undefined;
  currentFile: IFileNode | undefined;
}
const getFileIds = (
  audioFileGroups: IFileNode[][],
  audioId: string
): IGetFileFieldsReturn => {
  const result: IGetFileFieldsReturn = {
    currentFile: undefined,
    previousAudioId: undefined,
    nextAudioId: undefined,
  };
  console.log("audioFileGroups", audioFileGroups);
  const matchingAudioFileGroup = audioFileGroups.find((audioFileGroup) => {
    const containsAudioFile = audioFileGroup.find(
      (audioFile) => audioFile.id === audioId
    );
    return containsAudioFile;
  });

  if (!matchingAudioFileGroup) {
    console.warn(`Could not find audio file with id ${audioId}`);
    return result;
  }

  const matchingAudioFileIndex = matchingAudioFileGroup.findIndex(
    (audioFile) => audioFile.id === audioId
  );

  if (matchingAudioFileIndex < 0) {
    console.warn(`Could not find audio file with id ${audioId}`);
    return result;
  }

  const previousAudioIndex =
    matchingAudioFileIndex > 0 ? matchingAudioFileIndex - 1 : -1;
  const nextAudioIndex =
    matchingAudioFileIndex < matchingAudioFileGroup.length - 1
      ? matchingAudioFileIndex + 1
      : -1;

  result.previousAudioId = matchingAudioFileGroup[previousAudioIndex]?.id;
  result.nextAudioId = matchingAudioFileGroup[nextAudioIndex]?.id;
  result.currentFile = matchingAudioFileGroup[matchingAudioFileIndex];

  return result;
};

const getParentName = (fileTrees: IFileTreeDto[], audioId: string): string => {
  let parentName = "";

  if (fileTrees.length === 0 || !audioId) {
    return parentName;
  }

  fileTrees.forEach((fileTree) => {
    if (isPartOfSubTree(fileTree, audioId)) {
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
