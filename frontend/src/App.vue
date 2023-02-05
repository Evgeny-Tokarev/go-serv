<template>
  <header>
    <div class="wrapper">
      <v-card>
        <v-layout>
          <v-navigation-drawer permanent>
            <v-list>
              <v-list-item title="Add New Album" value="AddAlbum" to="/albums/add"></v-list-item>

              <v-list-item
                title="Get Album Info"
                value="AlbumsInfo"
                to="/albums/info"
              ></v-list-item>

              <v-list-item title="Update Album" value="UpdateAlbum" to="/albums/update"></v-list-item>
            </v-list>
          </v-navigation-drawer>

          <v-main class="main-content">
            <RouterView />
          </v-main>
        </v-layout>
      </v-card>
    </div>
  </header>

</template>
<script lang="ts" setup>
const onUpdate = (e: any) => {

console.log(e)
}
</script>
<style lang="scss" scoped>
</style>
