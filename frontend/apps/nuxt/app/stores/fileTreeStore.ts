import {
  getCurrentNodeWithSiblings,
  getFileTreeQuery,
  getFileTreeSelect,
  getFolderFromFileTree,
  getFoldersInActiveTree,
  Maybe,
  type IGetFileTreeSelectReturn,
} from '@videos-with-subtitle-player/core';
import { defineStore } from 'pinia';

const fileTreeStoreKey = 'file-tree-store';

type FileTreeStoreData = IGetFileTreeSelectReturn & { isInitialized: boolean };

export const useFileTreeStore = defineStore(fileTreeStoreKey, {
  state: (): FileTreeStoreData => ({
    fileGroups: [],
    fileTrees: [],
    isInitialized: false,
  }),
  actions: {
    async init() {
      const {
        data: { value },
        error,
      } = await useAsyncData('file-tree-query', () =>
        getFileTreeQuery('http://localhost:3000'),
      );
      const selectResult = getFileTreeSelect(value ?? []);

      this.fileGroups = selectResult.fileGroups;
      this.fileTrees = selectResult.fileTrees;
      this.isInitialized = true;
    },
  },
  getters: {
    getFolderFromFileTree: (state: FileTreeStoreData) => {
      return (folderId: string) =>
        getFolderFromFileTree(state.fileTrees, folderId);
    },
    getActiveFolderPathFromTree: (state: FileTreeStoreData) => {
      return (folderId: string) =>
        getFoldersInActiveTree(state.fileTrees, folderId);
    },

    getCurrentNodeWithSiblings: (state: FileTreeStoreData) => {
      return (
        fileId: Maybe<string>,
      ): ReturnType<typeof getCurrentNodeWithSiblings> => {
        if (!fileId) {
          return [[], undefined];
        }

        return getCurrentNodeWithSiblings(state.fileGroups, fileId);
      };
    },
  },
});
