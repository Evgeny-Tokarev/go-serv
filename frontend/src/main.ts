import "vuetify/styles"
import "@/assets/styles/main.scss"
import "@mdi/font/css/materialdesignicons.css"
import "@fortawesome/fontawesome-free/css/all.css"

import {createApp} from "vue"
import {createPinia} from "pinia"
import {createVuetify} from "vuetify"
import * as components from "vuetify/components"
import * as directives from "vuetify/directives"
import {aliases, fa} from 'vuetify/iconsets/fa'
import {mdi} from 'vuetify/iconsets/mdi'
import {createI18n} from "vue-i18n"
import en from '@/locales/en.json'
import ru from '@/locales/ru.json'

import App from "./App.vue"
import router from "./router"

export const i18n = createI18n({
    legacy: false,
    locale: 'ru',
    fallbackLocale: 'ru',
    globalInjection: true,
    messages: {
        ru: ru,
        en: en
    },
})
const app = createApp(App);
const vuetify = createVuetify({
    components,
    directives,
    icons: {
        defaultSet: 'fa',
        aliases,
        sets: {
            fa,
            mdi,
        }
    }
})

app.use(createPinia())
app.use(vuetify)
app.use(router)
app.use(i18n)
app.mount("#app")
