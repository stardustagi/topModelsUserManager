import { $t } from "@/plugins/i18n";

export default {
  path: "/person_center",
  redirect: "/person_center/package/index",
  meta: {
    icon: "ri/information-line",
    // showLink: false,
    // title: $t("menus.pureAbnormal"),
    title: "个人中心",
    rank: 9
  },
  children: [
    {
      path: "/person_center/wallet",
      name: "Wallet",
      component: () => import("@/views/person_center/wallet/index.vue"),
      meta: {
        title: "钱包",
        roles: ["admin"]
      }
    },
    {
      path: "/person_center/set",
      name: "Set",
      component: () => import("@/views/person_center/set/index.vue"),
      meta: {
        title: "个人设置"
      }
    }
  ]
} satisfies RouteConfigsTable;
