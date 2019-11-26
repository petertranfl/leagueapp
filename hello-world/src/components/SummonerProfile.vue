<template>
    <div class="main-container">
        <div v-if="summonerDataLoaded">
            <div class="profile-container">
                <img id="rankedCrest" :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/content/src/leagueclient/rankedcrests/' + leagueNumber + '_' + soloSummonerLeague + '/images/' + soloSummonerLeague + '_base_sheeng.png'" alt=""/>
                <img id="profileIcon" :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/profile-icons/' + profileIconID + '.jpg'"/>
                <img id="banner" :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/assets/regalia/banners/backgrounds/solidbanner_still.png'"/>
                <img id="trim" :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/assets/regalia/banners/trims/' + banner + '.png'"/>
            </div>
            <div class="text-container">
                <h2>{{summonerName}}</h2>
                <h3>Lvl: {{summonerLevel}}</h3>
            </div>
        </div>
    </div>
</template>

<script>
import AWS from 'aws-sdk';

export default {
    name: 'SummonerProfile',
    data() {
        return {
            summonerName: String,
            summonerLevel: Number,
            userSummonerName: "lolitspetey",
            profileIconID: Number,
            encryptedSummonerID: String,
            flexSummonerLeague: String,
            soloSummonerLeague: String,
            leagueResults: [],
            summonerDataLoaded: false,
        }
    },
    created() {
        this.getNASummonerData();
    },
    watch: {
        encryptedSummonerID(){
            this.getNALeagueData();
        }
    },
    computed: {
        leagueNumber() {
            switch (this.soloSummonerLeague) {
                case "iron":
                    return "01";
                case "bronze":
                    return "02";
                case "silver":
                    return "03";
                case "gold":
                    return "04";
                case "platinum":
                    return "05";
                case "diamond":
                    return "06";
                case "master":
                    return "07";
                case "grandmaster":
                    return "08";
                case "challenger":
                    return "09";
            }
        },
        banner() {
            switch(this.soloSummonerLeague) {
                case "iron":
                    return "trim_iron";
                case "bronze":
                    return "trim_bronze";
                case "silver":
                    return "trim_silver";
                case "gold":
                    return "trim_gold";
                case "platinum":
                    return "trim_plat";
                case "diamond":
                    return "trim_diamond";
                case "master":
                    return "trim_master";
                case "grandmaster":
                    return "trim_grandmaster";
                case "challenger":
                    return "trim_challenger";
                default:
                    return "defaulttrim";
            }
        }
    },
    methods: {
        getNASummonerData() {

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
            const summonerParams = {
                FunctionName: 'getNASummonerDataByName',
                InvocationType: 'RequestResponse',
                LogType: 'None',
                Payload: '{"SummonerName?": ' + `"` + this.userSummonerName + '"}'
            };
            //create variable to hold data returned by Lambda function
            let summonerResults;

            //Calls Lambda function 'GetSummonerDataByName' with given reqParams
            let getData = lambda.invoke(summonerParams).promise()
            getData.then(data => {
                summonerResults = JSON.parse(data.Payload)
                this.summonerName = summonerResults.name;
                this.summonerLevel = summonerResults.summonerLevel;
                this.profileIconID = summonerResults.profileIconId;
                this.encryptedSummonerID = summonerResults.id;
                this.$store.dispatch('commitSummonerID', summonerResults.id);
                this.$store.dispatch('commitAccountID', summonerResults.accountId)
            })
            getData.catch(error => {
                prompt(error)
            })
        },
        getNALeagueData() {
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
            const leagueParams = {
                FunctionName: 'getNALeagueDataByID',
                InvocationType: 'RequestResponse',
                LogType: 'None',
                Payload: '{"SummonerID?": ' + `"` + this.encryptedSummonerID + '"}'
            };

            //Calls Lambda function 'GetSummonerDataByName' with given reqParams
            let getLeagueData = lambda.invoke(leagueParams).promise()
            getLeagueData.then(data => {
                this.leagueResults = JSON.parse(data.Payload);
                this.sortNALeagueData(this.leagueResults);
                this.$store.dispatch('commitLeagueTier', this.soloSummonerLeague);
                this.summonerDataLoaded = true;
                })
            getLeagueData.catch(error => {
                prompt(error)
            })
        },
        sortNALeagueData(data) {
            for (let i = 0; i < data.length; i++) {
                if (data[i].queueType == "RANKED_SOLO_5x5") {
                    this.soloSummonerLeague = data[i].tier.toLowerCase();
                }
                else this.flexSummonerLeague = data[i].tier.toLowerCase();
            }
        }
    },
}
</script>

<style scoped>
    .main-container {
        display: flex;
        height: 500px;
        flex-direction: column;
        align-items: center;
        margin-bottom: 2em;
    }

    .profile-container {
        position: relative;
        width: 500px;
        height: 350px;
    }

    .profile-container #rankedCrest {
        position: absolute;
        top: 1.3em;
        left: 6.25em;
        z-index: 2;
    }
    .profile-container #profileIcon {
        position: absolute;
        width: 30%;
        height: auto;
        border-radius: 50%;
        top: 7em;
        right: 10.9em;
    }
    .profile-container #banner {
        height: 29em;
        width: 13.1em;
        z-index: -1;
    }

    .profile-container #trim {
        height: 6.6em;
        width: auto;
        position: absolute;
        top: 23.6em;
        left: 9em;
    }

    .text-container {
        position: relative
    }
    
    h2 {
        margin-top: -1.3em;
        color: #fad161
    }
    h3 {
        color: #fad161
    }
</style>