import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const state = {
    encryptedSummonerID: String,
    accountID: String,
}

const getters = {
    getEncryptedSummonerID(state) {
        return state.encryptedSummonerID
    },
    getAccountID(state) {
        return state.accountID
    }
}

const mutations = {
    saveSummonerID(state, sID) {
        state.encryptedSummonerID = sID
    },
    saveAccountID(state, aID) {
        state.accountID = aID
    }
}

const actions = {
    commitSummonerID(context, sID) {
    context.commit('saveSummonerID', sID)
    },
    commitAccountID(context, aID) {
        context.commit('saveAccountID', aID)
    }
}



export default new Vuex.Store({
    state,
    getters,
    mutations,
    actions
})