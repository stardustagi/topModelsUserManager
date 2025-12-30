const Layout = () => import("@/layout/index.vue");

export default {
  path: "/user_consumes",
  component: Layout,
  redirect: "/user_consumes/index",
  meta: {
    icon: "ri/information-line",
    title: "消费明细",
    rank: 11,
    showLink: true
  },
  children: [
    {
      path: "index",
      name: "UserConsumes",
      component: () => import("@/views/user_consumes/index.vue"),
      meta: {
        title: "消费明细"
      }
    }
  ]
} satisfies RouteConfigsTable;
