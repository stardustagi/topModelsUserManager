<template>
  <div
    class="bg-white rounded-2xl p-8 shadow-lg hover:shadow-2xl transition-all duration-300 hover:-translate-y-1 text-center"
  >
    <div class="text-5xl mb-3">{{ icon }}</div>
    <div class="flex items-center justify-center gap-2">
      <h3 class="text-lg font-semibold">{{ title }}</h3>
      <span
        v-if="badge"
        class="text-xs px-2 py-0.5 bg-red-500 text-white rounded-full"
        >{{ badge }}</span
      >
    </div>
    <p v-if="desc" class="text-gray-500 mt-2">{{ desc }}</p>
    <div v-if="tags?.length" class="flex flex-wrap gap-2 justify-center mt-3">
      <el-tag v-for="t in tags" :key="t">{{ t }}</el-tag>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  icon: string;
  title: string;
  desc?: string;
  tags?: string[];
  badge?: string;
}>();
</script>
