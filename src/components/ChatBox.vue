<template>
  <div>
    <div id="chatBox" class="flex">
      <ChatItem v-for="(comment,index) in comments" :key="index" :message="comment.message" :author="comment.author" />
      
    </div>
    <ChatInput :live-chat-id="liveChatId" />
  </div>
</template>

<script>
import ChatItem from './ChatItem.vue'
import ChatInput from './ChatInput.vue'

export default {
  name: "ChatBox",
  data() {
    return {
      isAtBottom: true
    }
  },
  props: {
    comments: Array,
    liveChatId: String,
  },
  components: {
    ChatItem,
    ChatInput
  },
  beforeUpdate() {
    this.isAtBottom = this.atBottom()
  },
  updated() {
    if(this.isAtBottom) {
      this.scrollToBottom();
    }
  },
  methods: {
    atBottom() {
      let chatBox = document.getElementById("chatBox")
      return (chatBox.scrollTop + chatBox.offsetHeight) === chatBox.scrollHeight;
    },
    scrollToBottom() {
      let chatBox = document.getElementById("chatBox")
      chatBox.scrollTop = chatBox.scrollHeight;
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.flex {
  width: 400px;
  height: 500px;
  display: flex;
  justify-content: flex-start;
  flex-wrap: wrap;
  margin: 15px;
  background-color: lightgray;
  overflow: auto;
}
</style>
