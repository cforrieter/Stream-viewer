<template>
  <div class="flexItem">
    <p>{{name}}</p>
    <img :src="thumbnail" />
    <p><router-link :to="'/stream/' + videoId">{{title}}</router-link></p>
  </div>
</template>

<script>
export default {
  name: "StreamPreview",
  props: {
    name: String,
    title: String,
    thumbnail: String,
    videoId: String
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.flexItem {
  width: 400px;
  display: inline-block;
  margin: 15px;
  background-color: aqua
}
</style>
