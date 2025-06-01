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
  <ul v-else className="list">
    <li v-for="child in files" :key="child.id" class="list-row">
      <NuxtLink
        :to="{
          name: 'folders-folderId-files-fileId',
          params: { folderId: folderId, fileId: child.id },
        }"
        class="hover:text-fuchsia-400"
      >
        {{ child.name }}
      </NuxtLink>
    </li>
  </ul>
</template>
