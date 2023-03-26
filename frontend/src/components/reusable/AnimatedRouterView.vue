<template>
  <router-view
      :key="$route.path + ( name + ($route.query[name]))"
      v-slot="{ Component }"
      :name="name"
  >
    <transition
        appear
        :name="name === undefined ? 'scale': 'fade'"
        mode="out-in"
    >
      <component :is="Component" />
    </transition>
  </router-view>
</template>

<script lang="ts">
export default {
  name: "AnimatedRouterView",
  props: {
    name: String
  }
}
</script>

<style scoped lang="scss">
.fade-enter-active {
  transition: opacity 0.2s ease;
}

.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.scale-enter-active,
.scale-leave-active {
  transition: all 0.5s ease;
}

.scale-enter-from,
.scale-leave-to {
  opacity: 0;
  transform: scale(0.95);
}
</style>
