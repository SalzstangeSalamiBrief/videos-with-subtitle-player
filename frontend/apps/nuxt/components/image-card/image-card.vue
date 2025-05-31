<script setup lang="ts">
import { IFileTree } from '@videos-with-subtitle-player/core';

interface Props {
  folder: IFileTree;
}

const { folder } = defineProps<Props>();
</script>
#
<style lang="css" scoped>
.imageCard:hover {
  transform: rotateX(25deg);
  box-shadow: 0px 10px 10px -5px var(--color-fuchsia-300);
}
</style>
<template>
  <article class="card card-border imageCard bg-slate-800">
    <figure class="h-56" :class="{ 'bg-fuchsia-800': !folder.thumbnailId }">
      <img
        v-if="folder.thumbnailId"
        loading="lazy"
        :src="imageHandler.getImageUrlForId(folder.thumbnailId)"
        :alt="`Cover image of the item ${folder.name}`"
      />
    </figure>
    <div class="card-body h-20">
      <!-- TODO hover styles -->
      <NuxtLink
        class="hover:text-fuchsia-400"
        :to="{ name: 'folders-folderId', params: { folderId: folder.id } }"
      >
        <h2
          class="line-clamp-3 overflow-hidden font-bold text-ellipsis"
          card-title
        >
          {{ folder.name }}
        </h2>
      </NuxtLink>
    </div>
  </article>
</template>
