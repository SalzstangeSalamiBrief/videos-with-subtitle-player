<script setup lang="ts">
import { FileType } from '@videos-with-subtitle-player/core';

interface IProps {
  audioId: string;
  subtitleId?: string;
  fileType: FileType;
}

const { audioId, fileType, subtitleId } = defineProps<IProps>();

// TODO ENV VARIABLE
const baseUrl = 'http://localhost:3000';

const audioSource = `${baseUrl}/api/file/continuous/${audioId}`;
const subtitleSource = `${baseUrl}/api/file/discrete/${subtitleId}`;
</script>
<template>
  <div class="grow">
    <video
      controls
      class="w-full"
      crossOrigin="anonymous"
      autoPlay
      data-testid="video"
    >
      <source type="audio/mp3" :src="audioSource" data-testid="source" />
      <track
        v-if="fileType === FileType.AUDIO && subtitleId"
        default
        kind="captions"
        srcLang="en"
        :src="subtitleSource"
        data-testid="track"
      />
    </video>
  </div>
</template>
