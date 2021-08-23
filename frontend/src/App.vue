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
            <li class="nav-item">
              <router-link class="nav-link active" aria-current="page" to="/">Board</router-link>
            </li>
            <!--            Feature is not ready-->
            <!--            <li class="nav-item">-->
            <!--              <router-link class="nav-link " to="/account">Accounting</router-link>-->
            <!--            </li>-->
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
  components: {}
}
</script>

<style>

</style>
















