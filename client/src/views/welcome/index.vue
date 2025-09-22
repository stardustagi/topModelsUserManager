<template>
  <div
    class="bg-gradient-to-b from-gray-50 via-white to-gray-100 text-gray-800"
  >
    <!-- Hero -->
    <section class="relative overflow-hidden">
      <div class="relative z-10 max-w-6xl mx-auto px-6 py-24 text-center">
        <span
          class="inline-flex items-center gap-2 px-4 py-1 rounded-full text-sm bg-purple-100 text-purple-600 shadow-sm mb-5 animate-pulse"
        >
          🚀 现已接入 GPT!
        </span>
        <h1
          class="text-6xl font-extrabold tracking-wide text-transparent bg-clip-text bg-gradient-to-r from-purple-600 to-indigo-600"
        >
          TopModels API
        </h1>
        <p class="mt-6 text-lg text-gray-600">
          一站式 AI
          模型接口服务，支持国内外主流大模型，官方接口，稳定可靠，开箱即用
        </p>
        <div class="mt-10 flex items-center justify-center gap-4">
          <el-button
            type="primary"
            size="large"
            class="!px-8 !py-5 !text-lg rounded-full shadow-lg hover:shadow-purple-400/40"
            @click="toConsole"
            >立即开始 →</el-button
          >
          <!-- <el-button
            plain
            size="large"
            class="!px-8 !py-5 !text-lg rounded-full border-2 hover:border-purple-500 hover:text-purple-600"
            >了解更多</el-button
          > -->
        </div>
      </div>
    </section>

    <!-- 支持的模型 -->
    <SectionHeader
      title="支持的模型"
      subtitle="覆盖全球主流 AI 模型，满足不同场景需求"
    />
    <div class="grid md:grid-cols-3 gap-6 max-w-7xl mx-auto px-6 pb-20">
      <FeatureCard
        icon="🌍"
        title="海外模型"
        desc="全球领先的 AI 模型"
        :tags="['GPT-5', 'Claude 4', 'Gemini Pro']"
      />
      <FeatureCard
        icon="🏠"
        title="国内模型"
        desc="本土优化模型"
        :tags="['ChatGLM', '讯飞星火', 'Kimi Chat']"
      />
      <FeatureCard
        icon="⭐"
        title="专属资源"
        desc="高性能集群"
        :tags="['DeepSeek-R1-32B']"
        badge="NEW"
      />
    </div>

    <!-- 为什么选择我们 -->
    <SectionHeader
      title="为什么选择TopModels API"
      subtitle="专业、稳定、高效的 AI 接口服务"
    />
    <div class="grid md:grid-cols-4 gap-6 max-w-7xl mx-auto px-6 pb-20">
      <FeatureCard icon="🛡️" title="稳定可靠" desc="99.9% 服务可用性保证" />
      <FeatureCard icon="💰" title="价格优惠" desc="大幅节省成本" />
      <FeatureCard icon="🔌" title="兼容性强" desc="完全兼容 OpenAI API" />
      <FeatureCard icon="👨‍💻" title="专业支持" desc="7×24 技术支持" />
    </div>

    <!-- 开发资源 -->
    <SectionHeader title="开发资源" subtitle="从入门到精通，全方位开发支持" />
    <div class="grid md:grid-cols-3 gap-6 max-w-7xl mx-auto px-6 pb-20">
      <ResourceCard
        icon="⚡"
        title="快速开始"
        desc="5分钟接入，含示例代码"
        tag="入门"
      />
      <ResourceCard icon="📊" title="服务监控" desc="实时 API 状态与性能指标" />
      <ResourceCard
        icon="💻"
        title="代码示例"
        desc="一键复制，立即使用"
        dark
        tag="热门"
      />
      <ResourceCard icon="📖" title="API 文档" desc="完整接口文档与参数说明" />
      <ResourceCard
        icon="🤝"
        title="开发者社区"
        desc="1000+ 开发者交流与经验分享"
        tag="推荐"
      />
      <ResourceCard
        icon="🛠️"
        title="开发工具"
        desc="SDK 与调试工具，提升效率"
      />
    </div>

    <!-- 代理合作优势 -->
    <SectionHeader title="代理合作计划" subtitle="加入我们，共享 AI 时代红利" />
    <div class="grid md:grid-cols-3 gap-6 max-w-7xl mx-auto px-6 pb-20">
      <FeatureCard
        icon="🤝"
        title="优质合作"
        desc="丰厚返佣比例，客户越多收益越高"
      />
      <FeatureCard
        icon="📞"
        title="专属支持"
        desc="1v1 客服经理 + 技术培训支持"
      />
      <FeatureCard
        icon="📈"
        title="数据透明"
        desc="可视化数据看板，返佣明细清晰"
      />
    </div>

    <!-- 代理等级 -->
    <SectionHeader title="代理等级与权益" subtitle="不同等级，不同奖励与支持" />
    <div class="grid md:grid-cols-4 gap-6 max-w-7xl mx-auto px-6 pb-20">
      <LevelCard icon="🥉" title="铜牌代理" desc="基础宣传素材、月度奖励" />
      <LevelCard icon="🥈" title="银牌代理" desc="优先支持、季度奖励" />
      <LevelCard icon="🥇" title="金牌代理" desc="专属品牌授权、年度大会" />
      <LevelCard
        icon="💎"
        title="钻石代理"
        desc="区域独家、定制化方案与高层对接"
      />
    </div>

    <!-- 加入流程 -->
    <SectionHeader title="加入流程" />
    <div class="grid md:grid-cols-4 gap-6 max-w-6xl mx-auto px-6 pb-20">
      <StepCard num="1" title="提交申请" desc="填写代理申请表" />
      <StepCard num="2" title="资质审核" desc="1-2 个工作日内完成" />
      <StepCard num="3" title="签约培训" desc="签署协议并完成培训" />
      <StepCard num="4" title="开始推广" desc="获取物料，正式开始合作" />
    </div>

    <!-- FAQ -->
    <SectionHeader title="常见问题" subtitle="帮助你快速了解代理与服务" />
    <div class="grid md:grid-cols-2 gap-6 max-w-6xl mx-auto px-6 pb-20">
      <FaqCard
        q="成为代理有什么要求？"
        a="无需资质要求，个人与企业均可申请。"
      />
      <FaqCard
        q="返佣如何计算和结算？"
        a="按订单或消费额比例结算，按月自动结算。"
      />
      <FaqCard
        q="需要预付代理费用吗？"
        a="完全免费，零成本加入，我们提供支持。"
      />
      <FaqCard
        q="客户停止使用后还有返佣吗？"
        a="返佣基于实际消费，客户停止使用后返佣随之停止。"
      />
    </div>

    <!-- 价格方案 -->
    <SectionHeader title="价格方案" subtitle="透明定价，按需付费" />
    <div class="grid md:grid-cols-3 gap-6 max-w-7xl mx-auto px-6 pb-20">
      <PricingCard
        icon="📦"
        title="按量计费"
        price="¥0.012/1K tokens"
        desc="100+ 模型支持 · 官方价低折扣"
        cta="立即开始"
        highlight
      />
      <PricingCard
        icon="🎁"
        title="新用户福利"
        price="¥0 免费额度"
        desc="注册即送10元额度"
        cta="免费试用"
      />
      <PricingCard
        icon="🏢"
        title="企业定制"
        price="联系商务"
        desc="专属API密钥 · SLA保障"
        cta="联系我们"
      />
    </div>

    <!-- Footer -->
    <footer class="bg-gray-900 text-gray-400">
      <div class="max-w-7xl mx-auto px-6 py-12 grid md:grid-cols-3 gap-6">
        <div>
          <h3 class="text-white font-bold text-xl mb-2">TopModels API</h3>
          <p>下一代 AI 接口服务平台</p>
        </div>
        <div>
          <h4 class="text-white font-semibold mb-2">服务</h4>
          <ul>
            <li>立即注册</li>
            <li>充值购买</li>
            <li>价格方案</li>
          </ul>
        </div>
        <div>
          <h4 class="text-white font-semibold mb-2">文档</h4>
          <ul>
            <li>使用教程</li>
            <li>服务状态</li>
          </ul>
        </div>
      </div>
      <p class="text-center text-gray-500 text-sm pb-8">
        © 2025 TopModels API. All rights reserved.
      </p>
    </footer>
  </div>
</template>

<script setup lang="ts">
import SectionHeader from "./SectionHeader.vue";
import FeatureCard from "./FeatureCard.vue";
import ResourceCard from "./ResourceCard.vue";
import LevelCard from "./LevelCard.vue";
import StepCard from "./StepCard.vue";
import FaqCard from "./FaqCard.vue";
import PricingCard from "./PricingCard.vue";
import { useRouter } from "vue-router";

const router = useRouter();
const toConsole = () => {
  router.push("/console");
};
</script>
