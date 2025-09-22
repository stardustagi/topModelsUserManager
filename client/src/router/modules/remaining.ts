import { $t } from "@/plugins/i18n";
const Layout = () => import("@/layout/index.vue");

export default [
  {
    path: "/login",
    name: "Login",
    component: () => import("@/views/login/index.vue"),
    meta: {
      title: $t("menus.pureLogin"),
      showLink: false,
      rank: 101
    }
  },
  {
    path: "/redirect",
    component: Layout,
    meta: {
      title: $t("status.pureLoad"),
      showLink: false,
      rank: 102
    },
    children: [
      {
        path: "/redirect/:path(.*)",
        name: "Redirect",
        component: () => import("@/layout/redirect.vue")
      }
    ]
  },
  {
    path: "/docs",
    name: "Document",
    component: Layout,
    redirect: "/docs/doc/mustread",
    meta: {
      icon: "ri/information-line",
      title: "使用说明",
      rank: 10,
      showLink: false
    },
    children: [
      // {
      //   path: "docs",
      //   name: "Docs",
      //   component: () => import("@/views/docs/docs.vue"),
      //   meta: {
      //     title: "使用说明"
      //   }
      // },
      {
        path: "doc/mustread",
        name: "Mustread",
        component: () => import("@/views/docs/doc/mustread.vue"),
        meta: {
          title: "使用说明"
        }
      },
      {
        path: "doc/doctest",
        name: "Doctest",
        component: () => import("@/views/docs/doc/doctest.vue"),
        meta: {
          title: "测试"
        }
      }
    ]
  }
] satisfies Array<RouteConfigsTable>;
