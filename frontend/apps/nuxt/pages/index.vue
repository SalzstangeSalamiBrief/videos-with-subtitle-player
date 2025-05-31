<script setup lang="ts">
import { useFileTreeStore } from '~/stores/fileTreeStore';

const fileTreeStore = useFileTreeStore();
</script>
<template>
  <p v-if="!fileTreeStore.fileTrees.length">
    The app does not contain any folders and files.
  </p>
  <folder-list v-else :folders="fileTreeStore.fileTrees" />
</template>
