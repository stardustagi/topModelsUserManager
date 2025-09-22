import { $t } from "@/plugins/i18n";

export default {
  path: "/company",
  redirect: "/company/models/index",
  meta: {
    icon: "ri/information-line",
    // showLink: false,
    // title: $t("menus.pureAbnormal"),
    title: "企业",
    rank: 11
  },
  children: [
    {
      path: "/company/models/index",
      name: "UpModel",
      component: () => import("@/views/company/models/index.vue"),
      meta: {
        title: "上传模型",
        roles: ["editor"]
      }
    }
  ]
} satisfies RouteConfigsTable;
