<template>
  <div class="hello">
    <h1>{{summonerName}}</h1>
    
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'SummonerProfile',
  data() {
    return {
      summonerName: String,
      loading: false,
      userSummonerName: 'lolitspetey',
      apiKey: 'RGAPI-564bf915-7f35-418b-8bd5-5db5bcfaf5da'
    }
  },

methods: {
  getSummonerData() {
    this.loading = true;
    axios.get('https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-name/' + this.userSummonerName + '?api_key=' + this.apiKey)
      .then((response) => {
        this.loading = false;
        console.log(response);
        this.summonerName = response.name
      })
    }
  },
  created() {
  this.getSummonerData();
},

}
</script>


<style scoped>
@font-face {
  font-family: 'Friz Quadrata Regular'; /* League of Legends Font (almost) */
  src: url('../assets/Friz Quadrata Regular.ttf') format('truetype'); /* Safari, Android, iOS */
  src: url('../assets/Friz Quadrata Regular.woff') format('woff') /* Modern Browsers */
}

h1 {
  font-family: 'Friz Quadrata Regular', sans-serif;
}

</style>
