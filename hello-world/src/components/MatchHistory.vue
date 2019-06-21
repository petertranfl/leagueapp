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
            matchHistoryData: Object,
        }
    },
    computed: {
        accountID() {
            return this.$store.getters.getAccountID
        }
    },
    watch: {
        accountID(oldSumID, newSumID) {
            this.getMatchHistory("0", "10")
        }
    },
    methods: {
        getMatchHistory(beginIndex, endIndex) {

        //load AWS credentials
        AWS.config.region = 'us-east-2';
        AWS.config.credentials = new AWS.CognitoIdentityCredentials({
            IdentityPoolId: 'us-east-2:20fd57cf-7bb6-4352-b5ce-69eca3907336'
            });

        //create AWS service object
        const lambda = new AWS.Lambda({region: 'us-east-2'});
        //create JSON object for parameters for invoking Lambda function
        var matchHistoryParams = {
            FunctionName: 'getMatchHistory',
            InvocationType: 'RequestResponse',
            LogType: 'None',
            Payload: '{"AccountID?": "' + this.accountID + '", "BeginIndex?": "' + beginIndex + '", "EndIndex?": "' + endIndex + '"}'
        }

        lambda.invoke(matchHistoryParams, (error, data) => {
            if (error) {
                prompt(error);
            } else {
                this.matchHistoryData = JSON.parse(data.Payload);
                for (var i = 0; i < this.matchHistoryData.matches.length; i++) {
                        this.getMatchDetails(i, this.matchHistoryData.matches[i].gameId, this.matchHistoryData.matches[i]);
                    }
                }
                console.log(this.matchHistoryData)
            })
        },
        getMatchDetails(index, matchID, match) {
            //load AWS credentials
        AWS.config.region = 'us-east-2';
        AWS.config.credentials = new AWS.CognitoIdentityCredentials({
            IdentityPoolId: 'us-east-2:20fd57cf-7bb6-4352-b5ce-69eca3907336'
            });

        //create AWS service object
        const lambda = new AWS.Lambda({region: 'us-east-2'});
        //create JSON object for parameters for invoking Lambda function
        var matchDetailsParams = {
            FunctionName: 'getMatchDetails',
            InvocationType: 'RequestResponse',
            LogType: 'None',
            Payload: '{"MatchID?": "' + matchID + '"}'
        };
        lambda.invoke(matchDetailsParams, (error, data) => {
            if (error) {
                prompt(error);
            } else {
                this.matchHistoryData.matches[index].matchDetails = JSON.parse(data.Payload);
                }
            })
        },
    }
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