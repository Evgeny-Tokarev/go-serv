import {createRouter, createWebHistory} from "vue-router";
import HomeView from "@/views/HomeView.vue";
import LoginPage from "@/views/LoginPage.vue";
import {useUserStore} from "@/stores"
import UserAccount from "@/views/UserAccount.vue"

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: "/",
            name: "home",
            component: HomeView,
        },
        {
            path: "/login",
            name: "Login",
            component: LoginPage
        },
        {
            path: "/account",
            name: "Account",
            component: UserAccount
        }
    ],
});
router.beforeEach(async (to, from) => {
    const store = useUserStore()
    if (
        !store.currentUser &&
        to.name !== 'Login'
    ) {
        return {name: 'Login'}
    }
})

export default router;
