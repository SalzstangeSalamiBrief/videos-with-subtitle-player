const baseUrl = import.meta.env.VITE_BASE_URL;

interface IAudioPlayerProps {
  audioId: string;
  subtitleId: string;
}

export function AudioPlayer({ audioId, subtitleId }: IAudioPlayerProps) {
  return (
    <video controls style={{ flexGrow: 1 }} crossOrigin="anonymous">
      <source type="audio/mp3" src={`${baseUrl}/api/file/audio/${audioId}`} />
      <track
        default
        kind="captions"
        srcLang="en"
        src={`${baseUrl}/api/file/subtitle/${subtitleId}`}
      />
    </video>
  );
}
