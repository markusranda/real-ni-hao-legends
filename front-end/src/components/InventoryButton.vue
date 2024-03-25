<template>
  <button class="inventory-button" :class="buttonClass" @click="handleClick">
    {{ buttonTitle }}
  </button>
</template>

<script lang="ts">
import { GameState } from '@/models/nh-gamestate'
import { useGame } from '@/store/game'
import clamp from 'lodash.clamp'

export default {
  props: {
    isInventory: {
      type: Boolean,
      required: true
    }
  },
  data() {
    return {
      newInventoryItems: 1
    }
  },
  computed: {
    state() {
      return useGame().state
    },
    buttonTitle() {
      if (this.newInventoryItems > 0) return `Inventory (${this.newInventoryItems})`
      return 'Inventory'
    },
    buttonClass() {
      return this.newInventoryItems > 0 ? 'glow-button' : ''
    }
  },
  watch: {
    state: {
      deep: true,
      handler(newState: GameState, oldState: GameState) {
        const diff = newState.inventory.buildings.length - oldState.inventory.buildings.length

        const newInventoryItems = clamp(diff, 0, 99)
        if (newInventoryItems > 0) {
          this.newInventoryItems += newInventoryItems
        }
      }
    }
  },
  methods: {
    handleClick() {
      this.$emit('update:isInventory', !this.isInventory)
      this.newInventoryItems = 0
    }
  }
}
</script>
<style lang="scss">
.inventory-button {
  position: absolute;
  bottom: 0;
  right: 0;
  z-index: 1;

  font-size: 1rem;
  margin: 1rem;
  padding: 1rem 1rem;

  border: solid 2px rgba(0, 0, 0, 0);
  border-radius: 5px;

  color: #3d3d3d;
  background: linear-gradient(116deg, rgb(181, 181, 232), rgb(241, 173, 255));
}

@keyframes glow-animation {
  0% {
    border-color: #fff;
    font-size: 1rem;
  }
  100% {
    border-color: linear-gradient(116deg, rgb(181, 181, 232), rgb(241, 173, 255));
    font-size: 1.05rem;
  }
}

.glow-button {
  animation: glow-animation 750ms ease-in-out infinite alternate;
}
</style>
