<!-- App.vue -->
<template>
  <div class="min-h-screen bg-dark-bg text-dark-text relative">
    <!-- 背景装饰 -->
    <div class="fixed inset-0 -z-10 overflow-hidden">
      <div
        class="absolute -top-40 -right-40 w-80 h-80 bg-primary rounded-full opacity-10 blur-3xl"
      ></div>
      <div
        class="absolute -bottom-40 -left-40 w-80 h-80 bg-primary-dark rounded-full opacity-10 blur-3xl"
      ></div>
    </div>

    <!-- 导航栏 -->
    <nav
      class="bg-dark-bg/90 backdrop-blur-md border-b border-gray-800 fixed w-full top-0 z-50"
    >
      <div class="max-w-7xl mx-auto px-6 lg:px-8">
        <div class="flex justify-between h-16 items-center">
          <!-- Logo -->
          <div class="flex-shrink-0 flex items-center">
            <div class="text-primary text-2xl font-bold">TopAPI</div>
          </div>

          <!-- 导航链接 -->
          <div class="hidden md:flex items-center space-x-8">
            <a
              href="#features"
              class="hover:text-primary transition-colors duration-200"
              >API</a
            >
            <a
              href="#how-it-works"
              class="hover:text-primary transition-colors duration-200"
              >文档</a
            >
            <a
              href="#pricing"
              class="hover:text-primary transition-colors duration-200"
              >定价</a
            >
            <a
              href="#console"
              class="hover:text-primary transition-colors duration-200"
              >控制台</a
            >
          </div>

          <!-- 按钮 -->
          <div class="flex items-center space-x-4">
            <button
              class="text-dark-text-secondary hover:text-white transition-colors duration-200"
            >
              登录
            </button>
            <button
              class="bg-primary hover:bg-primary-dark text-white px-4 py-2 rounded-lg transition-all duration-200 transform hover:scale-105"
            >
              注册
            </button>
          </div>
        </div>
      </div>
    </nav>

    <!-- 英雄区域 -->
    <section class="pt-24 pb-16 md:pt-32 md:pb-24 px-6 lg:px-8">
      <div class="max-w-7xl mx-auto text-center">
        <h1
          class="text-4xl md:text-6xl font-extrabold mb-6 bg-gradient-to-r from-white via-gray-200 to-gray-300 bg-clip-text text-transparent"
        >
          释放 API 的无限潜能
        </h1>
        <p
          class="text-xl md:text-2xl text-dark-text-secondary mb-10 max-w-3xl mx-auto leading-relaxed"
        >
          为开发者提供强大、稳定、易用的 API 服务，加速您的应用开发进程。
        </p>
        <div class="space-x-4">
          <button
            @click="scrollToSection('features')"
            class="bg-primary hover:bg-primary-dark text-white px-8 py-3 rounded-lg text-lg font-medium transition-all duration-200 transform hover:scale-105 shadow-lg hover:shadow-primary/25"
          >
            立即开始
          </button>
          <button
            class="border border-gray-600 hover:border-gray-400 text-white px-8 py-3 rounded-lg text-lg font-medium transition-all duration-200 hover:bg-gray-800"
          >
            查看文档
          </button>
        </div>

        <!-- 产品展示图 -->
        <div
          class="mt-16 bg-dark-card rounded-xl p-6 shadow-2xl border border-gray-700 max-w-5xl mx-auto transform hover:scale-105 transition-transform duration-300"
        >
          <div
            class="h-64 md:h-80 bg-gradient-to-br from-gray-800 to-gray-900 rounded-lg flex items-center justify-center text-gray-400 text-lg"
          >
            🖥️ API 控制台模拟界面
          </div>
        </div>
      </div>
    </section>

    <!-- 核心特性 -->
    <section id="features" class="py-16 px-6 lg:px-8 bg-dark-card/50">
      <div class="max-w-7xl mx-auto">
        <div class="text-center mb-16">
          <h2 class="text-3xl md:text-4xl font-bold mb-4">核心特性</h2>
          <p class="text-dark-text-secondary text-lg">
            我们提供业界领先的 API 服务，满足您的各种需求
          </p>
        </div>

        <div class="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
          <div
            v-for="feature in features"
            :key="feature.title"
            class="bg-dark-bg/70 p-8 rounded-xl shadow-lg border border-gray-700 hover:shadow-xl hover:border-primary/30 transition-all duration-300 group"
          >
            <div
              class="text-4xl mb-4 transform group-hover:scale-110 transition-transform duration-300"
            >
              {{ feature.icon }}
            </div>
            <h3 class="text-xl font-semibold mb-3">{{ feature.title }}</h3>
            <p class="text-dark-text-secondary leading-relaxed">
              {{ feature.description }}
            </p>
          </div>
        </div>
      </div>
    </section>

    <!-- 如何使用 -->
    <section id="how-it-works" class="py-16 px-6 lg:px-8">
      <div class="max-w-7xl mx-auto">
        <div class="text-center mb-16">
          <h2 class="text-3xl md:text-4xl font-bold mb-4">如何使用</h2>
          <p class="text-dark-text-secondary text-lg">
            三步快速接入我们的 API 服务
          </p>
        </div>

        <div class="grid md:grid-cols-3 gap-8">
          <div v-for="(step, index) in steps" :key="index" class="text-center">
            <div
              class="w-16 h-16 bg-primary text-white rounded-full flex items-center justify-center text-2xl font-bold mx-auto mb-6"
            >
              {{ index + 1 }}
            </div>
            <h3 class="text-xl font-semibold mb-3">{{ step.title }}</h3>
            <p class="text-dark-text-secondary">{{ step.description }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- 定价 -->
    <section id="pricing" class="py-16 px-6 lg:px-8 bg-dark-card/50">
      <div class="max-w-7xl mx-auto">
        <div class="text-center mb-16">
          <h2 class="text-3xl md:text-4xl font-bold mb-4">定价计划</h2>
          <p class="text-dark-text-secondary text-lg">选择最适合您需求的方案</p>
        </div>

        <div class="grid md:grid-cols-3 gap-8 max-w-5xl mx-auto">
          <div
            v-for="plan in pricingPlans"
            :key="plan.name"
            class="bg-dark-bg/70 rounded-xl p-8 shadow-xl border border-gray-700 transition-all duration-300 hover:shadow-2xl hover:border-primary/40"
          >
            <h3 class="text-2xl font-bold mb-2">{{ plan.name }}</h3>
            <div class="text-4xl font-bold mb-6">{{ plan.price }}</div>
            <p class="text-dark-text-secondary mb-6">{{ plan.description }}</p>
            <ul class="space-y-3 mb-8">
              <li
                v-for="feature in plan.features"
                :key="feature"
                class="flex items-center"
              >
                <svg
                  class="w-5 h-5 text-green-500 mr-2"
                  fill="currentColor"
                  viewBox="0 0 20 20"
                >
                  <path
                    fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd"
                  ></path>
                </svg>
                {{ feature }}
              </li>
            </ul>
            <button
              class="w-full bg-primary hover:bg-primary-dark text-white py-3 rounded-lg transition-colors duration-200 transform hover:scale-105"
            >
              {{ plan.buttonText }}
            </button>
          </div>
        </div>
      </div>
    </section>

    <!-- 页脚 -->
    <footer
      id="console"
      class="bg-dark-bg border-t border-gray-800 py-12 px-6 lg:px-8"
    >
      <div class="max-w-7xl mx-auto">
        <div class="grid grid-cols-2 md:grid-cols-4 gap-8 mb-8">
          <div>
            <h3 class="text-lg font-semibold mb-4">产品</h3>
            <ul class="space-y-2">
              <li>
                <a
                  href="#"
                  class="text-dark-text-secondary hover:text-white transition-colors duration-200"
                  >API 列表</a
                >
              </li>
              <li>
                <a
                  href="#"
                  class="text-dark-text-secondary hover:text-white transition-colors duration-200"
                  >SDK</a
                >
              </li>
              <li>
                <a
                  href="#"
                  class="text-dark-text-secondary hover:text-white transition-colors duration-200"
                  >定价</a
                >
              </li>
            </ul>
          </div>
          <div>
            <h3 class="text-lg font-semibold mb-4">开发者</h3>
            <ul class="space-y-2">
              <li>
                <a
                  href="#"
                  class="text-dark-text-secondary hover:text-white transition-colors duration-200"
                  >文档</a
                >
              </li>
              <li>
                <a
                  href="#"
                  class="text-dark-text-secondary hover:text-white transition-colors duration-200"
                  >状态页</a
                >
              </li>
              <li>
                <a
                  href="#"
                  class="text-dark-text-secondary hover:text-white transition-colors duration-200"
                  >博客</a
                >
              </li>
            </ul>
          </div>
          <div>
            <h3 class="text-lg font-semibold mb-4">公司</h3>
            <ul class="space-y-2">
              <li>
                <a
                  href="#"
                  class="text-dark-text-secondary hover:text-white transition-colors duration-200"
                  >关于我们</a
                >
              </li>
              <li>
                <a
                  href="#"
                  class="text-dark-text-secondary hover:text-white transition-colors duration-200"
                  >联系我们</a
                >
              </li>
              <li>
                <a
                  href="#"
                  class="text-dark-text-secondary hover:text-white transition-colors duration-200"
                  >隐私政策</a
                >
              </li>
            </ul>
          </div>
          <div>
            <h3 class="text-lg font-semibold mb-4">关注我们</h3>
            <div class="flex space-x-4">
              <a
                href="#"
                class="text-dark-text-secondary hover:text-white transition-colors duration-200"
                >𝕏</a
              >
              <a
                href="#"
                class="text-dark-text-secondary hover:text-white transition-colors duration-200"
                >GitHub</a
              >
            </div>
          </div>
        </div>
        <div
          class="border-t border-gray-800 pt-8 text-center text-dark-text-secondary"
        >
          &copy; {{ new Date().getFullYear() }} TopAPI. All rights reserved.
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref } from "vue";

// 特性数据
const features = ref([
  {
    icon: "⚡",
    title: "高性能",
    description: "毫秒级响应，99.99% 可用性，确保您的应用流畅运行。"
  },
  {
    icon: "🔒",
    title: "高安全",
    description: "采用 OAuth 2.0 和 HTTPS 加密，保护您的数据安全。"
  },
  {
    icon: "📊",
    title: "可扩展",
    description: "轻松应对流量高峰，支持弹性伸缩。"
  },
  {
    icon: "📚",
    title: "丰富文档",
    description: "详尽的 API 文档和 SDK，快速上手集成。"
  },
  {
    icon: "🛠️",
    title: "开发者友好",
    description: "直观的控制台，实时监控和调试工具。"
  },
  {
    icon: "🌍",
    title: "全球覆盖",
    description: "全球多个节点，降低延迟，提升用户体验。"
  }
]);

// 使用步骤
const steps = ref([
  {
    title: "注册账号",
    description: "创建您的 TopAPI 账户，获取专属 API 密钥。"
  },
  {
    title: "集成 API",
    description: "使用我们提供的 SDK 或直接调用 API 接口。"
  },
  {
    title: "开始使用",
    description: "在您的应用中调用 API，享受高效服务。"
  }
]);

// 定价计划
const pricingPlans = ref([
  {
    name: "免费版",
    price: "¥0",
    description: "适合个人项目和学习使用",
    features: ["1000 次/月调用", "基础 API 访问", "社区支持"],
    buttonText: "立即使用"
  },
  {
    name: "专业版",
    price: "¥99/月",
    description: "适合中小型项目",
    features: ["50,000 次/月调用", "所有 API 访问", "优先支持", "高级分析"],
    buttonText: "升级专业版"
  },
  {
    name: "企业版",
    price: "定制",
    description: "适合大型企业和高需求项目",
    features: [
      "无限次调用",
      "专属 API 端点",
      "24/7 专家支持",
      "SLA 保障",
      "私有部署"
    ],
    buttonText: "联系销售"
  }
]);

// 滚动到指定区域
const scrollToSection = id => {
  const element = document.getElementById(id);
  if (element) {
    element.scrollIntoView({ behavior: "smooth" });
  }
};
</script>

<style>
@import url("https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700;800&display=swap");

body {
  margin: 0;
  font-family:
    "Inter",
    -apple-system,
    BlinkMacSystemFont,
    "Segoe UI",
    Roboto,
    sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

/* 自定义颜色变量 */
/* :root {
  --dark-bg: #0a1128;
  --dark-card: #111827;
  --dark-text: #e5e7eb;
  --dark-text-secondary: #9ca3af;
  --primary: #0066ff;
  --primary-dark: #0052cc;
} */

.bg-dark-bg {
  background-color: var(--dark-bg);
}
.bg-dark-card {
  background-color: var(--dark-card);
}
.text-dark-text {
  color: var(--dark-text);
}
.text-dark-text-secondary {
  color: var(--dark-text-secondary);
}
.text-primary {
  color: var(--primary);
}
.bg-primary {
  background-color: var(--primary);
}
.hover\:bg-primary-dark:hover {
  background-color: var(--primary-dark);
}
.border-primary\/30 {
  border-color: rgba(0, 102, 255, 0.3);
}
.border-primary\/40 {
  border-color: rgba(0, 102, 255, 0.4);
}
</style>
