<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>

export default {
  name: 'app'
}
</script>

<style>
@font-face {
  font-family: 'Friz Quadrata Regular'; /* League of Legends Font (almost) */
  src: url('./assets/Friz Quadrata Regular.ttf') format('truetype'); /* Safari, Android, iOS */
  src: url('./assets/Friz Quadrata Regular.woff') format('woff') /* Modern Browsers */
}

#app {
  background: url("https://lolstatic-a.akamaihd.net/lolkit/1.1.6/resources/images/bg-default.jpg") no-repeat center center fixed;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #000000;
}

</style>
