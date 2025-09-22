<template>
  <div
    class="bg-white rounded-2xl p-8 shadow-lg hover:shadow-2xl transition-all duration-300 hover:-translate-y-1 text-center"
  >
    <div class="text-4xl mb-3">{{ icon }}</div>
    <h3 class="text-lg font-semibold">{{ title }}</h3>
    <p class="text-gray-500 mt-2">{{ desc }}</p>
  </div>
</template>

<script setup lang="ts">
defineProps<{ icon: string; title: string; desc: string }>();
</script>
