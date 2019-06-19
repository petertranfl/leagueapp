<template>
    <div v-if="matchHistoryDataLoaded">
        <div class="matchHistoryContainer">
           
        </div>
    </div>
</template>

<script>
import AWS from 'aws-sdk'
export default {
    name: 'MatchHistory',
    data() {
        return {
            matchHistoryDataLoaded: false,
        }
    },
    computed: {
        accountID() {
            return this.$store.getters.getAccountID
        }
    },
    watch: {
        accountID(oldSumID, newSumID) {
            this.getMatchHistory()
        }
    },
    methods: {
        getMatchHistory() {

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
            Payload: '{"AccountID?": ' + `"` + this.accountID + '"}'
            };

        lambda.invoke(masteryParams, (error, data) => {
            if (error) {
                prompt(error);
            } else {
                this.masteryResults = JSON.parse(data.Payload);
                this.matchHistoryDataLoaded = true;
                }
            })
        }
    },
}
</script>

<style scoped>
    .championsContainer {
        display: flex;
        position: relative;
        justify-content: space-evenly;
        align-items: baseline;
        height:2em;
        width: auto;
    }
</style>