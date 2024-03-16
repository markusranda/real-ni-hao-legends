<template>
  <Modal :show="true">
    <template #header>
      <h3>Age Verification Required</h3>
    </template>
    <template #body>
      <p>‼️‼️GAMBLING IS NOT FOR KIDS ‼️‼️ PLEASE CONFIRM THAT YOU ARE 18 YEARS OR OLDER</p>
    </template>
    <template #footer>
      <Button @click="() => handleAgeConfirmation()">Okay</Button>
    </template>
  </Modal>
</template>

<script lang="ts">
import Modal from './Modal.vue'

export default {
  props: {
    handleAgeConfirmation: {
      type: Function,
      required: true
    }
  },
  components: { Modal }
}
</script>

<style scoped></style>
