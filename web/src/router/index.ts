import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Index from '@/views/home/index.vue';


const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: '/ssh/node',
    component: Index,
    redirect: '/ssh/node',
    children: [
      // 首页看板
      {
        path: "/ssh/node",
        name: "/ssh/node",
        component: () => import("@/views/node.vue"),
      },
      {
        path: "/ssh/dial",
        name: "/ssh/dial",
        component: () => import("@/views/dial.vue"),
      },
    ]
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
