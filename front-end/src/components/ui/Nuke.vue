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
  constructor(public x: number, public y: number, public size: number, public xSpeed: number, public ySpeed: number, public color: string) {}

  // Update particle position
  update() {
    this.x += this.xSpeed;
    this.y += this.ySpeed;

    // If the particle is outside the canvas, reset its position
    if (this.x < 0 || this.x > canvas.value?.width || this.y < 0 || this.y > canvas.value?.height) {
      this.x = Math.random() * (canvas.value?.width || 0);
      this.y = Math.random() * (canvas.value?.height || 0);
    }
  }

  // Decrease the size of the particle over time
  fade() {
    this.size -= 0.002;
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
        if (newEffect) {
          drawSmoke(ctx);
          animate();
        }
      }
    }
);

function getRandomGreen() {
  const green = Math.floor(Math.random() * 256);
  const alpha = Math.random(); // Generate a random alpha value between 0 and 1
  return `rgba(0, ${green}, 0, ${alpha})`;
}

function drawSmoke(ctx: CanvasRenderingContext2D) {
  // Create 500 smoke particles
  for (let i = 0; i < 500; i++) { // Increase the number of particles
    // Randomly position the particles on the canvas
    const x = canvas.value ? Math.random() * canvas.value.width : 0;
    const y = canvas.value ? Math.random() * canvas.value.height : 0;

    // Randomly size the particles
    const size = Math.random() * 50;

    // Randomly set the particles' speed
    const xSpeed = Math.random() * 2 - 1;
    const ySpeed = Math.random() * 2 - 1;

    // Create the particle with a random green color
    particles.push(new Particle(x, y, size, xSpeed, ySpeed, getRandomGreen()));
  }
}

function animate() {
  if (ctx && canvas.value) {
    // Clear the canvas
    ctx.clearRect(0, 0, canvas.value.width, canvas.value.height);

    // Update and draw particles
    for (const particle of particles) {
      particle.update();
      particle.fade();

      // Only draw the particle if its size is greater than 0
      if (particle.size > 0) {
        // Set the fill color to the particle's color
        ctx.fillStyle = particle.color;

        // Draw the particle
        ctx.beginPath();
        ctx.arc(particle.x, particle.y, particle.size, 0, Math.PI * 2);
        ctx.closePath();
        ctx.fill();
      }
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