<template>
  <aside class="h-full flex flex-col bg-gray-50">
    <!-- Logo -->
    <!-- <div class="flex items-center h-14 px-4 font-bold text-blue-600 text-lg">
      TopModels
    </div> -->

    <!-- 搜索框 -->
    <div class="px-4 mb-2">
      <el-input
        v-model="keyword"
        placeholder="搜索"
        size="small"
        prefix-icon="el-icon-search"
        class="rounded"
      />
    </div>

    <!-- 导航菜单 -->
    <el-menu
      :default-active="activeMenu"
      class="flex-1 border-0"
      @select="handleSelect"
    >
      <el-menu-item index="/docs/doc/mustread">📖 (必读)使用说明</el-menu-item>

      <el-sub-menu index="openai">
        <template #title>OpenAI 接口</template>
        <el-menu-item index="/docs/doc/doctest">chat</el-menu-item>
        <el-menu-item index="/responses">responses</el-menu-item>
        <el-menu-item index="/images">images</el-menu-item>
      </el-sub-menu>

      <el-sub-menu index="midjourney">
        <template #title>Midjourney 接口</template>
        <el-menu-item index="/imagine">
          文生图 <el-tag size="small" type="danger">POST</el-tag>
        </el-menu-item>
        <el-menu-item index="/blend">
          图像混合 <el-tag size="small" type="danger">POST</el-tag>
        </el-menu-item>
        <el-menu-item index="/seed">
          获取种子 <el-tag size="small" type="success">GET</el-tag>
        </el-menu-item>
      </el-sub-menu>
    </el-menu>
  </aside>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";

const keyword = ref("");
const activeMenu = ref("/");
const router = useRouter();

const handleSelect = (index: string) => {
  router.push(index);
};
</script>
