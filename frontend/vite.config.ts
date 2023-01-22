import vue from "@vitejs/plugin-vue";
import { join } from "path";
import { defineConfig } from "vite";
import viteCompression from "vite-plugin-compression";
import VueTypeImports from "@rah-emil/vite-plugin-vue-type-imports";
import vuetify from "vite-plugin-vuetify";

export default defineConfig({
  plugins: [vue({}), viteCompression(), VueTypeImports(), vuetify()],
  resolve: {
    alias: [
      {
        // для алиасов в секции style
        find: /~(.+)/,
        replacement: join(__dirname, "src/$1"),
      },
      {
        find: /@\//,
        replacement: join(__dirname, "src") + "/",
      },
      {
        find: "vue",
        replacement: "vue/dist/vue.esm-bundler.js",
      },
    ],
  },
});
