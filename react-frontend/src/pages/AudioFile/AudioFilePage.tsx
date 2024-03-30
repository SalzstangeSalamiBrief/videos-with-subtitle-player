import { useParams } from "react-router-dom";
import { FileTreeContext } from "../../contexts/FileTreeContextWrapper";
import { useContext } from "react";
import { Button, Flex, Tooltip } from "antd";
import { LeftOutlined, RightOutlined } from "@ant-design/icons";
import { Link as ReactRouterLink } from "react-router-dom";
import { IFileTreeDto } from "../../models/dtos/fileTreeDto";
import { AudioPlayer } from "./components/AudioPlayer";
import { IFileNode } from "../../models/fileTree";

export function AudioFilePage() {
  const { audioFileGroups, fileTrees } = useContext(FileTreeContext);
  const { audioId } = useParams();
  const fields = getFileIds(audioFileGroups, audioId ?? "");
  console.log(fields);
  const {
    currentAudioId,
    currentSubtitleId,
    nextAudioId,
    previousAudioId,
    currentAudioName,
  } = fields;
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
        {currentAudioName}
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
        {currentAudioId && currentSubtitleId && (
          <AudioPlayer
            key={currentAudioId}
            audioId={currentAudioId}
            subtitleId={currentSubtitleId}
          />
        )}
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
  currentSubtitleId: string | undefined;
  currentAudioId: string | undefined;
  currentAudioName: string;
}
const getFileIds = (
  audioFileGroups: IFileNode[][],
  audioId: string
): IGetFileFieldsReturn => {
  const result: IGetFileFieldsReturn = {
    currentAudioId: undefined,
    currentSubtitleId: undefined,
    nextAudioId: undefined,
    previousAudioId: undefined,
    currentAudioName: "",
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

  const currentAudioId = matchingAudioFileGroup[matchingAudioFileIndex].id;
  const currentSubtitleId =
    matchingAudioFileGroup[matchingAudioFileIndex].subtitleFileId;

  result.previousAudioId = matchingAudioFileGroup[previousAudioIndex]?.id;
  result.nextAudioId = matchingAudioFileGroup[nextAudioIndex]?.id;
  result.currentSubtitleId = currentSubtitleId;
  result.currentAudioId = currentAudioId;
  result.currentAudioName = matchingAudioFileGroup[matchingAudioFileIndex].name;

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
