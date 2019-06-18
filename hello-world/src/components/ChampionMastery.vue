<template>
    <div v-if="masteryDataLoaded" class="main-container">

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
            masteryDataLoaded: false,
        }
    },
    // computed: {
    //     encryptedSummonerID() {
    //         return store.state.encryptedSummonerID
    //     }
    // },
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
            Payload: '{"SummonerID?": ' + `"` + this.summonerID + '"}'
            };
        //create variable to hold data returned by Lambda function
        var masteryResults;

        lambda.invoke(masteryParams, (error, data) => {
            if (error) {
                this.masteryDataLoaded = true;
                prompt(error);
            } else {
                masteryResults = JSON.parse(data.Payload);
                this.ChampionPoints = masteryResults.championPoints;
                this.ChampionID = masteryResults.championId;
                this.ChampionLevel = masteryResults.championLevel;
                this.masteryDataLoaded = true;
                console.log(masteryResults)
                }
            })
        }
    },
    computed: {
        summonerID() {
            return this.$store.getters.getEncryptedSummonerID
        }
    },
    watch: {
        summonerID(oldSumID, newSumID) {
            this.getChampionMasteryByID()
        }
    }
}
</script>

<style scoped>

</style>