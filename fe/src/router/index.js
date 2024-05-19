import { createRouter, createWebHistory } from "vue-router";
// 導入進度條
import NProgress from "nprogress";
import "nprogress/nprogress.css";

const routes = [
  {
    path: "/layout",
    icon: "odometer",
    component: () => import("@/layout/Layout.vue"),
    children: [
      {
        path: "/home",
        name: "home",
        icon: "odometer",
        meta: { title: "Home", requireAuth: false },
        component: () => import("@/views/home/Home.vue"),
      },
    ],
  },
  {
    path: "/404",
    name: "404",
    meta: { title: "Page Not Found", requireAuth: false },
    component: () => import("@/views/common/404.vue"),
  },
  {
    path: "/:pathMatch(.*)",
    redirect: "/404",
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// 進度條配置
NProgress.inc(0.2); // 設置進度條遞增
NProgress.configure({ easing: "ease", speed: 500, showSpinner: false }); // 設置進度條顯示方式

// 路由守衛
router.beforeEach((to, from, next) => {
  // 進度條開始
  NProgress.start();
  // 設置頭部
  if (to.meta.title) {
    document.title = to.meta.title;
  } else {
    document.title = "Kubernetes";
  }
  // 放行
  next();
});

router.afterEach(() => {
  // 進度條結束
  NProgress.done();
});

export default router;
