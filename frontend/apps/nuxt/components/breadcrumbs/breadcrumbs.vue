<script setup lang="ts">
const {
  params: { folderId, fileId },
} = useRoute();
const fileTreeStore = useFileTreeStore();
const activeFolderPathInTree = fileTreeStore.getActiveFolderPathFromTree(
  (folderId as string) ?? '',
);
const isFileSelected = Boolean(fileId);
</script>
<template>
  <nav className="breadcrumbs text-sm max-w-full">
    <menu v-if="activeFolderPathInTree.length">
      <li>
        <NuxtLink to="/">Home</NuxtLink>
      </li>
      <li v-for="(currentPathItem, index) in activeFolderPathInTree">
        <template
          v-if="index === activeFolderPathInTree.length - 1 || isFileSelected"
        >
          {{ currentPathItem.name }}
        </template>
        <NuxtLink
          :to="{
            name: 'folders-folderId',
            params: { folderId: currentPathItem.id },
          }"
          v-else
          >{{ currentPathItem.name }}</NuxtLink
        >
      </li>
    </menu>
  </nav>
</template>
