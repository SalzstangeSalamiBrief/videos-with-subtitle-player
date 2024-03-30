import { FileType } from "../../../models/enums/FileType";

const baseUrl = import.meta.env.VITE_BASE_URL;

interface IPlayerProps {
  audioId: string;
  subtitleId?: string;
  fileType: FileType;
}

export function Player({ audioId, subtitleId, fileType }: IPlayerProps) {
  return (
    <video controls style={{ flexGrow: 1 }} crossOrigin="anonymous" autoPlay>
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
  );
}
