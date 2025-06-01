<script setup lang="ts">
import { IFileNode } from '@videos-with-subtitle-player/core';

interface IProps {
  items: IFileNode[];
  folderId: string;
  fileId: string;
}

const { items, fileId, folderId } = defineProps<IProps>();
const labelId = useId();

function getAriaLabel(fileName: string): string {
  return `Play '${fileName}''`;
}
</script>

<template>
  <section :aria-labelledby="labelId">
    <h2 :id="labelId" class="mb-2 text-base font-bold">Playlist</h2>
    <ol class="list h-96 w-80 overflow-y-scroll">
      <li v-for="item in items" class="list-row">
        <NuxtLink
          :title="item.name"
          :aria-label="getAriaLabel(item.name)"
          :to="{
            name: 'folders-folderId-files-fileId',
            params: { fileId: item.id, folderId },
          }"
          :aria-selected="item.id === fileId ? 'true' : 'false'"
          class="hover:text-fuchsia-400"
          :class="{
            'text-fuchsia-500': item.id === fileId,
          }"
        >
          {{ item.name }}
        </NuxtLink>
      </li>
    </ol>
  </section>
</template>
