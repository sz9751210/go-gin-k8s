import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/home",
    name: "home",
    icon: "odometer",
    meta: { title: "Home", requireAuth: false },
    component: () => import("@/views/home/Home.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});
export default router;
