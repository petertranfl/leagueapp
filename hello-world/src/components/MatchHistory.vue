<template>
  <div v-if="matchHistoryDataLoaded">
    <div class="matchHistoryContainer">
      <table class="table">
        <tr v-for="(matches, index) in matchHistoryData.matches" :key="index">
          <td>
            <img
              :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/v1/champion-icons/' + matchHistoryData.matches[index].champion + '.png'"
            />
          </td>
          <td>
            <img
              :src="'http://raw.communitydragon.org/latest/plugins/rcp-be-lol-game-data/global/default/assets/ranked/positions/rankposition_' + leagueTier + '-' + lanePosition(index) + '.png'"
              alt
              id="roleImg"
            />
          </td>
          <td>{{getSummonerMatchStats(index)}}</td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
import AWS from "aws-sdk";

export default {
  name: "MatchHistory",
  data() {
    return {
      matchHistoryDataLoaded: false,
      matchHistoryData: Object
    };
  },
  computed: {
    accountID() {
      return this.$store.getters.getAccountID;
    },
    leagueTier() {
      return this.$store.getters.getLeagueTier;
    },
    KDA(index) {
      return this.getSummonerMatchStats(index).deaths;
    }
  },
  watch: {
    accountID() {
      this.getMatchHistory("0", "10");
    }
  },
  methods: {
    getMatchHistory(beginIndex, endIndex) {
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
          '{"AccountID?": "' +
          this.accountID +
          '", "BeginIndex?": "' +
          beginIndex +
          '", "EndIndex?": "' +
          endIndex +
          '"}'
      };

      let getMatchData = lambda.invoke(matchHistoryParams).promise();
      getMatchData
        .then(data => {
          this.matchHistoryData = JSON.parse(data.Payload);
          let promises = [];
          for (let i = 0; i < this.matchHistoryData.matches.length; i++) {
            let promise = this.getMatchDetails(
              i,
              this.matchHistoryData.matches[i].gameId
            );
            promises.push(promise);
          }
          Promise.all(promises).then(value => {
            this.matchHistoryDataLoaded = true;
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
          this.matchHistoryData.matches[index] = Object.assign(
            this.matchHistoryData.matches[index],
            JSON.parse(data.Payload)
          );
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
          console.log(this.accountID)
        if (this.matchHistoryData.matches[index].participantIdentities[j].player.accountId == this.accountID) {
          tempParticipantID = this.matchHistoryData.matches[index].participantIdentities[j].participantId;
        }
      }
      for (let k = 0; k < this.matchHistoryData.matches[index].participants.length; k++) {
        if (this.matchHistoryData.matches[index].participants[k].participantId == tempParticipantID) {
            console.log(this.matchHistoryData.matches[index].participants[k].championId)
            return this.matchHistoryData.matches[index].participants[k].stats.kills;
        }
      }
    }
  }
};
</script>

<style scoped>
.matchHistoryContainer {
  display: flex;
  position: relative;
  justify-content: center;
}

.table {
}

td {
  padding-bottom: 1em;
}

img {
  width: 40%;
  height: auto;
}

#roleImg {
  margin-left: -2em;
  width: 120%;
  height: auto;
}
</style>