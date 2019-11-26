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
        summonerID() {
            this.getChampionMasteryByID()
        }
    },
    methods: {
        getChampionMasteryByID() {

        //load AWS credentials
        AWS.config.credentials = new AWS.CognitoIdentityCredentials({
                IdentityPoolId: 'us-east-1:98b70204-c8a3-4336-b9be-ea2f4393f3b1',
            },
            {
                region: 'us-east-1'
            });

        //create AWS service object
        const lambda = new AWS.Lambda({region: 'us-east-1'});
        //create JSON object for parameters for invoking Lambda function
        const masteryParams = {
            FunctionName: 'getChampionMasteryByID',
            InvocationType: 'RequestResponse',
            LogType: 'None',
            Payload: '{"SummonerID?": ' + `"` + this.summonerID + '"}'
            };

        let getMastery = lambda.invoke(masteryParams).promise()
        getMastery.then(data => {
            this.masteryResults = JSON.parse(data.Payload)
            this.masteryDataLoaded = true;
        })
        getMastery.catch(error => {
            prompt(error)
        })
        }
    },
}
</script>

<style scoped>
    h3 {
        color:#fad161
    }
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