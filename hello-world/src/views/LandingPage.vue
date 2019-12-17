<template>
    <div class="landingPage">
        <div class="landingText">
            <h1> Welcome to the SummonerSpy! <br> <br> Look up a Summoner by searching for their Summoner Name below. </h1>
        </div>
        <div class="searchBar">
            <input v-model="summonerName" type="text" v-on:keyup.enter="submitSearch">
        </div>
    </div>
</template>

<script>
export default {
    data: function() {
        return {
        summonerName: ""
        }
    },
    methods: {
        submitSearch() {
            this.$store.commit('saveSummonerName', this.summonerName)
            this.$router.push('/profile/')
        },
    }
}
</script>

<style scoped>
    @font-face {
        font-family: 'Friz Quadrata Regular'; /* League of Legends Font (almost) */
        src: url('../assets/Friz Quadrata Regular.woff') format('truetype'); /* Safari, Android, iOS */
        src: url('../assets/Friz Quadrata Regular.woff') format('woff') /* Modern Browsers */
    }
    .landingPage {
        display: flex;
        flex-direction: column;
        height: 100vh;
    }
    .landingText {
        justify-content: center;
        width: 100vw;
        margin-top: 3rem;
        color: #fad161;
        text-shadow: 1px 0 0 rgb(77, 65, 1), 0 -1px 0 rgb(77, 65, 1), 0 1px 0 rgb(77, 65, 1), -1px 0 0 rgb(77, 65, 1);
    }
    .searchBar {
        margin-top: 10rem;
    }
    .searchBar input {
        height: 1.5rem;
        width: 20rem;
        font-size: 1.3rem;
        border-style: none;
        border-radius: 0.5rem;
        color: #fad161;
        /*border-color:rgb(77, 65, 1);*/
        background: rgba(255, 255, 255, 0.205);
    }
</style>