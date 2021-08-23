<template>
  <div class="container">
    <div class="row align-items-center justify-content-center">
      <div class="col">
        <h5>Backlog</h5>
      </div>
      <div class="col">
        <h5>In progress</h5>
      </div>
      <div class="col">
        <h5>Done</h5>
      </div>
      <div class="col">
        <h5>Rejected</h5>
      </div>
    </div>
    <div class="row align-items-start">
      <div class="col">
        <Card v-for="(elem, idx) in getBacklogList" :key="elem.id + idx" :card="elem"/>

      </div>
      <div class="col">
        <Card v-for="(elem, idx) in getInProgressList" :key="elem.id + idx" :card="elem"/>
      </div>
      <div class="col">
        <Card v-for="(elem, idx) in getDoneList" :key="elem.id + idx" :card="elem"/>
      </div>
      <div class="col">
        <Card v-for="(elem, idx) in getRejectedList" :key="elem.id + idx" :card="elem"/>
      </div>
    </div>
  </div>
</template>

<script>

import Card from "@/components/Card";
import {mapGetters} from "vuex"

export default {
  name: "Board",
  components: {Card},
  computed: {
    ...mapGetters(["getDoneList", "getInProgressList", "getBacklogList", "getRejectedList"])
  }
  ,
  mounted() {
    this.$store.dispatch("fetchCards")
  }
}
</script>

<style scoped>
.col {
  margin-top: 25px;
}
</style>