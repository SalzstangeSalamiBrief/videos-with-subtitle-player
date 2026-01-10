import { render, screen, within } from '@testing-library/react';
import { FileType } from '@videos-with-subtitle-player/core';
import { describe, expect, test } from 'vitest';
import { Player } from '$features/playerCompound/player/Player';

describe('Player', () => {
  test('Render video player without track', () => {
    render(
      <Player audioId="guid" fileType={FileType.VIDEO} subtitleId="guid" />,
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
        fileType={FileType.AUDIO}
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
      <Player audioId="guid" fileType={FileType.AUDIO} subtitleId="guid" />,
    );
    const video = screen.getByTestId('video');
    expect(video).toBeInTheDocument();
    const source = within(video).getByTestId('source');
    expect(source).toBeInTheDocument();
    const track = within(video).queryByTestId('track');
    expect(track).toBeInTheDocument();
  });
});
