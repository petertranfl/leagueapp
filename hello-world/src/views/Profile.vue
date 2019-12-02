<template>
  <div class="summonerProfileBg">
    <div v-show="isLoading" class="loadingIcon">
    </div>
    <div v-show="!isLoading">
      <SummonerProfile></SummonerProfile>
      <ChampionMastery></ChampionMastery>
      <MatchHistory @clicked="finishLoading"></MatchHistory>
    </div>
  </div>

</template>

<script>
import SummonerProfile from '../components/SummonerProfile'
import ChampionMastery from '../components/ChampionMastery'
import MatchHistory from '../components/MatchHistory'


export default {
  name: 'Profile',
  components: {
    SummonerProfile,
    ChampionMastery,
    MatchHistory
  },
  data() {
    return {
      isLoading: true,
    }
  },
  methods: {
    finishLoading(isLoaded) {
      console.log(isLoaded)
      this.isLoading = isLoaded
    }
  }
}
</script>


<style>
html, body {
  margin: 0;
  padding: 0;
}
.loadingIcon {
  border: 16px solid #f3f3f3; /* Light grey */
  border-top: 16px solid #3498db; /* Blue */
  border-radius: 50%;
  width: 60px;
  height: 60px;
  animation: spin 2s linear infinite;
  align-content: center;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.summonerProfileBg{
  background: url("https://lolstatic-a.akamaihd.net/lolkit/1.1.6/resources/images/bg-default.jpg") no-repeat center center fixed;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;
}
</style>
