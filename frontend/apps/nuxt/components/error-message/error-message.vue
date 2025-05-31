<script setup lang="ts">
interface IProps {
  message: string;
  description?: string;
}

const { message, description } = defineProps<IProps>();
</script>
<template>
  <section className="h-fit rounded-md bg-red-50 p-4 text-red-900">
    <h1 className="text-lg font-bold">{{ message }}</h1>
    <p v-if="description">{{ description }}</p>
  </section>
</template>
