<template>
  <div id="app">
    <Profile/>
  </div>
</template>

<script>
import Profile from './views/Profile.vue'
import store from './store/index.js'

export default {
  name: 'app',
  store,
  components: {
    Profile
  }
}
</script>

<style>
@font-face {
  font-family: 'Friz Quadrata Regular'; /* League of Legends Font (almost) */
  src: url('./assets/Friz Quadrata Regular.ttf') format('truetype'); /* Safari, Android, iOS */
  src: url('./assets/Friz Quadrata Regular.woff') format('woff') /* Modern Browsers */
}

#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #000000;

}

body {
  margin: 0px;
}
</style>
