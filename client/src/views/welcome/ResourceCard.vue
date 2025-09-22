<template>
  <div
    :class="[
      'rounded-2xl p-8 transition-all duration-300 hover:-translate-y-1',
      dark
        ? 'bg-gray-900 text-white shadow-xl hover:shadow-2xl'
        : 'bg-white shadow-lg hover:shadow-2xl'
    ]"
  >
    <div class="flex items-start justify-between">
      <div class="text-4xl mb-3">{{ icon }}</div>
      <span
        v-if="tag"
        class="text-[10px] px-2 py-0.5 rounded-full bg-purple-500 text-white"
        >{{ tag }}</span
      >
    </div>
    <h3 class="text-lg font-semibold">{{ title }}</h3>
    <p :class="['mt-2', dark ? 'text-gray-300' : 'text-gray-500']">
      {{ desc }}
    </p>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  icon: string;
  title: string;
  desc: string;
  dark?: boolean;
  tag?: string;
}>();
</script>
