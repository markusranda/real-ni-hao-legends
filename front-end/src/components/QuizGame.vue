<template>
  <Modal
    :show="true"
    background-color="background-color: rgba(0, 0, 0, 0.9)"
    modalBodyClass="quiz-body"
    modalContainerClass="quiz-container"
    modalHeaderClass="quiz-header"
    @close="() => handleClickClose()"
  >
    <template #header>
      <h1></h1>
      <h1>KNAUSER QUIZ</h1>
      <Button style="margin-left: auto; width: min-content" @click="() => handleClickClose()">
        Pussy out
      </Button>
    </template>

    <template #body>
      <h1>{{ question.question }}</h1>
      <div class="question-container">
        <Button
          v-for="answer of question.answers"
          :key="`question__${answer.answer}`"
          @click="() => handleClickAnswer(answer)"
        >
          {{ answer.answer }}
        </Button>
      </div>
    </template>

    <template #footer>
      <span></span>
    </template>
  </Modal>

  <audio id="audio-suspense" src="suspense.mp3"></audio>
  <audio id="audio-clap" src="clap.mp3"></audio>
  <audio id="audio-gun-shot" src="gun_shot_2.wav"></audio>
</template>

<script lang="ts">
import { QUESTIONS, question_1 } from '@/plugins/quiz'
import Modal from './Modal.vue'
import Button from './ui/button/Button.vue'
import { NHQuizAnswer } from '@/models/nh-quiz'

export default {
  components: { Modal, Button },
  props: {
    handleClickClose: {
      type: Function,
      required: true
    }
  },
  data() {
    return {
      question: question_1,
      questions: QUESTIONS,
      volume: 0.25
    }
  },
  mounted() {
    this.playSuspenseSound()
  },
  methods: {
    handleClickAnswer(answer: NHQuizAnswer) {
      if (answer.finished) {
        alert('You are the winner of this chicken dinner!')
        return
      }

      const questionStr = answer.question ?? ''

      if (questionStr === 'question_1') {
        this.playBadSound()
      }

      const newQuestion = this.questions[questionStr]
      if (newQuestion) {
        this.playClapSound()
        this.question = newQuestion
      }
    },
    playBadSound() {
      const audio = document.getElementById('audio-gun-shot') as HTMLAudioElement
      audio.volume = this.volume
      audio.play()
    },
    playSuspenseSound() {
      const audio = document.getElementById('audio-suspense') as HTMLAudioElement
      audio.volume = this.volume
      audio.loop = true
      audio.play()
    },
    playClapSound() {
      const audio = document.getElementById('audio-clap') as HTMLAudioElement
      audio.volume = this.volume
      audio.play()
    }
  }
}
</script>
<style lang="scss">
.quiz-header {
  width: 100%;
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  text-align: center;
  color: #fff;

  text-shadow:
    -2px -2px 0 black,
    2px -2px 0 black,
    -2px 2px 0 black,
    2px 2px 0 black;
}

.quiz-container {
  display: grid;
  grid-template-rows: auto minmax(1rem, 1fr) auto;
  width: 100%;

  background: radial-gradient(
    circle,
    rgba(0, 212, 255, 1) 0%,
    rgba(9, 9, 121, 1) 22%,
    rgba(2, 0, 36, 1) 47%
  );

  animation: moveGradient 10s ease infinite;
  margin: 2rem;
  padding: 2rem;
}

.quiz-body {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 2rem;

  padding: 2rem;
  color: #fff;
  text-shadow:
    -2px -2px 0 black,
    2px -2px 0 black,
    -2px 2px 0 black,
    2px 2px 0 black;

  width: 100%;
  height: 100%;
}

.question-container {
  display: flex;
  gap: 1rem;
}
</style>
