<template>
    <div v-if="masteryDataLoaded">
         <h1>Champion Mastery</h1>
        <img :src="'https://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/assets/loadouts/summonerbanners/frames/base/leagueclient/bannerframe_01_inventory.png'" 
             id="masteryTitle"
        />
        <div class="championsContainer">
            <div>
                    <img :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' + masteryResults[0].championId + '.png'"/>
                    <h2>{{champ(this.masteryResults[0].championId)}}</h2>
                    <h3>{{masteryResults[0].championPoints}} Pts</h3>
            </div>
            <div>
                     <img :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' + masteryResults[1].championId + '.png'"/>
                     <h2>{{champ(this.masteryResults[1].championId)}}</h2>
                     <h3>{{masteryResults[1].championPoints}} Pts</h3>
            </div>
            <div>
                     <img :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' + masteryResults[2].championId + '.png'"/>
                     <h2>{{champ(this.masteryResults[2].championId)}}</h2>
                     <h3>{{masteryResults[2].championPoints}} Pts</h3>
            </div>
        </div>
    </div>
</template>

<script>
import AWS from 'aws-sdk';
import Champions from '../../public/championKeys.json';

export default {
    name: 'ChampionMastery',
    data() {
        return {
            masteryDataLoaded: false,
            masteryResults: [],
            champions: Champions,
        }
    },
    computed: {
        summonerID() {
            return this.$store.getters.getEncryptedSummonerID
        },
    },
    watch: {
        summonerID() {
            this.getChampionMasteryByID()
        }
    },
    methods: {
        getChampionMasteryByID() {
            console.log(this.champions)

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
            //console.log(this.masteryResults)
        })
        getMastery.catch(error => {
            prompt(error)
         })
        },
    champ(Id) {
            return this.champions.keys[Id]
        }
    },
}
</script>

<style scoped>
    h1 {
        margin-top: 1rem;
        color: #fad161;
        font-weight: bolder;
        text-shadow: 1px 0 0 rgb(49, 42, 9), 0 -1px 0 rgb(49, 42, 9), 0 1px 0 rgb(49, 42, 9), -1px 0 0 rgb(49, 42, 9);
    }

    h2 {
        color: whitesmoke;
        text-shadow: 1px 0 0 rgb(46, 46, 46), 0 -1px 0 rgb(46, 46, 46), 0 1px 0 rgb(46, 46, 46), -1px 0 0 rgb(46, 46, 46);
    }

    h3 {
        color:#fad161;
        text-shadow: 1px 0 0 rgb(49, 42, 9), 0 -1px 0 rgb(49, 42, 9), 0 1px 0 rgb(49, 42, 9), -1px 0 0 rgb(49, 42, 9);
    }
    .championsContainer {
        display: flex;
        position: relative;
        justify-content: space-evenly;
        align-items: baseline;
        margin-top: -2rem;
    }
    #masteryTitle {
        height: 8rem;
        width: 25rem;
        margin-top: -3.8rem;
        margin-bottom: 3rem;
    }
    img {
        max-width: 90%;
        height: auto;
    }
</style>