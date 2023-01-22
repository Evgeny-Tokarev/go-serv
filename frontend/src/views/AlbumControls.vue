<template>
  <div class="album-form">
    <v-form
      class="album-form__form"
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

      <v-text-field
        v-model="artist"
        :rules="nameRules"
        label="Artist"
        required
      ></v-text-field>

      <v-text-field
        v-model="price"
        :counter="10"
        :rules="priceRools"
        label="Price"
        required
      ></v-text-field>
      <p>valid{{ valid }}</p>
      <div class="album-form__controls">
        <v-btn
          color="success"
          class="album-form__control-button"
          @click="validate"
        >
          Validate</v-btn
        >

        <v-btn color="error" class="album-form__control-button" @click="reset">
          Reset Form</v-btn
        >

        <v-btn color="success" type="submit" class="album-form__control-button">
          Send Album</v-btn
        >
      </div>
    </v-form>
  </div>
</template>
<script lang="ts" setup>
import { ref } from "vue";
import { addNewAlbum } from "@/api";

const form = ref<any>(null);
const valid = ref(true);
const title = ref("");
const artist = ref("");
const price = ref("");
const nameRules = [
  (v: string) => !!v || "Name is required",
  (v: string) => (v && v.length > 2) || "Name is too short",
];
const priceRools = [
  (v: string) => /^\d+$/.test(v) || "Price must be an integer number",
];

const sendAlbum = async () => {
  console.log(valid.value);
  if (title.value && artist.value && price.value && valid.value) {
    const res = await addNewAlbum({
      title: title.value,
      price: Number(price.value),
      artist: artist.value,
    });
    console.log(res);
    if (res.status >= 200 && res.status < 300) {
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
.album-form {
  min-height: 100vh;
  width: 60%;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  align-items: center;
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
}
</style>
