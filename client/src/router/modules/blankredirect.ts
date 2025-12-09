import { $t } from "@/plugins/i18n";
const { VITE_HIDE_HOME } = import.meta.env;

export default {
  path: "/blank_redirect",
  name: "BlankDirectHome",
  // component: Layout,
  component: () => import("@/views/blank_redirect/index.vue"),
  meta: {
    icon: "ep/home-filled",
    title: $t("menus.pureHome"),
    rank: 0,
    showLink: false
  }
  //   children: [
  //     {
  //       path: "/blank_redirect",
  //       name: "BlankDirect",
  //       component: () => import("@/views/blank_redirect/index.vue"),
  //       meta: {
  //         title: $t("menus.pureHome"),
  //         showLink: VITE_HIDE_HOME === "true" ? false : true
  //       }
  //     }
  //   ]
} satisfies RouteConfigsTable;
