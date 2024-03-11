<template>
  <div>
    <button @click="handleClickPlay">Play databingo</button>

    <Teleport to="body">
      <div v-if="showDatabingo">
        <DatabingoAgeConfirm
          v-if="!isAgeConfirmed"
          :handle-age-confirmation="handleAgeConfirmation"
        />
        <DatabingoGame v-else :handle-click-close="handleClickClose" />
      </div>
    </Teleport>
  </div>
</template>

<script lang="ts">
import DatabingoAgeConfirm from './DatabingoAgeConfirm.vue'
import DatabingoGame from './DatabingoGame.vue'

export default {
  components: {
    DatabingoAgeConfirm,
    DatabingoGame
  },
  data() {
    return {
      showDatabingo: false,
      isModalOpen: false,
      isAgeConfirmed: false
    }
  },
  methods: {
    handleClickPlay() {
      this.showDatabingo = true
    },
    handleClickClose() {
      this.showDatabingo = false
    },
    handleAgeConfirmation() {
      this.isAgeConfirmed = true
    }
  }
}
</script>
