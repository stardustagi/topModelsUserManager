<!-- src/layout/components/docs/DocsLayout.vue -->
<template>
  <div class="docs-layout">
    <!-- 中间内容 -->
    <main class="docs-content">
      <router-view />
    </main>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped>
.docs-layout {
  display: flex;
  width: 100%;
}

.docs-sidebar {
  width: 280px; /* 固定宽度 */
  min-width: 280px; /* 最小宽度 */
  border-right: 1px solid #e5e7eb;
  background-color: #f9fafb;
  overflow-y: auto;
}

.docs-content {
  flex: 1; /* 占据剩余空间 */
  overflow-y: auto;
  padding: 24px;
  background-color: white;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .docs-layout {
    flex-direction: column;
    height: auto;
  }

  .docs-sidebar {
    width: 100%;
    min-width: 100%;
    border-right: none;
    border-bottom: 1px solid #e5e7eb;
  }

  .docs-content {
    width: 100%;
  }
}
</style>
