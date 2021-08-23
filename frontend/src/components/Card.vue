<template>
  <div class="card" style="width: 14rem;">
    <div class="card-body">
      <Tag :label="card.tag"/>
      <h5 class="card-title">{{ card.title }}</h5>
      <p class="card-text">{{ card.description }}</p>
      <div class="move" @click="handleMove()"></div>
      <div class="reject"></div>
    </div>
  </div>
</template>

<script>
import Tag from "@/components/Tag";

export default {
  name: "Card",
  components: {Tag},
  props: {
    card: {
      type: Object,
      required: true
    },
  },
  methods: {
    handleMove() {
      this.$store.commit("updateCardById", {id: this.card.id, newState: "IN_PROGRESS"})
    }
  }
}
</script>

<style scoped>
.card {
  border: black 1px solid;
  border-radius: 3%;
  margin-top: 5px;
}

.reject {
  background-image: url("../assets/forbidden.png");
  background-size: 25px;
  cursor: pointer;
  padding: 2px;
  border-radius: 20%;
  top: 4px;
  right: 10px;
  width: 25px;
  height: 25px;
  position: absolute;
}

.move {
  background-image: url("../assets/move2.png");
  background-size: 23px;
  background-repeat: no-repeat;
  cursor: pointer;
  padding: 2px;
  border-radius: 20%;
  top: 30px;
  right: 10px;
  width: 25px;
  height: 25px;
  position: absolute;
}

</style>



















