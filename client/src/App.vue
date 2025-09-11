<template>
  <el-config-provider :locale="currentLocale">
    <TopNav />
    <div class="content-container">
      <router-view />
    </div>
    <ReDialog />
  </el-config-provider>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { ElConfigProvider } from "element-plus";
import { ReDialog } from "@/components/ReDialog";
import en from "element-plus/es/locale/lang/en";
import zhCn from "element-plus/es/locale/lang/zh-cn";
import TopNav from "./layout/components/lay-top/TopNav.vue";
// import LayNavBar from "./layout/components/lay-navbar/index.vue";

export default defineComponent({
  name: "app",
  components: {
    [ElConfigProvider.name]: ElConfigProvider,
    ReDialog,
    TopNav
  },
  computed: {
    currentLocale() {
      return this.$storage.locale?.locale === "zh" ? zhCn : en;
    }
  }
});
</script>

<style scoped>
.content-container {
  padding-top: 60px;
}
</style>
