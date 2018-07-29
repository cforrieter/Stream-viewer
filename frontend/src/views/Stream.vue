<template>
  <div>
    <div style="display:flex;">
      <youtube :video-id="videoId" :player-vars="{ autoplay: 1}" ref="youtube"></youtube>
      <ChatBox :comments="comments" :live-chat-id="liveChatId" />
    </div>
  </div>
</template>

<script>
import buildApiRequest from '@/lib/googleApi.js'

import ChatBox from '@/components/ChatBox.vue'

export default {
  name: "Stream",
  props: {
    videoId: String,
  },
  components: {
    ChatBox
  },
  data() {
    return {
      liveChatId: "",
      comments: [],
      nextPageToken: "",
      commentsTimeout: undefined
    }
  },
  computed: {
    player () {
      return this.$refs.youtube.player
    }
  },
  mounted() {
    this.getLiveChatId()
    
  },
  beforeDestroy() {
    clearTimeout(this.commentsTimeout)
  },
  methods: {
    getLiveChatId() {
      this.$getGapiClient()
        .then(gapi => {
              return buildApiRequest('GET',
              '/youtube/v3/videos',
              {'maxResults': '1',
              'part': 'liveStreamingDetails',
              'id': this.videoId})
          
        }).then(response => {
          console.log(response.result.items[0].liveStreamingDetails.activeLiveChatId)
          this.liveChatId = response.result.items[0].liveStreamingDetails.activeLiveChatId;
          this.fetchComments()
        })
    },
    fetchComments() {
       this.$getGapiClient()
        .then(gapi => {
              return buildApiRequest(
                'GET', '/youtube/v3/liveChat/messages',
              {
                'part': 'snippet, authorDetails',
                'liveChatId': this.liveChatId,
                'pageToken': this.nextPageToken
              }
            );
          
        }).then(response => {
          this.nextPageToken = response.result.nextPageToken
          response.result.items.forEach( item => {
            this.comments.push({
              "message": item.snippet.displayMessage,
              "author": item.authorDetails.displayName
            })
          })

          this.commentsTimeout = setTimeout(this.fetchComments, response.result.pollingIntervalMillis || 5000)
        })
    }
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
