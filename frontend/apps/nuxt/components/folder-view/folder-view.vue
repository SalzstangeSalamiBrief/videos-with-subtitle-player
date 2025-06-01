<script setup lang="ts">
import { IFileTree } from '@videos-with-subtitle-player/core';
import FileList from '../file-list/file-list.vue';
import FolderList from '../folder-list/folder-list.vue';

interface IProps {
  folder: IFileTree;
}
const { folder } = defineProps<IProps>();

const activeTab = ref(0);

interface ITab {
  title: string;
  component: Component;
  props: any;
}
const tabs: ITab[] = [
  {
    title: 'Folder tab',
    component: FolderList,
    props: { folders: folder.children },
  },
  {
    title: 'Audio files',
    component: FileList,
    props: {
      folderId: folder.id,
      files: folder.audios,
    },
  },
  {
    title: 'Video files',
    component: FileList,
    props: {
      folderId: folder.id,
      files: folder.videos,
    },
  },
  // TODO IMPLEMENT
  // {
  //   title: "Image tab",
  //   component:
  //   props:{}
  // }
];
</script>
<style lang="css" scoped>
.tabs {
  background-color: var(--color-slate-800);
}
</style>
<template>
  <div class="tabs tabs-box">
    <template v-for="(tab, index) in tabs" :key="tab.title">
      <input
        type="radio"
        :name="tab.title"
        class="tab"
        :aria-label="tab.title"
        :checked="activeTab === index"
        @input="() => (activeTab = index)"
      />
      <div class="tab-content bg-base-100 border-base-300 p-6">
        <component :is="tab.component" v-bind="tab.props" />
      </div>
    </template>
  </div>
</template>
