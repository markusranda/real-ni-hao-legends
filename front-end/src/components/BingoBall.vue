<template>
  <canvas ref="canvas"></canvas>
</template>

<script lang="ts">
export default {
  name: 'BingoBall',
  data() {
    return {
      ball: {
        x: Math.floor(Math.random() * window.innerWidth - 25) + 1,
        y: Math.floor(Math.random() * window.innerHeight - 25) + 1,
        radius: 25,
        speedX: Math.floor(Math.random() * 100) - 50,
        speedY: Math.floor(Math.random() * 100) - 50,
        number: Math.floor(Math.random() * 100) + 1
      },
      animationFrameId: null as null | number
    }
  },
  mounted() {
    this.resizeCanvas()
    window.addEventListener('resize', this.resizeCanvas)
    this.startAnimation()
  },
  beforeUnmount() {
    window.removeEventListener('resize', this.resizeCanvas)
    if (this.animationFrameId) cancelAnimationFrame(this.animationFrameId)
  },
  methods: {
    drawBall() {
      const canvas = this.$refs.canvas as HTMLCanvasElement | undefined
      if (!canvas) throw Error('No canvas, no fun')
      const ctx = canvas.getContext('2d')
      if (!ctx) throw Error('No canvas context, no fun')
      ctx.beginPath()
      ctx.arc(this.ball.x, this.ball.y, this.ball.radius, 0, Math.PI * 2)
      ctx.fillStyle = '#FFF'
      ctx.fill()
      ctx.closePath()

      // Draw the number on the ball
      ctx.font = '14px Arial'
      ctx.fillStyle = 'black'
      ctx.textAlign = 'center'
      ctx.fillText(`${this.ball.number}`, this.ball.x, this.ball.y + this.ball.radius / 4)
    },
    updateBallPosition() {
      const canvas = this.$refs.canvas as HTMLCanvasElement | undefined
      if (!canvas) throw Error('No canvas, no fun')
      this.ball.x += this.ball.speedX
      this.ball.y += this.ball.speedY

      // Bounce off the walls
      if (this.ball.x + this.ball.radius > canvas.width || this.ball.x - this.ball.radius < 0) {
        this.ball.speedX = -this.ball.speedX
      }
      if (this.ball.y + this.ball.radius > canvas.height || this.ball.y - this.ball.radius < 0) {
        this.ball.speedY = -this.ball.speedY
      }

      // Gradually decrease speed
      this.ball.speedX *= 0.99 // Adjust friction here
      this.ball.speedY *= 0.99 // Adjust friction here
    },
    animate() {
      const canvas = this.$refs.canvas as HTMLCanvasElement | undefined
      if (!canvas) throw Error('No canvas, no fun')
      const ctx = canvas.getContext('2d')
      if (!ctx) throw Error('No canvas context, no fun')

      ctx.clearRect(0, 0, canvas.width, canvas.height) // Clear the canvas for the new frame

      this.drawBall()
      this.updateBallPosition()

      // If speed is very low, stop the animation to avoid infinite loop
      if (Math.abs(this.ball.speedX) > 0.1 || Math.abs(this.ball.speedY) > 0.1) {
        this.animationFrameId = requestAnimationFrame(this.animate)
      }
    },
    startAnimation() {
      this.animate()
    },
    resizeCanvas() {
      const canvas = this.$refs.canvas as HTMLCanvasElement | undefined
      if (!canvas) throw Error('No canvas, no fun')

      canvas.width = window.innerWidth
      canvas.height = window.innerHeight

      // Adjust ball position to be within the new canvas dimensions
      this.ball.x = Math.min(this.ball.x, canvas.width - this.ball.radius)
      this.ball.y = Math.min(this.ball.y, canvas.height - this.ball.radius)
    }
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
}
</style>
