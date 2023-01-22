import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import AlbumControls from "@/views/AlbumControls.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/controls",
      name: "controls",
      component: AlbumControls,
    },
  ],
});

export default router;
