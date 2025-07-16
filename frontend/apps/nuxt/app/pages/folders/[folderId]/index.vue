<script setup lang="ts">
import { useFileTreeStore } from '~/stores/fileTreeStore';
const {
  params: { folderId },
} = useRoute();
const fileTreeStore = useFileTreeStore();
const folder = fileTreeStore.getFolderFromFileTree(folderId as string);
</script>
<template>
  <NuxtLayout name="folder-layout">
    <error-message
      v-if="!folder"
      :message="`Could not find folder with id '${folderId}'`"
    />
    <folder-view v-else :folder="folder" />
  </NuxtLayout>
</template>
