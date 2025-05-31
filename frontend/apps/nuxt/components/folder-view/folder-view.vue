<script setup lang="ts">
import { IFileTree } from '@videos-with-subtitle-player/core';

interface IProps {
  folder: IFileTree;
}

const { folder } = defineProps<IProps>();
const activeTab = ref(0);
</script>
<style lang="css" scoped>
.tabs {
  background-color: var(--color-slate-800);
}
</style>
<template>
  <div class="tabs tabs-box">
    <input
      type="radio"
      name="folder-tab"
      class="tab"
      aria-label="Folder tab"
      :checked="activeTab === 0"
      @input="() => (activeTab = 0)"
    />
    <div class="tab-content bg-base-100 border-base-300 p-6">
      <!-- TODO DOES NOT WORK WITH ROUTING? -->
      <file-list :folder="folder" />
    </div>

    <input
      type="radio"
      name="video-file-tab"
      class="tab"
      aria-label="Video file tab"
      :checked="activeTab === 1"
      @input="() => (activeTab = 1)"
    />
    <div class="tab-content bg-base-100 border-base-300 p-6">
      <file-list :folder="folder.videos" />
    </div>

    <input
      type="radio"
      name="audio-file-tab"
      class="tab"
      aria-label="Audio file tab"
      :checked="activeTab === 2"
      @input="() => (activeTab = 2)"
    />
    <div class="tab-content bg-base-100 border-base-300 p-6">Tab content 3</div>

    <input
      type="radio"
      name="image-tab"
      class="tab"
      aria-label="Image tab"
      :checked="activeTab === 3"
      @input="() => (activeTab = 3)"
    />
    <div class="tab-content bg-base-100 border-base-300 p-6">Tab content 3</div>
  </div>
</template>
