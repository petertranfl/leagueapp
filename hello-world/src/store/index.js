import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const state = {
    encryptedSummonerID: String
}

const getters = {
    getEncryptedSummonerID(state) {
        return state.encryptedSummonerID
    }
}

const mutations = {
    saveSummonerID(state, sID) {
        state.encryptedSummonerID = sID
    }
}

const actions = {
    commitSummonerID(context, sID) {
    context.commit('saveSummonerID', sID)
    }
}



export default new Vuex.Store({
    state,
    getters,
    mutations,
    actions
})