<script setup lang="ts">
import { Maybe } from '@videos-with-subtitle-player/core';

const {
  params: { fileId, folderId },
} = useRoute() as unknown as { params: { fileId: string; folderId: string } };
const fileTreeStore = useFileTreeStore();
const [siblings, currentFile] = fileTreeStore.getCurrentNodeWithSiblings(
  fileId as Maybe<string>,
);
</script>

<template>
  <NuxtLayout name="folder-layout">
    <!-- TODO CREATE ERROR MESSAGE COMPONENT -->
    <p v-if="!currentFile">Could not find file.</p>
    <div v-else class="grid">
      <h1 className="m-0 text-lg font-bold">{{ currentFile.name }}</h1>
      <player-compound
        :currentFile="currentFile"
        :siblings="siblings"
        :folderId="folderId"
      />
    </div>
  </NuxtLayout>
</template>
