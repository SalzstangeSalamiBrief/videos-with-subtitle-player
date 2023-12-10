import { useParams } from "react-router-dom";
import { FileTreeContext } from "../../contexts/FileTreeContextWrapper";
import { useContext } from "react";
import { IAudioFileDto } from "../../models/audioFileDto";
import { Button, Flex, Tooltip } from "antd";
import { LeftOutlined, RightOutlined } from "@ant-design/icons";
import { Link as ReactRouterLink } from "react-router-dom";

export function AudioFilePage() {
  const { audioFileGroups } = useContext(FileTreeContext);
  const { audioId } = useParams();
  const { currentAudioId, currentSubtitleId, nextAudioId, previousAudioId } =
    getFileIds(audioFileGroups, audioId ?? "");

  // TODO REMOVE LOCALHOST
  return (
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
      <video
        controls
        src={`http://localhost:3000/api/file/${currentAudioId}`}
        style={{ flexGrow: 1 }}
      >
        <track
          default
          kind="captions"
          srcLang="en"
          src={`http://localhost:3000/api/file/${currentSubtitleId}`}
        />
      </video>
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
  );
}

interface IGetFileFieldsReturn {
  previousAudioId: string | undefined;
  nextAudioId: string | undefined;
  currentSubtitleId: string | undefined;
  currentAudioId: string | undefined;
}
const getFileIds = (
  audioFileGroups: IAudioFileDto[][],
  audioId: string
): IGetFileFieldsReturn => {
  const result: IGetFileFieldsReturn = {
    currentAudioId: undefined,
    currentSubtitleId: undefined,
    nextAudioId: undefined,
    previousAudioId: undefined,
  };

  const matchingAudioFileGroup = audioFileGroups.find((audioFileGroup) => {
    const containsAudioFile = audioFileGroup.find(
      (audioFile) => audioFile.audioFile.id === audioId
    );
    return containsAudioFile;
  });

  if (!matchingAudioFileGroup) {
    console.warn(`Could not find audio file with id ${audioId}`);
    return result;
  }

  const matchingAudioFileIndex = matchingAudioFileGroup.findIndex(
    (audioFile) => audioFile.audioFile.id === audioId
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

  const currentAudioId =
    matchingAudioFileGroup[matchingAudioFileIndex].audioFile.id;
  const currentSubtitleId =
    matchingAudioFileGroup[matchingAudioFileIndex].subtitleFile.id;

  result.previousAudioId =
    matchingAudioFileGroup[previousAudioIndex]?.audioFile?.id;
  result.nextAudioId = matchingAudioFileGroup[nextAudioIndex]?.audioFile?.id;
  result.currentSubtitleId = currentSubtitleId;
  result.currentAudioId = currentAudioId;

  return result;
};
