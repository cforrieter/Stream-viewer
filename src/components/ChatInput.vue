<template>
  <div>
    <textarea placeholder="Send a message" v-model="messageValue" class="chatInput" @keyup.enter="submitMessage" type="text" />
  </div>
</template>

<script>
import buildApiRequest from '@/lib/googleApi.js'

export default {
  name: "ChatInput",
  data() {
    return {
      messageValue: ""
    }
  },
  props: {
    liveChatId: String
  },
  methods: {
    submitMessage() {
      let messageText = this.messageValue;
      this.messageValue = "";
      this.$getGapiClient()
        .then(gapi => {
            
              return buildApiRequest(
                'POST', '/youtube/v3/liveChat/messages',
              {
                'part': 'snippet',
              },
              {
                'snippet': {
                  'liveChatId': this.liveChatId,
                  'type': 'textMessageEvent',
                  'textMessageDetails' : {
                    'messageText': messageText
                  }
                }
              }
            );

        }).then(response => {
  
          console.log(response)
        })
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.chatInput {
  resize: none;
  width: 100%;
  border-radius: 2px;
  border: 1px solid lightgrey;
  padding: 10px;
  margin: 0 15px 0;
  box-sizing: border-box;
}
</style>
