<template>
  <div class="summonerProfileBg">
    <div v-if="isLoading" class="loadingIcon">
    </div>
    <div v-else>
      <SummonerProfile></SummonerProfile>
    </div>
  </div>

</template>

<script>
import AWS from 'aws-sdk'
import SummonerProfile

export default {
  name: 'HomePage',
  data() {
    return {
      summonerName: String,
      summonerLevel: Number,
      isLoading: false,
      userSummonerName: "lolitspetey",
      profileIconID: Number,
      encryptedSummonerID: String,
      championLevel: Number,
      championMasteryPoints: Number,
      championID: Number,
    }
  },
methods: {
  getChampionMasteryByID() {

    //load AWS credentials
    AWS.config.region = 'us-east-2';
    AWS.config.credentials = new AWS.CognitoIdentityCredentials({
        IdentityPoolId: 'us-east-2:20fd57cf-7bb6-4352-b5ce-69eca3907336'
      });

    //create AWS service object
    var lambda = new AWS.Lambda({region: 'us-east-2'});
    //create JSON object for parameters for invoking Lambda function
    var masteryParams = {
        FunctionName: 'getChampionMasteryByID',
        InvocationType: 'RequestResponse',
        LogType: 'None',
        Payload: '{"SummonerID?": ' + `"` + this.encryptedSummonerID + '"}'
      };
      //create variable to hold data returned by Lambda function
    var masteryResults;

    lambda.invoke(masteryParams, (error, data) => {
        if (error) {
          this.isLoading = false;
          prompt(error);
        } else {
          masteryResults = JSON.parse(data.Payload)
          console.log(masteryResults);
        }
      })
    }
  },
computed: {
  
  },
  created() {
  this.getNASummonerData();
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
