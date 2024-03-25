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
      <span></span>
      <h1>KNAUSER QUIZ</h1>
      <QuizButton
        style="margin-left: auto; width: min-content"
        text="Pussy out"
        @click="() => handleClickClose()"
      >
      </QuizButton>
    </template>

    <template #body>
      <div class="question-text">{{ question.question }}</div>
      <div class="question-container">
        <QuizButton
          v-for="(answer, index) of question.answers"
          :key="`question__${answer.answer}`"
          :index="index"
          :text="answer.answer"
          @click="() => handleClickAnswer(answer)"
        >
        </QuizButton>
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
import QuizButton from './QuizButton.vue'
import { NHQuizAnswer } from '@/models/nh-quiz'

export default {
  // eslint-disable-next-line vue/no-reserved-component-names
  components: { Modal, QuizButton },
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

  align-items: center;

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

.question-text {
  color: #fff;
  background-color: #0e0e0e;
  border: solid 2px #6097c4;
  border-radius: 2rem;

  padding: 2rem 10rem 2rem 10rem;

  font-size: 2rem;
}

.question-container {
  display: flex;
  gap: 1rem;

  strong {
    color: #e0aa49;
  }
}
</style>
