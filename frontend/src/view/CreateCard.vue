<template>
  <div class="container">

    <select class="form-select" v-model="tag" aria-label="Default select example">
      <option value="">Peek the tag</option>
      <option value="Course">Course</option>
      <option value="Book">Book</option>
      <option value="Video">Video</option>
      <option value="Pet">Pet</option>
    </select>

    <div class="input-group mb-3 mt-3">
      <span class="input-group-text" id="basic-addon1">Title</span>
      <input type="text" class="form-control" v-model="title"
             aria-describedby="basic-addon1">
    </div>

    <div class="mb-3">
      <label for="exampleFormControlTextarea1" class="form-label">Description</label>
      <textarea class="form-control" id="exampleFormControlTextarea1" rows="3" v-model="description"></textarea>
    </div>
    <div class="mb-3">
      <button type="button" class="btn btn-secondary" @click="onSubmit">Create</button>
    </div>
    <ul class="err">
      <li v-for="(e, idx) in errors" :key="idx">{{ e }}</li>
    </ul>

    <div v-if="success" class="success">
      {{ success }}
    </div>
  </div>
</template>

<script>
import Constants from "@/commom/constants"

export default {
  name: "CreateCard",

  data() {
    return {
      tag: "",
      title: "",
      description: "",
      errors: [],
      success: ""
    }
  },

  methods: {
    async onSubmit() {
      // validation
      this.errors = []
      if (!this.tag) this.errors.push("peek the tag")
      if (this.title.length > 250 || this.title.length === 0)
        this.errors.push("title must be have length between 0 and 250")
      if (this.description.length > 500) this.errors.push("title must be have length not great than 500")
      if (!this.errors) return
      // create card
      const body = {
        tag: this.tag,
        title: this.title,
        description: this.description
      }

      const resp = await fetch(`${Constants.BASE_URL}/card`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "origin": "localhost:8080",
        },
        body: JSON.stringify(body)
      })
      this.errors = []
      if (resp.status === 200) {
        this.tag = ""
        this.title = ""
        this.description = ""
        this.success = "the card was successfully created"
        setTimeout(() => this.success = "", 2000)
        return
      }
      this.errors.push("server responded by error")
    }
  }

}
</script>

<style scoped>
.container {
  margin-top: 10px;
}

.err {
  color: darkred;
}

.success {
  color: green;
}
</style>