<template>
  <div class="home-container">
    <!-- 导航栏 -->
    <!-- <el-menu
      :default-active="activeIndex"
      class="el-menu-demo bg-white shadow-lg fixed w-full z-50"
      mode="horizontal"
      @select="handleSelect"
      :ellipsis="false"
      :class="{ 'px-4': isMobile, 'px-20': !isMobile }"
    >
      <div class="flex items-center">
        <img
          src="https://via.placeholder.com/40?text=Logo"
          alt="Logo"
          class="h-12 mr-4"
        />
        <span class="text-2xl font-bold text-gray-800">TopModels AI</span>
      </div>
      <div class="flex-grow"></div>
      <el-menu-item index="1">首页</el-menu-item>
      <el-menu-item index="2">API 文档</el-menu-item>
      <el-menu-item index="3">定价</el-menu-item>
      <el-menu-item index="4">解决方案</el-menu-item>
      <el-menu-item index="5">关于我们</el-menu-item>
      <div class="flex items-center space-x-4 ml-4">
        <el-button plain round>登录</el-button>
        <el-button type="primary" round @click="scrollToSection('try-now')"
          >免费试用</el-button
        >
      </div>
    </el-menu> -->

    <!-- Hero 区域 -->
    <div
      class="hero-bg text-white pt-20 pb-20 md:pt-20 md:pb-28 text-center relative"
    >
      <div class="container mx-auto px-4 relative z-10">
        <h1 class="text-2xl md:text-5xl font-bold mb-6 scroll-reveal">
          释放 AI 潜能，赋能智能未来
        </h1>
        <p class="text-lg md:text-2xl !mb-10 !mt-6 scroll-reveal">
          TopModels AI 提供高性能、易集成的
          API，助力开发者构建智能、可靠的应用。
        </p>
        <div class="flex justify-center space-x-4">
          <el-button
            type="primary"
            size="large"
            @click="scrollToSection('try-now')"
            >立即试用</el-button
          >
          <el-button plain size="large" class="text-white border-white"
            >查看文档</el-button
          >
        </div>
      </div>
      <div
        class="absolute bottom-0 left-0 w-full h-24 bg-gradient-to-t from-gray-50 to-transparent"
      ></div>
    </div>

    <!-- 技术优势 -->
    <div class="container mx-auto px-4 py-20">
      <h2
        class="text-3xl md:text-4xl font-bold text-center text-gray-800 mb-12 scroll-reveal"
      >
        为什么选择 TopModels AI
      </h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div
          v-for="advantage in advantages"
          :key="advantage.title"
          class="feature-card bg-white p-8 rounded-xl shadow-lg scroll-reveal"
        >
          <i :class="advantage.icon" class="text-5xl text-blue-500 mb-4"></i>
          <h3 class="text-xl font-semibold text-gray-800 mb-3">
            {{ advantage.title }}
          </h3>
          <p class="text-gray-600 leading-relaxed">
            {{ advantage.description }}
          </p>
        </div>
      </div>
    </div>

    <!-- 动态统计 -->
    <div class="bg-gray-100 py-20">
      <div class="container mx-auto px-4">
        <h2
          class="text-3xl md:text-4xl font-bold text-center text-gray-800 mb-12 scroll-reveal"
        >
          我们的影响力
        </h2>
        <div class="grid grid-cols-1 md:grid-cols-4 gap-8 text-center">
          <div v-for="stat in stats" :key="stat.label" class="scroll-reveal">
            <p :id="'stat-' + stat.label" class="stats-number">
              {{ stat.value }}
            </p>
            <p class="text-gray-600">{{ stat.label }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- 案例展示 -->
    <div class="container mx-auto px-4 py-20">
      <h2
        class="text-3xl md:text-4xl font-bold text-center text-gray-800 mb-12 scroll-reveal"
      >
        成功案例
      </h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <div
          v-for="caseStudy in caseStudies"
          :key="caseStudy.title"
          class="case-card bg-white p-6 rounded-xl shadow-lg scroll-reveal"
        >
          <img
            :src="caseStudy.image"
            alt="case"
            class="w-full h-40 object-cover rounded-lg mb-4"
          />
          <h3 class="text-xl font-semibold text-gray-800 mb-2">
            {{ caseStudy.title }}
          </h3>
          <p class="text-gray-600 mb-4">{{ caseStudy.description }}</p>
          <el-button type="text" class="text-blue-500">了解更多</el-button>
        </div>
      </div>
    </div>

    <!-- API 文档入口 -->
    <div class="bg-blue-50 py-20">
      <div class="container mx-auto px-4 text-center">
        <h2
          class="text-3xl md:text-4xl font-bold text-gray-800 mb-6 scroll-reveal"
        >
          探索 API 文档
        </h2>
        <p class="text-lg text-gray-600 mb-8 max-w-2xl mx-auto scroll-reveal">
          全面的开发文档和示例代码，助力快速集成和开发。
        </p>
        <el-button type="primary" size="large" round>查看 API 文档</el-button>
      </div>
    </div>

    <!-- 客户评价 -->
    <div class="container mx-auto px-4 py-20">
      <h2
        class="text-3xl md:text-4xl font-bold text-center text-gray-800 mb-12 scroll-reveal"
      >
        客户评价
      </h2>
      <el-carousel :interval="6000" type="card" height="320px">
        <el-carousel-item
          v-for="testimonial in testimonials"
          :key="testimonial.name"
        >
          <div
            class="bg-white p-8 rounded-xl shadow-lg h-full flex flex-col justify-between"
          >
            <p class="text-gray-600 mb-4">{{ testimonial.content }}</p>
            <div class="flex items-center">
              <img
                :src="testimonial.avatar"
                alt="avatar"
                class="w-14 h-14 rounded-full mr-4"
              />
              <div>
                <p class="font-semibold text-gray-800">
                  {{ testimonial.name }}
                </p>
                <p class="text-gray-500 text-sm">{{ testimonial.role }}</p>
              </div>
            </div>
          </div>
        </el-carousel-item>
      </el-carousel>
    </div>

    <!-- 试用表单 -->
    <div id="try-now" class="bg-blue-600 text-white py-20">
      <div class="container mx-auto px-4">
        <h2
          class="text-3xl md:text-4xl font-bold text-center mb-8 scroll-reveal"
        >
          立即体验 TopModels AI
        </h2>
        <p class="text-lg text-center mb-12 max-w-xl mx-auto scroll-reveal">
          填写信息，免费试用我们的 API，开启智能化开发之旅。
        </p>
        <div class="max-w-md mx-auto bg-white p-8 rounded-xl shadow-lg">
          <el-form
            :model="form"
            :rules="rules"
            ref="formRef"
            label-position="top"
          >
            <el-form-item label="姓名" prop="name">
              <el-input v-model="form.name" placeholder="请输入您的姓名" />
            </el-form-item>
            <el-form-item label="邮箱" prop="email">
              <el-input v-model="form.email" placeholder="请输入您的邮箱" />
            </el-form-item>
            <el-form-item label="公司名称" prop="company">
              <el-input v-model="form.company" placeholder="请输入公司名称" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="submitForm" round
                >提交</el-button
              >
            </el-form-item>
          </el-form>
        </div>
      </div>
    </div>

    <!-- 页脚 -->
    <footer class="bg-gray-900 text-white py-16">
      <div class="container mx-auto px-4">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-12">
          <div>
            <h3 class="text-lg font-semibold mb-4">TopModels AI</h3>
            <p class="text-gray-400">
              领先的 AI API 平台，赋能开发者构建智能未来。
            </p>
            <div class="flex space-x-4 mt-4">
              <a href="#" class="text-gray-400 hover:text-white"
                ><i class="fab fa-weixin"></i
              ></a>
              <a href="#" class="text-gray-400 hover:text-white"
                ><i class="fab fa-weibo"></i
              ></a>
              <a href="#" class="text-gray-400 hover:text-white"
                ><i class="fab fa-github"></i
              ></a>
              <a href="#" class="text-gray-400 hover:text-white"
                ><i class="fab fa-twitter"></i
              ></a>
            </div>
          </div>
          <div>
            <h3 class="text-lg font-semibold mb-4">产品</h3>
            <ul class="space-y-2">
              <li>
                <a href="#" class="text-gray-400 hover:text-white">API 文档</a>
              </li>
              <li>
                <a href="#" class="text-gray-400 hover:text-white">定价</a>
              </li>
              <li>
                <a href="#" class="text-gray-400 hover:text-white"
                  >开发者社区</a
                >
              </li>
              <li>
                <a href="#" class="text-gray-400 hover:text-white">SDK 下载</a>
              </li>
            </ul>
          </div>
          <div>
            <h3 class="text-lg font-semibold mb-4">资源</h3>
            <ul class="space-y-2">
              <li>
                <a href="#" class="text-gray-400 hover:text-white">博客</a>
              </li>
              <li>
                <a href="#" class="text-gray-400 hover:text-white">教程</a>
              </li>
              <li>
                <a href="#" class="text-gray-400 hover:text-white">常见问题</a>
              </li>
              <li>
                <a href="#" class="text-gray-400 hover:text-white">支持中心</a>
              </li>
            </ul>
          </div>
          <div>
            <h3 class="text-lg font-semibold mb-4">联系我们</h3>
            <p class="text-gray-400">邮箱: support@TopModels-ai.com</p>
            <p class="text-gray-400">电话: +86 123-456-7890</p>
            <p class="text-gray-400">地址: 中国上海科技园区</p>
          </div>
        </div>
        <div class="mt-12 text-center text-gray-400">
          <p>&copy; 2025 TopModels AI. All rights reserved.</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from "vue";
import { ElMessage, ElForm, FormRules } from "element-plus";
import gsap from "gsap";
import ScrollTrigger from "gsap/ScrollTrigger";

gsap.registerPlugin(ScrollTrigger);

// 类型定义
interface Advantage {
  title: string;
  description: string;
  icon: string;
}

interface Stat {
  label: string;
  value: number;
  target: number;
}

interface CaseStudy {
  title: string;
  description: string;
  image: string;
}

interface Testimonial {
  name: string;
  role: string;
  content: string;
  avatar: string;
}

interface Form {
  name: string;
  email: string;
  company: string;
}

// 数据
const activeIndex = ref("1");
const isMobile = ref(false);
const formRef = ref<InstanceType<typeof ElForm>>();
const form = ref<Form>({
  name: "",
  email: "",
  company: ""
});
const rules: FormRules = {
  name: [{ required: true, message: "请输入姓名", trigger: "blur" }],
  email: [
    { required: true, message: "请输入邮箱", trigger: "blur" },
    { type: "email", message: "请输入有效邮箱", trigger: "blur" }
  ],
  company: [{ required: true, message: "请输入公司名称", trigger: "blur" }]
};
const advantages = ref<Advantage[]>([
  {
    title: "高性能 AI",
    description: "基于先进模型，提供毫秒级响应，满足高并发需求。",
    icon: "fas fa-rocket"
  },
  {
    title: "易于集成",
    description: "简洁的 API 设计和详细文档，快速上手，轻松集成。",
    icon: "fas fa-plug"
  },
  {
    title: "安全可靠",
    description: "多层加密和隐私保护，确保数据安全无忧。",
    icon: "fas fa-shield-alt"
  }
]);
const stats = ref<Stat[]>([
  { label: "活跃用户", value: 0, target: 100000 },
  { label: "API 调用", value: 0, target: 5000000 },
  { label: "合作伙伴", value: 0, target: 200 },
  { label: "国家覆盖", value: 0, target: 50 }
]);
const caseStudies = ref<CaseStudy[]>([
  {
    title: "智能客服系统",
    description: "为电商平台提供 24/7 智能客服，降低 40% 运营成本。",
    image: "https://via.placeholder.com/300x200?text=Case1"
  },
  {
    title: "情感分析工具",
    description: "帮助品牌分析用户反馈，提升产品满意度 30%。",
    image: "https://via.placeholder.com/300x200?text=Case2"
  },
  {
    title: "自动化内容生成",
    description: "为媒体公司提供 AI 内容生成，效率提升 3 倍。",
    image: "https://via.placeholder.com/300x200?text=Case3"
  }
]);
const testimonials = ref<Testimonial[]>([
  {
    name: "张伟",
    role: "CTO, 智能科技公司",
    content: "TopModels AI 的 API 性能卓越，帮助我们快速构建了智能客服系统。",
    avatar: "https://via.placeholder.com/48?text=ZW"
  },
  {
    name: "李娜",
    role: "产品经理, 创新企业",
    content: "文档清晰、集成简单，极大缩短了我们的开发周期。",
    avatar: "https://via.placeholder.com/48?text=LN"
  },
  {
    name: "王磊",
    role: "开发者, 创业公司",
    content: "支持团队响应迅速，API 稳定性让我们非常放心。",
    avatar: "https://via.placeholder.com/48?text=WL"
  }
]);

// 方法
const handleSelect = (key: string) => {
  activeIndex.value = key;
};

const scrollToSection = (id: string) => {
  document.getElementById(id)?.scrollIntoView({ behavior: "smooth" });
};

const checkMobile = () => {
  isMobile.value = window.innerWidth <= 768;
};

const submitForm = () => {
  formRef.value?.validate(valid => {
    if (valid) {
      ElMessage.success("提交成功，我们将尽快联系您！");
      form.value = { name: "", email: "", company: "" };
    }
  });
};

const handleScroll = () => {
  const elements = document.querySelectorAll(".scroll-reveal");
  elements.forEach(el => {
    const rect = el.getBoundingClientRect();
    if (rect.top <= window.innerHeight * 0.8) {
      el.classList.add("visible");
    }
  });
};

const animateStats = () => {
  stats.value.forEach(stat => {
    gsap.to(`#stat-${stat.label}`, {
      duration: 2,
      innerText: stat.target,
      roundProps: "innerText",
      ease: "power1.out",
      scrollTrigger: {
        trigger: `#stat-${stat.label}`,
        start: "top 80%"
      }
    });
  });
};

// 生命周期
onMounted(() => {
  checkMobile();
  window.addEventListener("resize", checkMobile);
  window.addEventListener("scroll", handleScroll);
  handleScroll();
  animateStats();
  gsap.from(".hero-bg", {
    opacity: 0,
    duration: 1,
    ease: "power2.out"
  });
});

onBeforeUnmount(() => {
  window.removeEventListener("resize", checkMobile);
  window.removeEventListener("scroll", handleScroll);
});
</script>

<style lang="scss" scoped>
.home-container {
  font-family:
    "Inter",
    -apple-system,
    BlinkMacSystemFont,
    sans-serif;
}

.hero-bg {
  background: linear-gradient(135deg, #1e40af 0%, #3b82f6 100%);
  position: relative;
  overflow: hidden;

  &::before {
    content: "";
    position: absolute;
    top: -50%;
    left: -50%;
    width: 200%;
    height: 200%;
    background: radial-gradient(
      circle,
      rgba(255, 255, 255, 0.15) 0%,
      transparent 70%
    );
    animation: rotate 20s linear infinite;
    z-index: 0;
  }
}

@keyframes rotate {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

.feature-card,
.case-card {
  transition:
    transform 0.3s ease,
    box-shadow 0.3s ease;

  &:hover {
    transform: translateY(-10px);
    box-shadow: 0 12px 24px rgba(0, 0, 0, 0.15);
  }
}

.scroll-reveal {
  opacity: 0;
  transform: translateY(30px);
  transition:
    opacity 0.6s ease,
    transform 0.6s ease;

  &.visible {
    opacity: 1;
    transform: translateY(0);
  }
}

.stats-number {
  font-size: 2.5rem;
  font-weight: bold;
  color: #1e40af;
}

@media (max-width: 768px) {
  .hero-bg::before {
    background: none;
  }

  .stats-number {
    font-size: 2rem;
  }
}
</style>
