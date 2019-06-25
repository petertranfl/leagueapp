<template>
    <div v-if="masteryDataLoaded">
        <div class="championsContainer">
            <div>
                    <img :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' + masteryResults[0].championId + '.png'"/>
                    <h3>{{masteryResults[0].championPoints}} Pts</h3>
            </div>
            <div>
                     <img :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' + masteryResults[1].championId + '.png'"/>
                     <h3>{{masteryResults[1].championPoints}} Pts</h3>
            </div>
            <div>
                     <img :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' + masteryResults[2].championId + '.png'"/>
                     <h3>{{masteryResults[2].championPoints}} Pts</h3>
            </div>
        </div>
    </div>
</template>

<script>
import AWS from 'aws-sdk';

export default {
    name: 'ChampionMastery',
    data() {
        return {
            masteryDataLoaded: false,
            masteryResults: []
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
            Payload: '{"SummonerID?": ' + `"` + this.summonerID + '"}'
            };

        lambda.invoke(masteryParams, (error, data) => {
            if (error) {
                prompt(error);
            } else {
                this.masteryResults = JSON.parse(data.Payload)
                this.masteryDataLoaded = true;
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
    }

    img {
        max-width: 90%;
        height: auto;
    }
</style>