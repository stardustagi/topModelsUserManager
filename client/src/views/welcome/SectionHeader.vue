<template>
  <div class="text-center py-14">
    <h2 class="text-4xl font-bold mb-3">{{ title }}</h2>
    <p v-if="subtitle" class="text-gray-500">{{ subtitle }}</p>
  </div>
</template>

<script setup lang="ts">
defineProps<{ title: string; subtitle?: string }>();
</script>
