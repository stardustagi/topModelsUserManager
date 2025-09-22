<template>
  <div
    class="bg-white rounded-2xl p-6 shadow-lg hover:shadow-xl transition-all duration-300 cursor-pointer"
    @click="open = !open"
  >
    <div class="flex items-center justify-between">
      <div class="flex items-center gap-2">
        <span class="text-purple-600 font-bold">Q</span>
        <h3 class="text-lg font-semibold">{{ q }}</h3>
      </div>
      <el-icon
        :class="[
          open ? 'rotate-180' : 'rotate-0',
          'transition-transform duration-300'
        ]"
        ><ArrowDown
      /></el-icon>
    </div>
    <transition name="faq">
      <p v-if="open" class="mt-4 text-gray-600 leading-relaxed border-t pt-4">
        <span class="text-green-600 font-bold mr-1">A</span>{{ a }}
      </p>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { ArrowDown } from "@element-plus/icons-vue";
defineProps<{ q: string; a: string }>();
const open = ref(false);
</script>

<style scoped>
.faq-enter-active,
.faq-leave-active {
  transition: all 0.25s ease;
}
.faq-enter-from,
.faq-leave-to {
  opacity: 0;
  transform: translateY(-6px);
}
</style>
