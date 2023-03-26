import vue from "@vitejs/plugin-vue";
import { join, resolve } from "path";
import { defineConfig } from "vite";
import viteCompression from "vite-plugin-compression";
import VueTypeImports from "@rah-emil/vite-plugin-vue-type-imports";
import vuetify from "vite-plugin-vuetify"
import vueI18n from '@intlify/vite-plugin-vue-i18n'

export default defineConfig({
  plugins: [vue({}), viteCompression(), VueTypeImports(), vuetify(), vueI18n({
    include: resolve(__dirname, 'src/locales/**'),
    runtimeOnly: false
  })],
  define: {
    __VUE_I18N_FULL_INSTALL__: true,
    __VUE_I18N_LEGACY_API__: false,
    __INTLIFY_PROD_DEVTOOLS__: false,
  },
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
