import { createRouter, createWebHistory } from "vue-router";
import UpdateAlbum from "@/views/UpdateAlbum.vue";
import AddAlbum from "@/views/AddAlbum.vue";
import AlbumInfo from "@/views/AlbumInfo.vue";
import HomeView from "@/views/HomeView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/albums",
      children: [
        {
          path: "add",
          name: "addAlbum",
          component: AddAlbum,
        },
        {
          path: "info",
          name: "albumInfo",
          component: AlbumInfo,
        },
        {
          path: "update",
          name: "updateAlbum",
          component: UpdateAlbum,
        },
      ],
    },
  ],
});

export default router;
