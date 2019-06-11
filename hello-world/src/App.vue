<template>
  <div id="app">
    <Profile/>
  </div>
</template>

<script>
import Profile from './views/Profile.vue'

export default {
  name: 'app',
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
  color: #2c3e50;
}
</style>
