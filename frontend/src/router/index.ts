import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";

const routes: RouteRecordRaw[] = [
  {
    path: "/login",
    name: "Login",
    component: () => import("../views/Login.vue"),
    meta: { requiresAuth: false },
  },
  {
    path: "/",
    name: "Layout",
    component: () => import("../views/Layout.vue"),
    meta: { requiresAuth: true },
    children: [
      {
        path: "",
        name: "Home",
        component: () => import("../views/Home.vue"),
      },
      {
        path: "products",
        name: "Products",
        component: () => import("../views/Products.vue"),
      },
      {
        path: "licenses",
        name: "Licenses",
        component: () => import("../views/Licenses.vue"),
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// 路由守卫
router.beforeEach(async (to, _, next) => {
  const token = localStorage.getItem("token");
  const requiresAuth = to.matched.some((record) => record.meta.requiresAuth);

  // 处理未授权情况
  const handleUnauthorized = () => {
    localStorage.removeItem("token");
    next({
      path: "/login",
      query: { redirect: to.fullPath },
    });
  };

  if (requiresAuth && !token) {
    handleUnauthorized();
  } else {
    next();
  }
});

export default router;
