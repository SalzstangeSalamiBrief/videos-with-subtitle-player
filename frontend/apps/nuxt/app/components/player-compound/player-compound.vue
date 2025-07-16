<script setup lang="ts">
import { IFileNode } from '@videos-with-subtitle-player/core';

interface IProps {
  currentFile: IFileNode;
  siblings: IFileNode[];
  folderId: string;
}

const { siblings, currentFile, folderId } = defineProps<IProps>();
</script>
<template>
  <div className="flex gap-4">
    <player
      :audioId="currentFile.id"
      :fileType="currentFile.fileType"
      :subtitleId="currentFile.subtitleFileId"
    />
    <playlist :items="siblings" :folderId="folderId" :fileId="currentFile.id" />
  </div>
</template>
