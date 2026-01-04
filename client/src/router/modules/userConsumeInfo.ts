const Layout = () => import("@/layout/index.vue");

export default {
  path: "/user_consumes",
  component: Layout,
  redirect: "/user_consumes/index",
  meta: {
    icon: "ri/information-line",
    title: "数据统计",
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
    },
    {
      path: "dailyReport",
      name: "DailyReport",
      component: () => import("@/views/user_consumes/daily_report.vue"),
      meta: {
        title: "日报"
      }
    }
  ]
} satisfies RouteConfigsTable;
