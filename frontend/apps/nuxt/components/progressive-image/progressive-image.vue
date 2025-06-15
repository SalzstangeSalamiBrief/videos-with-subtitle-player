<script setup lang="ts">
import { ref } from 'vue';
import { IProgressiveImageProps } from './progressive-image-props';

const { alt, highQualityImageUrl, lowQualityImageUrl } =
  defineProps<IProgressiveImageProps>();

const isHighQualityImageLoaded = ref<boolean>(false);
</script>
<style lang="css" scoped>
.progressive-image-container {
  position: relative;
  height: 100%;
  width: 100%;
  overflow: hidden;
}

.progressive-image {
  position: absolute;
  height: 100%;
  width: 100%;
  object-fit: contain;

  /* hide alt text */
  font-size: 0;

  transition:
    opacity 0.6s ease-in-out,
    visibility 0.6s ease-in-out;

  &.low-quality {
    filter: blur(10px);
  }

  &.high-quality {
    filter: none;
  }

  &.hiddenImage {
    visibility: hidden;
    opacity: 0;
  }
}
</style>
<template>
  <div class="progressive-image-container">
    <img
      v-if="lowQualityImageUrl"
      loading="eager"
      :src="lowQualityImageUrl"
      class="progressive-image low-quality"
      :class="{ hiddenImage: isHighQualityImageLoaded }"
      :alt="alt"
    />
    <ClientOnly>
      <img
        v-if="highQualityImageUrl"
        loading="lazy"
        :src="highQualityImageUrl"
        class="progressive-image high-quality"
        :class="{ hiddenImage: !isHighQualityImageLoaded }"
        :alt="alt"
        @load="isHighQualityImageLoaded = true"
      />
    </ClientOnly>
  </div>
</template>
