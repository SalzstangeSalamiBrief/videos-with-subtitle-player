<script setup lang="ts">
import { IFileTree } from '@videos-with-subtitle-player/core';

interface IProps {
  folders: IFileTree[];
}

const { folders } = defineProps<IProps>();
</script>

<template>
  <p v-if="!folders.length">no sub folders</p>
  <ul
    v-else
    className="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4"
  >
    <li v-for="folder in folders">
      <image-card
        :imageUrls="{
          highQualityImageUrl: folder.thumbnailId,
          lowQualityImageUrl: folder.lowQualityThumbnailId,
        }"
        :title="folder.name"
        :linkProps="{
          to: { name: 'folders-folderId', params: { folderId: folder.id } },
        }"
        :folder="folder"
        :key="folder.id"
      />
    </li>
  </ul>
</template>
