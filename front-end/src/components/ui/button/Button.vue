<script setup lang="ts">
const audio = new Audio("https://www.zapsplat.com/wp-content/uploads/2015/sound-effects-46416/zapsplat_multimedia_button_click_005_53866.mp3");
</script>

<template>
  <button
      @click="() => audio.play()"
      class="cool-button"
  >
    <slot/>
  </button>
</template>

<style scoped lang="scss">
.cool-button {
  background: linear-gradient(116deg, rgba(181, 181, 232, 0.4), rgba(241, 173, 255, 0.4));
  border: 0;
  border-radius: 24px;
  user-select: none;
  color: #313131;
  cursor: pointer;
  transition-property: box-shadow;
  -webkit-tap-highlight-color: transparent;
  font-weight: 500;

  box-shadow: 0 0 0 1px white,
  0em 0.1em 0.3em rgba(0, 0, 0, 0.3);
  font-size: 12px;
  padding: 12px 24px;

  &, span {
    display: block;
    transition: {
      duration: 0.1s;
      timing-function: linear;
    }

    &:focus {
      outline: none;
    }
  }

  span {
    transition-property: transform;
    will-change: transform;
  }

  transition: transform 0.4s ease-in-out;

  &:hover {
    transform: scale3d(1.05, 1.10, 1.05);
  }

  &:active {
    transform: scale3d(0.85, 0.95, 0.85);
  }
}

</style>
