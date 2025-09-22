<template>
  <div
    class="bg-white rounded-2xl p-8 shadow-lg hover:shadow-xl transition-all duration-300 text-center"
  >
    <div class="text-4xl font-bold text-purple-600 mb-3">{{ num }}</div>
    <h3 class="text-lg font-semibold">{{ title }}</h3>
    <p class="text-gray-500 mt-2">{{ desc }}</p>
  </div>
</template>

<script setup lang="ts">
defineProps<{ num: string; title: string; desc: string }>();
</script>
