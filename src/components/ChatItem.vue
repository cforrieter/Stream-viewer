<template>
  <div class="flexItem">
    <span>{{author}}: {{message}}</span>
  </div>
</template>

<script>
export default {
  name: "ChatItem",
  props: {
    author: String,
    message: String
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.flexItem {
  width: 100%;
  text-align: left;
  margin: 1px;
  padding: 3px;
  background-color: aqua;
  
}
</style>
