import { FileType } from '$enums/FileType';

const baseUrl = import.meta.env.VITE_BASE_URL;

interface IPlayerProps {
  audioId: string;
  subtitleId?: string;
  fileType: FileType;
}

export function Player({ audioId, subtitleId, fileType }: IPlayerProps) {
  return (
    <div className="grow rounded-md border border-slate-700">
      <video
        controls
        className="h-full w-full"
        crossOrigin="anonymous"
        autoPlay
        data-testid="video"
      >
        <source
          type="audio/mp3"
          src={`${baseUrl}/api/file/continuous/${audioId}`}
          data-testid="source"
        />
        {fileType === FileType.Audio && subtitleId && (
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
