<template>
  <div>
    <button @click="show = true">Lars side project</button>

    <Teleport to="body">
      <Modal :show="show" @close="() => (show = false)">
        <template #header>
          <h1>Lars side project</h1>
        </template>
        <template #body>
          <div>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="200"
              height="200"
              fill="none"
              viewBox="0 0 63 58"
              version="1.1"
              xmlns:xlink="http://www.w3.org/1999/xlink"
              xmlns:svgjs="http://svgjs.dev/svgjs"
            >
              <path
                fill="rgba(190, 0, 0, 1)"
                d="M17.3 24.6C29.8 27.8 12 9.2 12 9.2l11.4 9.1c5.7 4.7 4.9 2.5 4.5-3L26.4 3l4.8 12.1c1.3 3.3 3.7 3.7 5-1L41.7.8l-1.5 14.6c-.6 5 .8 3.6 5.6-.9L58.9 3.1 48 16.7c-3.2 4.4-5.5 6.7-1.3 7.1l15.1-.7-13 3.2c-7.4 2.7-2.1 4.3 4.8 7l8.7 3.9-11-2.2c-8.2-2.3-10.4-.8-4.6 5.8l10.1 11.4-11.5-8.8c-2.5-2.3-6.8-4.2-6.6-.3l1.3 14L35.7 43c-1-3.7-2.1-3.8-4 2l-3.8 11.4.8-10.3c.2-6.4-.7-6.7-4.8-3.3L12 51.9l8.8-11.3c2.7-3.2 3.8-5.5.2-5.7L6.3 37.3l10.1-4.2c3.3-1.5 7.5-2.6-.3-5.6L.6 21l16.7 3.6Z"
              ></path>
            </svg>

            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="100"
              height="100"
              fill="none"
              viewBox="0 0 47 44"
              version="1.1"
              xmlns:xlink="http://www.w3.org/1999/xlink"
              xmlns:svgjs="http://svgjs.dev/svgjs"
            >
              <path
                fill="rgba(5, 0, 190, 1)"
                d="M40.1 14.6c-1-.1-2.3-2.2-1.2-3.9.7-1.1 1.4-2.1.3-3.9-.8-1.4-1.7-1.2-2.7-.9-1.6.5-2.2-1.8-2.2-3.7 0-1-.5-1.9-2.3-1.9-1 0-1.5.8-2 1.7-.8 1.3-3.6 1.7-4.1.3-.5-1.4-3-1.5-4.7-.4-1.6 1.1-3.5.7-3.6.4-.2-.2-.3-.4-.4-.4-.1 0-.5-.9-1.5-1.2-.5-.2-1.1-.1-1.9.2-3 1.2-1.7 4-2.7 6.1-1.4 2.7-2.9-.3-5.5.4-1.9.5-3.3 2.8-3.3 5.2.1 2.2 3 4.7 2.7 6.3-.3 1.3-3.4-.6-4.6 3.2-.4 1.2-.2 2.3.1 3.7.6 2.9 2.1 1.7 3.3 2.9 1.4 1.5 1.6 3.1.2 4.9-2.8 3.9.5 6.7 4.3 6.1 2.1-.3 5.5-1.1 4.9.4-.3 1-1.1 2-.6 2.7 2.7 3 4.8-2 7.5-1.8h.3c.1 0 .4.2.7.4.2.2.5.3.6.4.1.1.4 1.1 1.3 1.6.5.3 1.1.5 1.8.5 1.4 0 3.5-.6 4.4-2 .3-.4.3-1.1.3-1.7.1-1 2.1-.2 4.2.1 2.9.4 5.6-1.4 4.3-5.6-.1-.3-.2-.5-.4-.7-.3-.4-.4-1 0-1.4.2-.2.5-.3.9-.2 3.9.8 6.7-.5 5.5-4-.6-1.8-4.5-2.6-1.5-3.6 1.1-.4 2.3-.8 3-2.3 1.7-3.4.5-8.6-3.8-8-.3.1-.9.2-1.6.1Z"
              ></path>
            </svg>
          </div>
        </template>
      </Modal>
    </Teleport>
  </div>
</template>
<script setup lang="ts">
import SambaTimeGame from '@/components/SambaTimeGame.vue'
import { ref } from 'vue'
import Modal from '@/components/Modal.vue'

const show = ref(false)
</script>

<style>
svg {
  position: absolute;
}
path:hover {
  transform: scale(1.1);
  transition: transform 0.2s;
}
</style>
