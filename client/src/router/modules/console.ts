import { $t } from "@/plugins/i18n";
const Layout = () => import("@/layout/index.vue");

export default {
  path: "/console",
  name: "Console",
  component: Layout,
  redirect: "/console/dashboard",
  meta: {
    icon: "ri/information-line",
    title: "控制台",
    rank: 10
  },
  children: [
    {
      path: "dashboard",
      name: "Dashboard",
      component: () => import("@/views/console/dashboard.vue"),
      meta: {
        title: "数据看板"
      }
    },
    {
      path: "person_center/wallet",
      name: "Wallet",
      component: () => import("@/views/console/person_center/wallet/index.vue"),
      meta: {
        title: "钱包",
        roles: ["admin"],
        showLink: false
      }
    },
    {
      path: "person_center/keymanager",
      name: "KeyMgr",
      component: () =>
        import("@/views/console/person_center/keymanager/index.vue"),
      meta: {
        title: "令牌管理"
      }
    },
    {
      path: "person_center/set",
      name: "Set",
      component: () => import("@/views/console/person_center/set/index.vue"),
      meta: {
        title: "个人设置"
      }
    }
  ]
} satisfies RouteConfigsTable;
