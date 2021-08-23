<template>
  <p class="align-middle txt">{{ label }} <span id="cycle" :class="{
    course: isCourse,
    pet: isPet,
    book: isBook,
    video: isVideo
  }"></span></p>
</template>

<script>
export default {
  name: "Tag",
  props: {
    label: {
      required: true,
      type: String
    }
  },

  computed: {
    isCourse() {
      return this.label === "Course"
    },
    isBook() {
      return this.label === "Book"
    },
    isPet() {
      return this.label === "Pet"
    },
    isVideo() {
      return this.label === "Video"
    }
  }
}

</script>

<style scoped>
#cycle {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  display: inline-block;
}

.course {
  background-color: darkorange;
}

.book {
  background-color: purple;
}

.pet {
  background-color: green;
}

.video {
  background-color: blue;
}

.txt {
  color: rgba(0, 0, 0, 0.5);
  font-size: 10pt;
  font-family: Arial, sans-serif;
}
</style>