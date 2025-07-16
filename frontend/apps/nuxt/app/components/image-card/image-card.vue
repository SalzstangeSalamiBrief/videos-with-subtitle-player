<script setup lang="ts">
import { RouterLinkProps } from 'vue-router';
import { IProgressiveImageProps } from '../progressive-image/progressive-image-props';

interface IProps {
  linkProps: RouterLinkProps;
  title: string;
  imageUrls?: Omit<IProgressiveImageProps, 'alt'>;
}

const { imageUrls, linkProps, title } = defineProps<IProps>();
const alt = `Cover image of the item ${title}`;
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
    <div class="relative h-56" :class="{ 'bg-fuchsia-800': !imageUrls }">
      <ProgressiveImage
        v-if="imageUrls"
        :alt="alt"
        :highQualityImageUrl="
          imageHandler.getImageUrlForId(imageUrls?.highQualityImageUrl)
        "
        :lowQualityImageUrl="
          imageHandler.getImageUrlForId(imageUrls?.lowQualityImageUrl)
        "
      />
    </div>
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
