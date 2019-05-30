<template>
  <div id="app">
    <SummonerProfile/>
  </div>
</template>

<script>
import SummonerProfile from './components/SummonerProfile.vue'

export default {
  name: 'app',
  components: {
    SummonerProfile
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
