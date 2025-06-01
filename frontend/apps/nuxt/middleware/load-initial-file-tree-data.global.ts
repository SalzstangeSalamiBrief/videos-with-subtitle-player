export default defineNuxtRouteMiddleware(async () => {
  const fileTreeStore = useFileTreeStore();
  if (fileTreeStore.isInitialized) {
    return;
  }

  await callOnce(fileTreeStore.init);
});
