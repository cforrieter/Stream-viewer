<template>
  <div class="home">
    <h1>Login page</h1>

    <button @click="login">Login with Google</button>
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: "login",
  components: {
  },
  methods: {
    login() {
       this.$getGapiClient()
        .then(gapi => {
          let GoogleAuth = gapi.auth2.getAuthInstance();

          console.log(this)
          if(!GoogleAuth.isSignedIn.get()) {
            GoogleAuth.signIn().then( () => {
              this.$router.push({ name: 'home'})
            });
          } else {
            this.$router.push({ name: 'home'})
          }
        })
    }
  }
};
</script>
