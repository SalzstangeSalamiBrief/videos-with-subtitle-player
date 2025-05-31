import {
  getFileTreeQuery,
  getFileTreeSelect,
  type IGetFileTreeSelectReturn,
} from '@videos-with-subtitle-player/core';

const fileTreeStoreKey = 'file-tree-store';

export const useFileTreeStore = defineStore(fileTreeStoreKey, {
  state: (): IGetFileTreeSelectReturn => ({
    fileGroups: [],
    fileTrees: [],
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
    },
  },
});
