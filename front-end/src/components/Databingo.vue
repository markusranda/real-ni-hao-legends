<template>
  <div>
    <Button @click="handleClickPlay">Databingo</Button>

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
import Button from "@/components/ui/button/Button.vue";

export default {
  components: {
    Button,
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
