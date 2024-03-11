<script lang="ts" setup>
import { computed, onMounted, ref, watch } from 'vue'
import { useGame } from '@/store/game'
import { NHCommand } from '@/models/nh-command'
import { useWebsocket } from '@/store/websocketStore'

const chat = computed(() => useGame().chat)
const currentMessage = ref('')
const chatMessagesContainer = ref<HTMLDivElement>() // Define the ref here

function sendChatMessage() {
  if (currentMessage.value) {
    useWebsocket().send({
      type: 'chat',
      userId: localStorage.getItem('UserId')!,
      data: { message: currentMessage.value, type: '' }
    } as NHCommand)
    currentMessage.value = ''
  }
}

watch(chat, () => {
  chatMessagesContainer.value?.scrollTo({
    top: chatMessagesContainer.value?.scrollHeight,
    behavior: 'smooth'
  })
})

onMounted(() => {
  chatMessagesContainer.value?.scrollTo({
    top: chatMessagesContainer.value?.scrollHeight,
    behavior: 'instant'
  })
})
</script>

<template>
  <div class="the-game-chat nihao-box">
    <div class="the-game-chat-inner">
      <div class="the-game-chat-messages" ref="chatMessagesContainer">
        <div
          class="the-game-chat-message"
          v-for="(chatMessage, index) in chat"
          :key="`msg__${index}`"
        >
          <span>{{ chatMessage.userName }}: {{ chatMessage.message }}</span>
        </div>
      </div>
      <div class="the-game-chat-input">
        <input
          type="text"
          placeholder="message.."
          v-model="currentMessage"
          @keydown.enter="sendChatMessage"
          class="col-2"
        />
        <button @click="sendChatMessage">send</button>
      </div>
    </div>
  </div>
</template>
