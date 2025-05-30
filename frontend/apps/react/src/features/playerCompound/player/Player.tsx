import { FileType } from '@videos-with-subtitle-player/core';

const baseUrl = import.meta.env.VITE_BASE_URL;

interface IPlayerProps {
  audioId: string;
  subtitleId?: string;
  fileType: FileType;
}

export function Player({ audioId, subtitleId, fileType }: IPlayerProps) {
  return (
    <div style={{ flexGrow: 1 }}>
      <video
        controls
        style={{ width: '100%' }}
        crossOrigin="anonymous"
        autoPlay
        data-testid="video"
      >
        <source
          type="audio/mp3"
          src={`${baseUrl}/api/file/continuous/${audioId}`}
          data-testid="source"
        />
        {fileType === FileType.AUDIO && subtitleId && (
          <track
            default
            kind="captions"
            srcLang="en"
            src={`${baseUrl}/api/file/discrete/${subtitleId}`}
            data-testid="track"
          />
        )}
      </video>
    </div>
  );
}
