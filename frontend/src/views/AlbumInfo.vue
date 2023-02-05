<template>
  <div class="album-info">
    <h2 class="album-info__title">Album Info</h2>
    <v-form
      class="album-info__form"
      ref="form"
      v-model="valid"
      @submit.prevent="sendAlbum"
      lazy-validation
    >
      <v-text-field
        v-model="title"
        :counter="10"
        :rules="nameRules"
        label="Title"
        required
      ></v-text-field>
      <div class="album-info__controls">
        <v-btn
          color="success"
          class="album-info__control-button"
          @click="validate"
        >
          Validate</v-btn
        >

        <v-btn color="error" class="album-info__control-button" @click="reset">
          Reset Form</v-btn
        >

        <v-btn color="success" type="submit" class="album-info__control-button">
          Get Info</v-btn
        >
      </div>
    </v-form>
    <div class="album-info__content">
      <h3 class="album-info__subtitle">
        {{ album.title }}
      </h3>
      <p class="album-info__info">{{ album.artist }}</p>
      <p class="album-info__info">{{ album.price }}</p>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { ref, reactive } from "vue";
import { getAlbumInfoByTitle } from "@/api";

const form = ref<any>(null);
const valid = ref(true);
const title = ref("");
const nameRules = [
  (v: string) => !!v || "Name is required",
  (v: string) => (v && v.length > 2) || "Name is too short",
];
const album = reactive({
  title: "",
  artist: "",
  price: "",
});
const sendAlbum = async () => {
  console.log(valid.value);
  if (title.value && valid.value) {
    const res = await getAlbumInfoByTitle(title.value);
    console.log(res);
    if (res.status >= 200 && res.status < 300) {
      [album.title, album.artist, album.price] = [
        res.data.title,
        res.data.artist,
        res.data.price,
      ];
      form.value?.reset();
    }
  }
};
const validate = async () => {
  if (form.value) {
    const { valid } = await form.value.validate();

    if (valid) alert("Form is valid");
  }
};
const reset = () => {
  form.value?.reset();
};
</script>

<style lang="scss" scoped>
.album-info {
  padding: 2rem;
  min-height: 100vh;
  width: 80%;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  align-items: center;
  &__title {
    margin: 2rem auto;
  }
  &__form {
    width: 100%;
    max-width: 500px;
  }
  &__controls {
    width: 100%;
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    align-items: center;
    gap: 5%;
    margin-top: 30px;
  }
  &__control-button {
    flex: 0 0 auto;
    margin-bottom: 20px;
  }
  &__content {
    margin-top: 2rem;
    text-align: center;
  }
}
</style>
