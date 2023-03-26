<template>
  <v-app v-if="store.currentUser">
    <v-card
        class="app-wrapper"
    >
      <v-layout>
        <v-navigation-drawer
            v-model="drawer"
            temporary
        >
          <v-list-item
              prepend-avatar="store.currentUser.photo"
              :title="store.currentUser.username"
          ></v-list-item>
          <v-divider></v-divider>
          <v-list density="compact" nav
                  class="menu-list d-flex flex-column"
          >
            <v-list-item prepend-icon="mdi:mdi-account-box" :title="$t('menu.account')" value="account" to="/account">
            </v-list-item>
            <v-list-item prepend-icon="mdi:mdi-file" :title="$t('menu.myResume')" value="about" to="/someRoute2"></v-list-item>
            <v-btn
                color="primary"
                @click.stop="drawer = false"
                icon="mdi:mdi-close"
                style="margin-top: auto; margin-right: auto; margin-left: auto"
            />
          </v-list>
        </v-navigation-drawer>
        <v-main style="width: 100%; height: 100%">
          <div class="pa-4" style="height: 100vh">
            <v-btn
                color="primary"
                icon="mdi:mdi-menu"
                @click.stop="drawer = !drawer"
                class="menu-button"
            />
            <h1 class="app_title text-h4 mb-4">
              {{ $t('mainPage.title') }}
            </h1>
            <AnimatedRouterView/>
          </div>
        </v-main>
      </v-layout>
    </v-card>
  </v-app>
</template>

<script setup lang="ts">
import AnimatedRouterView from '@/components/reusable/AnimatedRouterView.vue'
import NavigationBar from "@/components/reusable/NavigationBar.vue"
import {ref, onMounted} from "vue"
import {useUserStore} from "@/stores"
import {useRouter} from "vue-router";

const router = useRouter()
const drawer = ref(false)
const store = useUserStore()
onMounted(async () => {
  const res = await store.setCurrentUser()
  if (res) await router.push("/")
})
</script>
<style lang="scss" scoped>
.app-wrapper {
  height: 100vh;
  display: flex;

  .menu-button {
    position: absolute;
    top: 1rem;
    left: 1rem;
  }

  :deep(.v-navigation-drawer__content) {
    display: flex;
    flex-direction: column;

    .menu-list {
      flex-grow: 1;
    }
  }

  .app_title {
    text-transform: uppercase !important;
    text-align: center;
  }
}


</style>
