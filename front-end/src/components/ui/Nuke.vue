<template>
  <canvas ref="canvas"></canvas>
</template>
<script setup lang="ts">

import {NHEffect} from "@/models/NHEffect";
import {useGame} from "@/store/game";
import {computed, onMounted, ref, watch} from "vue";

const canvas = ref<HTMLCanvasElement>();
let ctx: CanvasRenderingContext2D | null = null;

const effect = computed(() => [...useGame().state.effects].reverse().find((effect: NHEffect) => effect.effectType === 'nuke') as NHEffect | undefined);
class Particle {
  constructor(public x: number, public y: number, public size: number, public xSpeed: number, public ySpeed: number) {}

  // Update particle position
  update() {
    this.x += this.xSpeed;
    this.y += this.ySpeed;
  }
}

let particles: Particle[] = [];

onMounted(() => {
  if (canvas.value) {
    ctx = canvas.value.getContext('2d');
  }
});

watch(
    () => effect.value,
    (newEffect, oldEffect) => {
      if (newEffect?.uuid !== oldEffect?.uuid && ctx) {
        // Reset particles
        particles = [];
        // Draw and animate smoke
        drawSmoke(ctx);
        animate();
      }
    }
);

function drawSmoke(ctx: CanvasRenderingContext2D) {
  // Set the fill color to green
  ctx.fillStyle = 'green';

  // Create 100 smoke particles
  for (let i = 0; i < 100; i++) {
    // Randomly position the particles on the canvas
    const x = canvas.value ? Math.random() * canvas.value.width : 0;
    const y = canvas.value ? Math.random() * canvas.value.height : 0;

    // Randomly size the particles
    const size = Math.random() * 50;

    // Randomly set the particles' speed
    const xSpeed = Math.random() * 2 - 1;
    const ySpeed = Math.random() * 2 - 1;

    // Create the particle
    particles.push(new Particle(x, y, size, xSpeed, ySpeed));
  }
}

function animate() {
  if (ctx && canvas.value) {
    // Clear the canvas
    ctx.clearRect(0, 0, canvas.value.width, canvas.value.height);

    // Update and draw particles
    for (const particle of particles) {
      particle.update();

      // Draw the particle
      ctx.beginPath();
      ctx.arc(particle.x, particle.y, particle.size, 0, Math.PI * 2);
      ctx.closePath();
      ctx.fill();
    }

    // Request next animation frame
    requestAnimationFrame(animate);
  }
}
</script>

<style scoped>
canvas {
  height: 100vh;
  width: 100vw;
  position: fixed;
  z-index: 99999999;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  pointer-events: none;

}
</style>