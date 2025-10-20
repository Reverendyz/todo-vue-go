<script setup lang="ts">
import { ref } from 'vue'
import { RouterLink, RouterView } from 'vue-router'

const drawer = ref(true)
</script>

<template>
  <v-responsive class="border rounded">
    <v-app>
      <v-navigation-drawer v-model="drawer">
        <v-list>
          <v-list-item prepend-icon="mdi-home"><router-link to="/">Home</router-link></v-list-item>
          <v-list-item prepend-icon="mdi-help"
            ><router-link to="/about">About</router-link></v-list-item
          >
        </v-list>
      </v-navigation-drawer>

      <v-app-bar title="Tooldo" class="d-flex mt-3 p-5 border rounded" />

      <v-main>
        <v-btn class="mt-5" color="success" @click.stop="drawer = !drawer">Toggle drawer</v-btn>
        <RouterView />
      </v-main>
    </v-app>
  </v-responsive>
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
