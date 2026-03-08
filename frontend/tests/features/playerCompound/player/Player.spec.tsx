import { FileType } from '$enums/FileType';
import { Player } from '$features/playerCompound/player/Player';
import { render, screen, within } from '@testing-library/react';
import { describe, expect, test } from 'vitest';

describe('Player', () => {
  test('Render video player without track', () => {
    render(
      <Player audioId="guid" fileType={FileType.Video} subtitleId="guid" />,
    );
    const video = screen.getByTestId('video');
    expect(video).toBeInTheDocument();
    const source = within(video).getByTestId('source');
    expect(source).toBeInTheDocument();
    const track = within(video).queryByTestId('track');
    expect(track).not.toBeInTheDocument();
  });

  test('Render audio player without track', () => {
    render(
      <Player
        audioId="guid"
        fileType={FileType.Audio}
        subtitleId={undefined}
      />,
    );
    const video = screen.getByTestId('video');
    expect(video).toBeInTheDocument();
    const source = within(video).getByTestId('source');
    expect(source).toBeInTheDocument();
    const track = within(video).queryByTestId('track');
    expect(track).not.toBeInTheDocument();
  });

  test('Render audio player with track', () => {
    render(
      <Player audioId="guid" fileType={FileType.Audio} subtitleId="guid" />,
    );
    const video = screen.getByTestId('video');
    expect(video).toBeInTheDocument();
    const source = within(video).getByTestId('source');
    expect(source).toBeInTheDocument();
    const track = within(video).queryByTestId('track');
    expect(track).toBeInTheDocument();
  });
});
