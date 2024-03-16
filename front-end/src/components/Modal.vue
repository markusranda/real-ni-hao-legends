<template>
  <Transition name="modal">
    <div v-if="show" class="modal-mask" :style="backgroundColor">
      <div :class="modalContainerClass">
        <div :class="modalHeaderClass">
          <slot name="header">default header</slot>
        </div>

        <div :class="modalBodyClass">
          <slot name="body">default body</slot>
        </div>

        <div class="modal-footer">
          <slot name="footer">
            <Button class="modal-default-button" @click="$emit('close')">OK</Button>
          </slot>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script lang="ts">
export default {
  props: {
    show: {
      type: Boolean,
      required: true
    },
    backgroundColor: {
      type: String,
      default: '  background-color: rgba(0, 0, 0, 0.5);'
    },
    modalContainerClass: {
      type: String,
      default: 'modal-container'
    },
    modalHeaderClass: {
      type: String,
      default: 'modal-header'
    },
    modalBodyClass: {
      type: String,
      defaul: 'modal-body'
    }
  },
  emits: ['close']
}
</script>

<style>
.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  transition: opacity 0.3s ease;
}

.modal-header h3 {
  margin-top: 0;
  color: #42b983;
}

.modal-default-button {
  float: right;
}

/*
 * The following styles are auto-applied to elements with
 * transition="modal" when their visibility is toggled
 * by Vue.js.
 *
 * You can easily play with the modal transition by editing
 * these styles.
 */

.modal-enter-from {
  opacity: 0;
}

.modal-leave-to {
  opacity: 0;
}

.modal-enter-from .modal-container,
.modal-leave-to .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
</style>
