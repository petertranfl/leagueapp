<template>
  <div v-if="isLoading" class="loadingIcon">

  </div>
  <div v-else>
    
    <h1>{{summonerName}}</h1>
    <p>{{summonerLevel}}</p>
    <h2><img :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/profile-icons/' + profileIconID + '.jpg'"/></h2>
    
  </div>
</template>

<script>
import AWS from 'aws-sdk'

export default {
  name: 'SummonerProfile',
  data() {
    return {
      summonerName: String,
      summonerLevel: Number,
      isLoading: false,
      userSummonerName: "lolitspetey",
      profileIconID: Number
    }
  },
methods: {
  getNASummonerData() {
    this.isLoading = true;

    //load AWS credentials
    AWS.config.region = 'us-east-2';
    AWS.config.credentials = new AWS.CognitoIdentityCredentials({
        IdentityPoolId: 'us-east-2:20fd57cf-7bb6-4352-b5ce-69eca3907336'
      });

    //create AWS service object
    var lambda = new AWS.Lambda({region: 'us-east-2'});
    //create JSON object for parameters for invoking Lambda function
    var reqParams = {
        FunctionName: 'getNASummonerDataByName',
        InvocationType: 'RequestResponse',
        LogType: 'None',
        Payload: '{"SummonerName?": ' + `"` + this.userSummonerName + '"}'
      };
      //create variable to hold data returned by Lambda function
    var reqResults;

      //Calls Lambda function 'GetSummonerDataByName' with given reqParams
    lambda.invoke(reqParams, (error, data) => {
        if (error) {
          this.isLoading = false;
          prompt(error);
        } else {
          reqResults = JSON.parse(data.Payload)
          this.summonerName = reqResults.name
          this.summonerLevel = reqResults.summonerLevel
          this.profileIconID = reqResults.profileIconId
          this.isLoading = false;
        }
      });
    }
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
  width: 120px;
  height: 120px;
  animation: spin 2s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
</style>
