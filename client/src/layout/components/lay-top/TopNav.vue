<!-- TopNav.vue -->
<template>
  <div class="top-nav">
    <div class="left-content">
      <!-- <img src="@/assets/logo.png" alt="Logo" class="logo" /> -->
      <span class="title">TopModels</span>
      <ul class="nav-links">
        <li><router-link to="/">首页</router-link></li>
        <li><router-link to="/console">控制台</router-link></li>
        <li><router-link to="/">定价</router-link></li>
        <li><router-link to="/">文档</router-link></li>
        <li><router-link to="/">关于</router-link></li>
      </ul>
    </div>
    <div class="right-content">
      <button class="icon-button"><i class="bell-icon"></i></button>
      <button class="icon-button"><i class="dark-mode-icon"></i></button>
      <button class="icon-button"><i class="language-icon"></i></button>
      <div class="user-dropdown">
        <span>yuanmu</span>
        <i class="dropdown-icon"></i>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup></script>

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
  background-color: #f8f9fa;
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

.right-content {
  display: flex;
  align-items: center;
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
