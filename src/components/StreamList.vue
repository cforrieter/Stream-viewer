<template>
  <div>
    <h1>Stream list page</h1>
    <div class="flex">
      <StreamPreview v-for="(stream, index) in streams" :key="index" :name="stream.channelTitle" :title="stream.title"  :thumbnail="stream.thumbnails.medium.url" :videoId="stream.videoId"/>
    </div>
  </div>
</template>

<script>
import StreamPreview from './StreamPreview.vue'
import buildApiRequest from "@/lib/googleApi.js"

export default {
  name: "StreamList",
  components: {
    StreamPreview
  },
  data() {
    return {
      streams: []
    }
  },
  methods: {
    fetchStreams() {

      this.$getGapiClient()
        .then(gapi => {
            
              return buildApiRequest('GET',
              '/youtube/v3/search',
              {'maxResults': '50',
              'part': 'snippet',
              'q': 'gaming',
              'eventType': 'live',
              'order': 'viewCount',
              'relevanceLanguage': 'en',
              'type': 'video'})

        }).then(response => {
          console.log(response)
          this.streams = response.result.items.map( item => {
            let stream = item.snippet
            stream["videoId"] = item.id.videoId
            return stream
          })
        })
    }
  },
  mounted() {
    this.fetchStreams();
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.flex {
  display: flex;
  width: 80%;
  flex-wrap: wrap;
  margin: 0 auto;
  justify-content: space-around
}
</style>
