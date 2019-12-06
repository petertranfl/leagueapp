<template>
    <div class="landingPage">
        <div id="landingText">
            <h1> Welcome to the SummonerSpy! <br> Look up Summoners by searching for their Summoner Name below. </h1>
        </div>
    </div>
</template>

<script>
export default {
    
}
</script>

<style scoped>
    .landingPage {
        display: flex;
        height: 100vh;
    }
    #landingText {
        margin-top: 3rem;
        color: #fad161;
        text-shadow: 1px 0 0 rgb(49, 42, 9), 0 -1px 0 rgb(49, 42, 9), 0 1px 0 rgb(49, 42, 9), -1px 0 0 rgb(49, 42, 9);
    }
</style>