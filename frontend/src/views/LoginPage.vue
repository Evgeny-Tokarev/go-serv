<script setup lang="ts">
import {ref} from "vue";
// import {useUserStore} from '@/stores/userStore'
import {useRouter} from 'vue-router'
import {useUserStore} from "@/stores/userStore"
// import mainIcon from '@/assets/img/main-icon.svg'
import {useI18n} from "vue-i18n"

const {t} = useI18n()
const router = useRouter()
const store = useUserStore()
const name = ref('')
const password = ref('')
const message = ref('')
const login = async () => {
  if (!name.value || !password.value) return
  const res = await store.authenticateUser(name.value.trim(), password.value.trim())
  if (res) {
    await router.push('/')
  } else {
    message.value = t('login.message.fail')
  }
}
</script>

<template>
  <div class="modal">
    <div class="modal__content">
      <p class="modal__text">
        {{$t('login.title')}}
      </p>
      <div>
        <v-text-field
            v-model="name"
            :label="$t('login.name')"
            bg-color="white"
            class="text-field"
            @focus="message=''"
        />
        <v-text-field
            v-model="password"
            :label="$t('login.password')"
            bg-color="white"
            class="text-field"
            @focus="message=''"
        />
      </div>
      <div class="modal__footer">
        <v-btn
            class="send-button"
            color="white"
            @click="login"
        >
          {{$t('login.enter')}}
        </v-btn>
        <p class="modal__message modal__text">
          {{ message }}
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.modal {
  position: fixed;
  top: 0;
  left: 0;
  z-index: 1;
  width: 100vw;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background: var(--color-accent);

  &__content {
    display: flex;
    width: 50%;
    height: 65%;
    min-width: 380px;
    min-height: 380px;
    max-height: 450px;
    max-width: 600px;
    flex-direction: column;
    align-items: stretch;
    justify-content: space-between;
    background: var(--color-accent);
    border-radius: 5px;
    padding: 2rem 4rem;
    text-align: center;
  }

  &__logo {
    display: flex;
    justify-content: center;
    background: #16A5F0;
    border-radius: 5px;
    padding: 10px;
  }

  &__logo-text {
    font-size: 30px;
    color: white;
    font-weight: 700;
  }

  &__text {
    color: white;
    font-weight: 700;
    font-size: 2rem;
  }
  &__footer {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
  &__message {
    color: red;
    font-weight: 700;
    font-size: 0.8rem;
  }
  @media screen and (min-width: 600px) {
    background: white;
    &__content {
      box-shadow: 0 0 1rem black;
    }
  }
}

.text-field {
  :deep(label) {
    color: var(--color-accent);
    font-size: 1.3rem;
  }

  :deep(input) {
    font-size: 1.5rem;
  }
}

.send-button {
  font-weight: 700;
  font-size: 1.3rem;

  :deep(.v-btn__content) {
    color: var(--color-accent);
  }
}
</style>
