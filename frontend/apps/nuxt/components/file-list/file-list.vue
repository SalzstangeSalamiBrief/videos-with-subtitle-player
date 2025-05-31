<script setup lang="ts">
import { IFileNode } from '@videos-with-subtitle-player/core';

interface IProps {
  folderId: string;
  files: IFileNode[];
}

const { folderId, files } = defineProps<IProps>();
</script>

<template>
  <p v-if="!files.length">No sub folders</p>
  <ul
    v-else
    className="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4"
  >
    <li v-for="child in files" :key="child.id">
      <image-card
        :linkProps="{
          to: {
            name: 'folders-folderId-files-fileId',
            params: { folderId: folderId, fileId: child.id },
          },
        }"
        :title="child.name"
      />
    </li>
  </ul>
</template>
