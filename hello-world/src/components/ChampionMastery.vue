<template>
    <div v-if="encryptedSummonerID" class="main-container">

    </div>
</template>

<script>
import AWS from 'aws-sdk'
export default {
    name: 'SummonerProfile',
    data() {
        return {
            ChampionPoints: Number,
            ChampionID: "lolitspetey",
            ChampionLevel: Number,
        }
    },
    props: {
        encryptedSummonerID: {
            type: String
        }
    },
    computed: {
        updateEncryptedSummonerID() {
            return this.encryptedSummonerID
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
                masteryResults = JSON.parse(data.Payload);
                this.ChampionPoints = masteryResults.championPoints;
                this.ChampionID = masteryResults.championId;
                this.ChampionLevel = masteryResults.championLevel;
                console.log(masteryResults)
                }
            })
        }
    },
}
</script>

<style scoped>

</style>