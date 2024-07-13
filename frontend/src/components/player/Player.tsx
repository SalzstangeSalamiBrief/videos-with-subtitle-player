import { FileType } from '../../enums/FileType';

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
      >
        <source type="audio/mp3" src={`${baseUrl}/api/file/audio/${audioId}`} />
        {fileType === FileType.AUDIO && subtitleId && (
          <track
            default
            kind="captions"
            srcLang="en"
            src={`${baseUrl}/api/file/subtitle/${subtitleId}`}
          />
        )}
      </video>
    </div>
  );
}
