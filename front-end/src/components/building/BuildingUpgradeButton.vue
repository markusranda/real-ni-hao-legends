<script setup lang="ts">

const audio = new Audio("https://www.zapsplat.com/wp-content/uploads/2015/sound-effects-46416/zapsplat_multimedia_button_click_005_53866.mp3");

const props = defineProps<{
  canUpgrade: boolean |undefined,
}>()
</script>

<template>
  <button :data-can-upgrade="props.canUpgrade" @click="audio.play()">
    <slot/>
  </button>
</template>

<style scoped lang="scss">

/* Text */
span {
  background: linear-gradient(180deg, #F3F3F3 0%, #A5A5A5 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  text-fill-color: transparent;
  text-shadow: 0px 2px 4px rgba(0, 0, 0, 0.16);
  letter-spacing: 0.5px;
  font-size: 18px;
}

/* Button */
button {
  width: 100%;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  padding: 18px 22px;
  background: var(--gradient);
  border: 0;
  position: relative;
  border-radius: var(--border-radius);
  z-index: 1;
  text-decoration: none;
  cursor: pointer;
  border-radius: 0 0 24px 24px;
}

/* Light notch */
button::after {
  content: '';
  display: block;
  width: 40px;
  height: 1px;
  position: absolute;

  top: 0px;
  box-shadow: 0px 1px 1px rgba(255, 255, 255, .1);
  transition: width 1s cubic-bezier(0.19, 1, 0.22, 1);
}

button {
  &[data-can-upgrade="true"]::after {
    background: #2A2C34;
    border: .5px solid #34FFAA;
    box-shadow: 0px -22px 26px rgba(0, 255, 163, 0.43), 0px 22px 26px rgba(75, 218, 166, 0.4), 0px -22px 20px rgba(89, 255, 195, 0.1), 0px 0px 16px #59FFC3;
  }

  &[data-can-upgrade="false"]::after {
    background: #FFEBEB;
    border: .5px solid #FF3434;
    box-shadow: 0px -22px 26px rgba(255, 0, 0, 0.43), 0px 22px 26px rgba(218, 75, 75, 0.4), 0px -22px 20px rgba(255, 89, 89, 0.1), 0px 0px 16px #FF5959;
  }
}

/* Light notch on button click */
button:active::after {
  width: 140px; /* Adjust this value to your desired width */
}

/* Button on click */
button:active {
  color: #EAFFFC;
  background: rgb(238, 174, 202);
  background: linear-gradient(190deg, rgba(238, 174, 202, 1) 0%, rgba(148, 187, 233, 0.6719999999999999) 100%);
}

/* Light notch on button click */
button:active::after {
  [data-can-upgrade="true"] {
    background: #EBFDFF;
    border: .5px solid #34FFAA;
    box-shadow: 0px -22px 26px rgba(0, 255, 163, 0.43), 0px 22px 26px rgba(75, 218, 166, 0.4), 0px -22px 20px rgba(89, 255, 195, 0.1), 0px 0px 16px #59FFC3;
  }
}

/* Light notch on button hover */
button:hover::after {
  [data-can-upgrade="true"] {
    background: #EBFDFF;
    border: .5px solid #34FFAA;
    box-shadow: 0px -22px 26px rgba(0, 255, 163, 0.43), 0px 22px 26px rgba(75, 218, 166, 0.4), 0px -22px 20px rgba(89, 255, 195, 0.1), 0px 0px 16px #59FFC3;
  }
}


</style>