<template>
    <div class="main-container">
        <img :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/profile-icons/' + profileIconID + '.jpg'"/>
        <h2>{{summonerName}}</h2>
        <p>Lvl: {{summonerLevel}}</p>
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
            userSummonerName: "lolitspetey",
            profileIconID: Number,
            encrypedSummonerID: String,
        }
    },
    methods: {
        getNASummonerData() {

            //load AWS credentials
            AWS.config.region = 'us-east-2';
            AWS.config.credentials = new AWS.CognitoIdentityCredentials({
                IdentityPoolId: 'us-east-2:20fd57cf-7bb6-4352-b5ce-69eca3907336'
            });

            //create AWS service object
            var lambda = new AWS.Lambda({region: 'us-east-2'});
            //create JSON object for parameters for invoking Lambda function
            var summonerParams = {
                FunctionName: 'getNASummonerDataByName',
                InvocationType: 'RequestResponse',
                LogType: 'None',
                Payload: '{"SummonerName?": ' + `"` + this.userSummonerName + '"}'
            };
            //create variable to hold data returned by Lambda function
            var summonerResults;

            //Calls Lambda function 'GetSummonerDataByName' with given reqParams
            lambda.invoke(summonerParams, (error, data) => {
                if (error) {
                prompt(error);
                } else {
                summonerResults = JSON.parse(data.Payload)
                this.summonerName = summonerResults.name;
                this.summonerLevel = summonerResults.summonerLevel;
                this.profileIconID = summonerResults.profileIconId;
                this.encryptedSummonerID = summonerResults.id;
                this.$store.dispatch('commitSummonerID', this.encryptedSummonerID)
                }
            });
        }
    },
    created() {
        this.getNASummonerData()
    }
}
</script>

<style scoped>
    .main-container{
        display: flex;
        height: 300px;
        flex-direction: column;
        align-items: center;
    }
</style>