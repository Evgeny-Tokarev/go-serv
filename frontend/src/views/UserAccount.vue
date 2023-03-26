<template>
  <div style="height: 100%">
    <v-card
        class="mx-auto"
        max-width="40%"
    >
      <div class="userPhoto d-flex justify-center position-relative">
        <v-icon
            v-if="user.image !== store.currentUser?.photo"
            color="error"
            icon="mdi:mdi-file-undo-outline"
            class="undo position-absolute"
            @click="undoPhotoPick"
        ></v-icon>
        <label class="position-absolute d-flex justify-center align-content-start">
          <input
              ref="imgInput"
              type="file"
              @change="onImageChange"
          >
        </label>
        <v-img
            :src="user.image || dummyImage"
            contain
            class="userImg elevation-1"
        ></v-img>
      </div>


      <div class="userName d-flex justify-space-between align-center">
        <label class="d-flex justify-center align-content-start pa-5">
          <input
              ref="imgInput"
              v-model=" user.name"
              type="text"
          >
        </label>
        <v-icon
            v-if="user.name !== store.currentUser?.username"
            color="error"
            icon="mdi:mdi-file-undo-outline"
            class="undo"
            @click="undoNameChange"
        ></v-icon>
      </div>

      <v-card-subtitle>
        {{ $t("account.created") }}: {{ user.createdAt.toLocaleString(locale, {dateStyle: 'medium'}) }}
      </v-card-subtitle>

      <v-card-actions>
        <v-btn
            color="orange-lighten-2"
            variant="outlined"
            @click="saveChanges"
        >
          {{ $t("account.save") }}
        </v-btn>

        <v-btn
            color="orange-lighten-2"
            @click="show = !show"
        >
          {{ $t("account.cvs") }} {{ store.currentUser?.CVList.length }}
        </v-btn>

        <v-spacer></v-spacer>

        <v-btn
            :icon="show ? 'mdi:mdi-chevron-up' : 'mdi:mdi-chevron-down'"
            @click="show = !show"
        ></v-btn>
      </v-card-actions>

      <v-expand-transition>
        <div v-show="show">
          <v-divider></v-divider>

          <v-card
          >
            <v-list
                :items="[{name: 'item1'}, {name: 'item2'}]"
                item-title="name"
                item-value="name"
            ></v-list>
          </v-card>
        </div>
      </v-expand-transition>
    </v-card>
  </div>
</template>
<script setup lang="ts">
import {ref, nextTick} from "vue"
import dummyImage from "@/assets/img/person.svg"
import {useUserStore} from "@/stores"
import {useI18n} from "vue-i18n";

interface HTMLInputEvent extends Event {
  target: HTMLInputElement & EventTarget;
}

const {locale} = useI18n()
const store = useUserStore()
const show = ref(false)
const imgInput = ref<HTMLInputElement | null>(null)
const imageFile = ref<File | null>(null)
const user = ref({
  image: store.currentUser?.photo,
  name: store.currentUser?.username,
  createdAt: store.currentUser?.CreatedAt
})
const onImageChange = (e: HTMLInputEvent) => {
  const target = e.target as HTMLInputElement
  const files = target.files
  if (!files || !files[0])
    return
  user.value.image = URL.createObjectURL(files[0])
  imageFile.value = files[0]
}
const saveChanges = async() => {
  if (  user.value.image === store.currentUser?.photo && user.value.name === store.currentUser?.username && user.value.createdAt === store.currentUser?.CreatedAt) return
  await store.changeUser(user.value)
}
const undoNameChange = () => {
  user.value.name = store.currentUser?.username
}
const undoPhotoPick = async () => {
  user.value.image = store.currentUser?.photo
}
</script>
<style scoped lang="scss">
.userPhoto {
  height: 30vh;
  z-index: 1;
  margin-top: 1rem;

  .undo {
    z-index: 1;
    right: 0;
  }

  label {
    top: 0;
    bottom: 0;
    left: 10%;
    right: 10%;
    cursor: pointer;
    z-index: 1;

    input {
      display: none;
    }
  }


  .userImg {
    height: 100%;
    flex: 0 0 80%;
  }
}
.userName {
  input {
    outline: none;
  }
}
</style>
