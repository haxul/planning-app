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
        <Card v-for="(elem, idx) in backlogList" :key="elem.id + idx" :title="elem.title"
              :description="elem.description"
              :tag="elem.tag"
        />
      </div>
      <div class="col">
        <Card v-for="(elem, idx) in inProgressList" :key="elem.id + idx" :title="elem.title"
              :description="elem.description"
              :tag="elem.tag"
        />
      </div>
      <div class="col">
        <Card v-for="(elem, idx) in doneList" :key="elem.id + idx" :title="elem.title"
              :description="elem.description"
              :tag="elem.tag"
        />
      </div>
      <div class="col">
        <Card v-for="(elem, idx) in rejectedList" :key="elem.id + idx" :title="elem.title"
              :description="elem.description"
              :tag="elem.tag"
        />
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
    ...mapGetters(["getCards"]),
    backlogList() {
      return this.getCards.filter(e => e.cur_state === "BACKLOG")
    },
    inProgressList() {
      return this.getCards.filter(e => e.cur_state === "IN_PROGRESS")
    },
    doneList() {
      return this.getCards.filter(e => e.cur_state === "DONE")
    },
    rejectedList() {
      return this.getCards.filter(e => e.cur_state === "REJECTED")
    }
  },

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