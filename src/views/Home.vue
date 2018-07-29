<template>
  <div class="home">
    <img src="../assets/logo.png">
    <StreamList/>
  </div>
</template>

<script>
// @ is an alias to /src
import StreamList from "@/components/StreamList.vue";

export default {
  name: "home",
  components: {
    StreamList
  }
};
</script>
