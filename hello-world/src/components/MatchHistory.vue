<template>
  <div v-if="matchHistoryDataLoaded">
    <div class="matchHistoryContainer">
      <h1>Match History</h1>
      <div class="matchHistoryCard" v-for="(matches, index) in matchHistoryData.matches" :key="index">
        <div class="timeStamp">{{TimeStamp(index)}}</div>
        <!-- <div class="winLoss">{{win(index)}}</div> -->
        <img
          :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' + matchHistoryData.matches[index].champion + '.png'"
          class="champImg" />
        <img
          :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/assets/ranked/positions/rankposition_' + leagueTier + '-' + lanePosition(index) + '.png'"
          @error="HandleBrokenImage" alt class="roleImg" />
        <img
          :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/data/spells/icons2d/' + SummonerSpell1(index) + '.png'"
          alt class="summonerSpellIcon1" />
        <img
          :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/data/spells/icons2d/' + SummonerSpell2(index) + '.png'"
          alt class="summonerSpellIcon2" />
        <img :src="'http://raw.communitydragon.org/latest/plugins' + Keystone(index)" alt class="keystone" />
        <img :src="'http://raw.communitydragon.org/latest/plugins' + SubStyle(index)" alt class="substyle" />
        <div class="kdaStats">
          <div class="kda">
            <p class="kills">{{getSummonerMatchStats(index).stats.kills}}</p>
            <p class="kdaDivider">/</p>
            <p class="deaths">{{getSummonerMatchStats(index).stats.deaths}}</p>
            <p class="kdaDivider">/</p>
            <p class="assists">{{getSummonerMatchStats(index).stats.assists}}</p>
          </div>
          <p class="ratio">{{KDA(index)}} KDA</p>
        </div>
        <div class="itemsBox">
          <img :src="'/item/' + getSummonerMatchStats(index).stats.item0 + '.png'" class="item" @error="noItem" />
          <img :src="'/item/' + getSummonerMatchStats(index).stats.item1 + '.png'" class="item" @error="noItem" />
          <img :src="'/item/' + getSummonerMatchStats(index).stats.item2 + '.png'" class="item" @error="noItem" />
          <img :src="'/item/' + getSummonerMatchStats(index).stats.item3 + '.png'" class="item" @error="noItem" />
          <img :src="'/item/' + getSummonerMatchStats(index).stats.item4 + '.png'" class="item" @error="noItem" />
          <img :src="'/item/' + getSummonerMatchStats(index).stats.item5 + '.png'" class="item" @error="noItem" />
        </div>
      </div>
    </div>
  </div>

</template>

<script>
import AWS from "aws-sdk";
import { green, red } from 'color-name';

export default {
  name: "MatchHistory",
  data() {
    return {
      matchHistoryDataLoaded: false,
      matchHistoryData: Object,
      // item: LeagueItems
    };
  },
  computed: {
    accountID() {
      return this.$store.getters.getAccountID;
    },
    leagueTier() {
      return this.$store.getters.getLeagueTier;
    },
  },
  watch: {
    accountID() {
      this.getMatchHistory("0", "10");
    }
  },
  methods: {
    getMatchHistory(beginIndex, endIndex) {
      this.matchHistoryDataLoaded = false;
      AWS.config.credentials = new AWS.CognitoIdentityCredentials(
        {
          IdentityPoolId: "us-east-1:98b70204-c8a3-4336-b9be-ea2f4393f3b1"
        },
        {
          region: "us-east-1"
        }
      );

      //create AWS service object
      const lambda = new AWS.Lambda({ region: "us-east-1" });
      //create JSON object for parameters for invoking Lambda function
      const matchHistoryParams = {
        FunctionName: "getMatchHistory",
        InvocationType: "RequestResponse",
        LogType: "None",
        Payload:
          '{"AccountID?": "' + this.accountID + '", "BeginIndex?": "' + beginIndex + '", "EndIndex?": "' +endIndex +'"}'
        };

      let getMatchData = lambda.invoke(matchHistoryParams).promise();
      getMatchData
        .then(data => {
          this.matchHistoryData = JSON.parse(data.Payload);
          let promises = [];
          for (let i = 0; i < this.matchHistoryData.matches.length; i++) {
            let promise = this.getMatchDetails(i, this.matchHistoryData.matches[i].gameId);
            promises.push(promise);
            }
          Promise.all(promises).then(value => {
            console.log(this.matchHistoryData)
            this.matchHistoryDataLoaded = true;
            this.$emit('loaded', false)
          });
        })
        .catch(error => {
          prompt(error);
        });
    },
    getMatchDetails(index, matchID) {
      return new Promise((resolve, reject) => {
        //load AWS credentials
        AWS.config.credentials = new AWS.CognitoIdentityCredentials(
          {
            IdentityPoolId: "us-east-1:98b70204-c8a3-4336-b9be-ea2f4393f3b1"
          },
          {
            region: "us-east-1"
          }
        );
        //create AWS service object
        const lambda = new AWS.Lambda({ region: "us-east-1" });
        //create JSON object for parameters for invoking Lambda function
        const matchDetailsParams = {
          FunctionName: "getMatchDetails",
          InvocationType: "RequestResponse",
          LogType: "None",
          Payload: '{"MatchID?": "' + matchID + '"}'
        };
        let getMatchDetailsData = lambda.invoke(matchDetailsParams).promise();
        getMatchDetailsData.then(data => {
          this.matchHistoryData.matches[index] = Object.assign(this.matchHistoryData.matches[index], JSON.parse(data.Payload));
          });
        getMatchDetailsData.catch(error => {
          prompt(error);
        });
        resolve(getMatchDetailsData);
      });
    },
    lanePosition(index) {
      if (this.matchHistoryData.matches[index].lane == "BOTTOM") {
        return (this.matchHistoryData.matches[index].lane = "bot");
      } else {
        return this.matchHistoryData.matches[index].lane.toLowerCase();
      }
    },
    getSummonerMatchStats(index) {
      var tempParticipantID;
      for (let j = 0; j < this.matchHistoryData.matches[index].participantIdentities.length; j++) {
        if (this.matchHistoryData.matches[index].participantIdentities[j].player.accountId == this.accountID) {
          tempParticipantID = this.matchHistoryData.matches[index].participantIdentities[j].participantId;
        }
      }
      for (let k = 0; k < this.matchHistoryData.matches[index].participants.length; k++) {
        if (this.matchHistoryData.matches[index].participants[k].participantId == tempParticipantID) {
            return this.matchHistoryData.matches[index].participants[k];
        }
      }
    },
    win(index) {
      let win = this.getSummonerMatchStats(index).stats.win
      if (win) {
        winLoss.style.setProperty('--element-color', green)
        return "VICTORY"
      } else {
        winLoss.style.setProperty('--element-color', red)
        return "DEFEAT"
      }
    },
     KDA(index) {
      let kda = ((this.getSummonerMatchStats(index).stats.kills + this.getSummonerMatchStats(index).stats.assists) / this.getSummonerMatchStats(index).stats.deaths).toFixed(1);
      if (kda == "Infinity") {
        return "PERFECT"
      }
      if (kda == "NaN") {
        return 0
      }
      return kda
    },
    SummonerSpell1(index) {
      switch (this.getSummonerMatchStats(index).spell1Id) {
        case 1:
          return "summoner_boost";
        case 3:
          return "summoner_exhaust";
        case 4:
          return "summoner_flash";
        case 6:
          return "summoner_haste";
        case 7:
          return "summoner_heal";
        case 11:
          return "summoner_smite";
        case 12:
          return "summoner_teleport";
        case 13:
          return "summonermana";
        case 14:
          return "summonerignite";
        case 21:
          return "summonerbarrier";
        case 32:
          return "summoner_mark";
      }
    },
    SummonerSpell2(index) {
      switch (this.getSummonerMatchStats(index).spell2Id) {
        case 1:
          return "summoner_boost";
        case 3:
          return "summoner_exhaust";
        case 4:
          return "summoner_flash";
        case 6:
          return "summoner_haste";
        case 7:
          return "summoner_heal";
        case 11:
          return "summoner_smite";
        case 12:
          return "summoner_teleport";
        case 13:
          return "summonermana";
        case 14:
          return "summonerignite";
        case 21:
          return "summonerbarrier";
        case 32:
          return "summoner_mark";
      }
    },
    Keystone(index) {
      switch (this.getSummonerMatchStats(index).stats.perk0) {
        case 8005:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/precision/presstheattack/presstheattack.png";
        case 8008:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/precision/lethaltempo/lethaltempotemp.png";
        case 8021:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/precision/fleetfootwork/fleetfootwork.png";
        case 8010:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/precision/conqueror/conqueror.png";
        case 8112:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/domination/electrocute/electrocute.png";
        case 8124:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/domination/predator/predator.png";
        case 8128:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/domination/darkharvest/darkharvest.png";
        case 9923:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/domination/hailofblades/hailofblades.png";
        case 8214:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/sorcery/summonaery/summonaery.png";
        case 8230:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/sorcery/phaserush/phaserush.png";
        case 8437:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/resolve/graspoftheundying/graspoftheundying.png";
        case 8439:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/resolve/veteranaftershock/veteranaftershock.png";
        case 8465:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/resolve/guardian/guardian.png";
        case 8351:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/inspiration/glacialaugment/glacialaugment.png";
        case 8360:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/inspiration/unsealedspellbook/unsealedspellbook.png";
        case 8358:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/inspiration/masterkey/masterkey.png";
      }
    },
    SubStyle(index) {
      switch (this.getSummonerMatchStats(index).stats.perkSubStyle) {
        case 8400:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/7204_resolve.png";
        case 8100:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/7200_domination.png";
        case 8000:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/7201_precision.png";
        case 8200:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/7202_sorcery.png";
        case 8300:
          return "/rcp-be-lol-game-data/global/default/v1/perk-images/styles/7203_whimsy.png"
      }
    },
    TimeStamp(index) {
      //time.now() - this = time in seconds.
      //turn this into hours and return
      let d = new Date();
      let currentTime = d.getTime();
      let timeElapsed = currentTime - this.matchHistoryData.matches[index].GameCreation;
      let seconds = (timeElapsed / 1000).toFixed(0);
      let minutes = (timeElapsed / (1000 * 60)).toFixed(0);
      let hours = (timeElapsed / (1000 * 60 * 60)).toFixed(0);
      let days = (timeElapsed / (1000 * 60 * 60 * 24)).toFixed(0);

        if (seconds < 60) {
            return seconds + " Secs ago";
        } else if (minutes < 60) {
            return minutes + " Mins ago";
        } else if (hours < 24) {
            return hours + " Hrs ago";
        } else {
            return days + " Days ago"
        }
    },
    HandleBrokenImage(event) {
      event.target.src = "questionMark.png"
    },
    noItem(event) {
      event.target.src = "/item/no_item.png";
      event.target.className += " noItem";
    }
  }
};
</script>

<style scoped>
h1 {
  color: #fad161;
  text-shadow: 1px 0 0 rgb(49, 42, 9), 0 -1px 0 rgb(49, 42, 9), 0 1px 0 rgb(49, 42, 9), -1px 0 0 rgb(49, 42, 9);
}
.matchHistoryContainer {
  display: flex;
  flex-direction: column;
  position: relative;
  align-items: center;
  margin-top: 3rem;
}
.matchHistoryCard {
  display: flex;
  flex-direction: row;
  align-items: center;
  width: 50rem;
  height: 6rem;
  padding: 1rem;
  background: rgba(194, 194, 194, 0.226);
  border-style: ridge;
  border-color: silver;
  border-width: 2px;
  border-radius: 5px;
}
.timeStamp {
  font-size: 1.1rem;
  color: rgb(233, 233, 233);
  font-weight: bold;
  width: 6rem;
}
.win-loss {
  font-size: 1.1rem;
  color: var(--element-color);
}
.champImg {
  height: 4.5rem;
  width: 4.5rem;
  margin-right: 3rem;
  margin-left: 1rem;
}
.summonerSpellIcon1 {
  width: 2rem;
  height: 2rem;
  margin-top: -2rem;
}
.summonerSpellIcon2 {
  width: 2rem;
  height: 2rem;
  margin-left: -2rem;
  margin-bottom: -2rem;
}
.roleImg {
  width: 4rem;
  height: 4rem;
  margin-right: 3rem;
}
.keystone {
  margin-top: -2rem;
  width: 2.5rem;
  height: 2.5rem;
}
.substyle{
  margin-bottom: -1.8rem;
  margin-left: -2.2rem;
  width: 1.8rem;
  height: 1.8rem;
}
.kdaStats {
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
  height: 6rem;
  width: 10rem;
  font-size: 1.3rem;
  margin-left: 1rem;
}
.kills {
  color: white;
}
.deaths {
  color: rgb(219, 42, 42);
}
.assists {
  color: white;
}
.kdaDivider {
  color: black;
  margin-left: 0.3rem;
  margin-right: 0.3rem;
}
.kda {
  display: flex;
  flex-direction: row;
}
.ratio {
  color: rgb(223, 14, 195);
}
.item {
  border-radius: 8px;
  height: 2rem;
  width: 2rem;
}
.noItem {
  opacity: 0.3;
}
</style>