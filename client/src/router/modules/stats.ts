import { $t } from "@/plugins/i18n";
const Layout = () => import("@/layout/index.vue");

export default {
  path: "/stats",
  name: "Stats",
  component: Layout,
  redirect: "/stats/index",
  meta: {
    icon: "ri/information-line",
    title: "统计服务",
    rank: 9,
    showLink: false
  },
  children: [
    {
      path: "index",
      name: "StatsOverview",
      component: () => import("@/views/stats/index.vue"),
      meta: {
        title: "数据总览"
      }
    },
    {
      path: "budgetalert",
      name: "BudgetAlert",
      component: () => import("@/views/stats/budgetalert.vue"),
      meta: {
        title: "预算预警"
        // roles: ["admin"],
      }
    },
    {
      path: "apikeysummary",
      name: "ApiKeySummary",
      component: () => import("@/views/stats/apikeysummary.vue"),
      meta: {
        title: "ApiKey汇总"
        // roles: ["admin"],
      }
    }
  ]
} satisfies RouteConfigsTable;
