<template>
  <div class="hello">
    <h1>{{summonerName}}</h1>
    <p>{{summonerLevel}}</p>
    
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
      loading: false,
      userSummonerName: "EarthRune"
    }
  },

methods: {
  getNASummonerData() {
    this.loading = true;

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
          this.loading = false;
          prompt(error);
        } else {
          this.loading = false;
          reqResults = JSON.parse(data.Payload)
          this.summonerName = reqResults.name
          this.summonerLevel = reqResults.summonerLevel
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

</style>
