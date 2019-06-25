import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const state = {
    encryptedSummonerID: String,
    accountID: String,
    leagueTier : String,
}

const getters = {
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