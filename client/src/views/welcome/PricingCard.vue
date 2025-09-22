<template>
  <div
    :class="[
      'bg-white rounded-2xl p-8 shadow-lg hover:shadow-2xl transition-all duration-300 hover:-translate-y-1 flex flex-col text-center',
      highlight ? 'border-2 border-purple-500' : ''
    ]"
  >
    <div class="text-5xl mb-3">{{ icon }}</div>
    <h3 class="text-xl font-semibold">{{ title }}</h3>
    <div v-if="price" class="my-3 text-2xl font-bold">{{ price }}</div>
    <p class="text-gray-500">{{ desc }}</p>
    <div class="mt-6">
      <el-button
        :type="highlight ? 'primary' : 'default'"
        class="w-full !h-11 rounded-full"
        >{{ cta }} →</el-button
      >
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  icon: string;
  title: string;
  price?: string;
  desc: string;
  cta: string;
  highlight?: boolean;
}>();
</script>
