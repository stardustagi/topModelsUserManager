import { $t } from "@/plugins/i18n";

export default {
  path: "/cansole",
  redirect: "/console/dashboard",
  meta: {
    icon: "ri/information-line",
    title: "控制台",
    rank: 10
  },
  children: [
    {
      path: "/console/dashboard",
      name: "Dashboard",
      component: () => import("@/views/console/dashboard.vue"),
      meta: {
        title: "数据看板"
      }
    }
  ]
} satisfies RouteConfigsTable;
