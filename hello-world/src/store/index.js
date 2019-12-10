import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const state = {
    summonerName: String,
    encryptedSummonerID: String,
    accountID: String,
    leagueTier : String,
}

const getters = {
    getSummonerName(state) {
        return state.summonerName
    },
    getEncryptedSummonerID(state) {
        return state.encryptedSummonerID
    },
    getAccountID(state) {
        return state.accountID
    },
    getLeagueTier(state) {
        return state.leagueTier
    }
}

const mutations = {
    saveSummonerName(state, name) {
        state.summonerName = name
    },
    saveSummonerID(state, sID) {
        state.encryptedSummonerID = sID
    },
    saveAccountID(state, aID) {
        state.accountID = aID
    },
    saveLeagueTier(state, leagueTier) {
        state.leagueTier = leagueTier
    }
}

const actions = {
    commitSummonerName(context, name) {
        context.commit('saveSummonerName', name)
    },
    commitSummonerID(context, sID) {
    context.commit('saveSummonerID', sID)
    },
    commitAccountID(context, aID) {
        context.commit('saveAccountID', aID)
    },
    commitLeagueTier(context, leagueTier) {
        context.commit('saveLeagueTier', leagueTier)
    }
}



export default new Vuex.Store({
    state,
    getters,
    mutations,
    actions
})