<script setup lang="ts">
import { RouterLinkProps } from 'vue-router';

interface IProps {
  linkProps: RouterLinkProps;
  title: string;
  imageUrl?: string;
}

const { imageUrl, linkProps, title } = defineProps<IProps>();
</script>
<style lang="css" scoped>
.card:hover {
  transform: rotateX(25deg);
  box-shadow: 0px 10px 10px -5px var(--color-fuchsia-300);
}
</style>
<template>
  <!-- TODO DOES NOT WORK WITH FILETREE AND FILENODE TOGETHER => FIX -->
  <article class="card card-border bg-slate-800">
    <figure class="h-56" :class="{ 'bg-fuchsia-800': !imageUrl }">
      <img
        v-if="imageUrl"
        loading="lazy"
        :src="imageHandler.getImageUrlForId(imageUrl)"
        :alt="`Cover image of the item ${title}`"
      />
    </figure>
    <div class="card-body h-28">
      <!-- TODO hover styles -->
      <NuxtLink class="hover:text-fuchsia-400" :to="linkProps.to">
        <h2
          class="line-clamp-3 overflow-hidden font-bold text-ellipsis"
          card-title
        >
          {{ title }}
        </h2>
      </NuxtLink>
    </div>
  </article>
</template>
