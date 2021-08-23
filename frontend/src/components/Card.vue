<template>
  <div class="card" style="width: 14rem;">
    <div class="card-body">
      <Tag :label="tag"/>
      <h5 class="card-title">{{ title }}</h5>
      <p class="card-text">{{ description }}</p>
    </div>
  </div>
</template>

<script>
import Tag from "@/components/Tag";

export default {
  name: "Card",
  components: {Tag},
  props: {
    title: {
      type: String,
      required: true
    },
    description: {
      type: String,
      required: false
    },
    tag: {
      type: String,
      required: true
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
</style>



















