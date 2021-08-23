<template>
  <div id="app">
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
      <div class="container-fluid">
        <span class="navbar-brand"><img src="@/assets/tree.png" alt="not found"></span>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarText"
                aria-controls="navbarText" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarText">
          <ul class="navbar-nav me-auto mb-2 mb-lg-0">
            <li class="nav-item" v-for="page in pages" :key=page.name
                @click="handleLi(page.name)">
              <router-link class="nav-link" :class="{active: page.name === curPage}" aria-current="page" :to=page.link>
                {{ page.name }}
              </router-link>
            </li>
          </ul>
        </div>
      </div>
    </nav>
    <router-view/>
  </div>
</template>

<script>

export default {
  name: 'App',
  components: {},
  data() {
    return {
      curPage: "Board",
      pages: [
        {name: "Board", link: "/"},
        {name: "Create Card", link: "/create-card"}
      ]
    }
  },
  computed: {
    isCurPage(page) {
      return this.curPage === page
    }
  },
  methods: {
    handleLi(page) {
      this.curPage = page
    }
  }
}
</script>

<style>

</style>
















