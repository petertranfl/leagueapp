<template>
  <div class="summonerProfileBg">
    <div v-if="isLoading" class="loadingIcon">
    </div>
    <div v-else>
      <SummonerProfile v-on:getSummonerID="getSummonerID"></SummonerProfile>
      <ChampionMastery :encrypedSummonerID="encryptedSummonerID" ></ChampionMastery>
    </div>
  </div>

</template>

<script>
import SummonerProfile from '../components/SummonerProfile'
import ChampionMastery from '../components/ChampionMastery'

export default {
  name: 'Profile',
  components: {
    SummonerProfile,
    ChampionMastery
  },
  data() {
    return {
      isLoading: false,
      encryptedSummonerID: String,
    }
  },
  methods: {
    getSummonerID(summonerID) {
      this.encryptedSummonerID = summonerID
      console.log(summonerID)
    }
  }
}
</script>


<style>
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
  height: 100vh;
  width: 100vw;
  background-image: url("https://lolstatic-a.akamaihd.net/lolkit/1.1.6/resources/images/bg-default.jpg");
  background-size: 100% 100%;
}
</style>
