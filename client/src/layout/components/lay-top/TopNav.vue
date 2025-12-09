<!-- TopNav.vue -->
<template>
  <div class="top-nav">
    <div class="left-content">
      <!-- <img src="@/assets/logo.png" alt="Logo" class="logo" /> -->
      <span class="title">TopModels</span>
      <ul class="nav-links">
        <li><router-link to="/">首页</router-link></li>
        <li><router-link to="/console">控制台</router-link></li>
        <!-- <li><router-link to="/docs">文档</router-link></li> -->
        <!-- <li><router-link to="/">下载</router-link></li> -->
      </ul>
    </div>

    <div class="right-content">
      <!-- <div v-if="username"> -->
      <div>
        <el-dropdown trigger="click">
          <span class="el-dropdown-link navbar-bg-hover">
            <!-- <img :src="userAvatar" :style="avatarsStyle" /> -->
            <!-- <p v-if="username" class="dark:text-white">{{ username }}</p> -->
            <!-- <p>{{ useStore.username }}</p> -->
            <p v-if="username" class="dark:text-white">{{ username }}</p>
          </span>
          <template #dropdown>
            <el-dropdown-menu class="logout">
              <el-dropdown-item @click="logout">
                <IconifyIconOffline :icon="LogoutCircleRLine" style="margin: 5px" />
                <!-- {{ t("buttons.pureLoginOut") }} -->
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
      <!-- <div v-else>
        <div class="flex items-center space-x-4 ml-4">
          <el-button plain round @click="onLogin">登录</el-button>
        </div>
      </div> -->
    </div>
  </div>
</template>

<script lang="ts" setup>
const { logout, username } = useNav();

import { useNav } from "@/layout/hooks/useNav";
import { useTranslationLang } from "@/layout/hooks/useTranslationLang";
import LogoutCircleRLine from "~icons/ri/logout-circle-r-line";
// import { isLoggedIn } from "@/utils/auth";
// import { useUserStoreHook } from "@/store/modules/user";
import { useRouter } from "vue-router";
// import { computed } from "vue";
// import { getToken } from "@/utils/auth";
// const { t } = useTranslationLang();
const router = useRouter();
// const useStore = useUserStoreHook();

// const username = computed(() => useStore.username);

// const isLoggedIn = computed(() => {
//   const token = getToken();
//   console.log(useStore.username, "   000");
//   // console.log("token ==666 ", token?.accessToken);
//   return !!token?.accessToken;
// });

const onLogin = () => {
  router.push("/login");
};
</script>

<style scoped>
.top-nav {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 20px;
  /* background-color: #f8f9fa; */
  background: linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%);
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.05);
  /* box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1); */
  height: 50px;
}

.left-content {
  display: flex;
  align-items: center;
}

.logo {
  width: 30px;
  height: 30px;
  margin-right: 10px;
}

.title {
  font-size: 18px;
  font-weight: bold;
  margin-right: 20px;
}

.nav-links {
  display: flex;
  list-style: none;
  font-weight: 700;
  gap: 10px;
  margin: 0;
  padding: 0;
}

.nav-links li {
  margin-right: 20px;
}

.nav-links a {
  text-decoration: none;
  color: #333;
}

.nav-links a:hover {
  text-decoration: none;
  color: rgb(29, 29, 233);
}

.right-content {
  display: flex;
  align-items: center;
  justify-content: space-around;
  height: 48px;
  padding: 10px;
  color: #000000d9;
  cursor: pointer;

  p {
    font-size: 14px;
  }
}

.icon-button {
  background: none;
  border: none;
  cursor: pointer;
  margin-right: 10px;
}

.user-dropdown {
  display: flex;
  align-items: center;
  cursor: pointer;
}

.dropdown-icon {
  margin-left: 5px;
}
</style>
