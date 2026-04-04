<script setup>
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { EventsOn } from '../wailsjs/runtime/runtime'
import { getStudyCard, loadDashboard, submitAnswer } from './api'

const dashboard = ref(null)
const selectedLanguage = ref('zh-TW')
const selectedAnswer = ref('')
const feedback = ref(null)
const loading = ref(true)
const submitting = ref(false)
let unsubscribe = null

const card = computed(() => dashboard.value?.currentCard ?? null)
const stats = computed(() => dashboard.value?.stats ?? { studiedToday: 0, correctRate: 0 })
const reviewMode = computed(() => dashboard.value?.reviewMode ?? false)
const reviewQueue = computed(() => dashboard.value?.reviewQueue ?? [])
const summary = computed(() => dashboard.value?.summary ?? { nextReviewAt: '', weakTopics: [] })
const explanation = computed(() => {
  if (!card.value) {
    return ''
  }

  return selectedLanguage.value === 'en' ? card.value.explanationEn : card.value.explanationZh
})

const formattedCorrectRate = computed(() => `${Math.round((stats.value.correctRate ?? 0) * 100)}%`)
const nextReviewText = computed(() => summary.value.nextReviewAt || 'Not scheduled')

onMounted(async () => {
  dashboard.value = await loadDashboard()
  selectedLanguage.value = dashboard.value.preferredLanguage || dashboard.value.info.defaultLanguage
  loading.value = false

  if (window?.runtime) {
    EventsOn('notification:open-card', async (cardId) => {
      const nextCard = await getStudyCard(cardId)
      dashboard.value = {
        ...dashboard.value,
        currentCard: nextCard,
      }
      feedback.value = null
      selectedAnswer.value = ''
    })
    unsubscribe = true
  }
})

onUnmounted(() => {
  unsubscribe = null
})

async function handleSubmit() {
  if (!card.value || !selectedAnswer.value) {
    return
  }

  submitting.value = true
  feedback.value = await submitAnswer({
    cardId: card.value.id,
    sessionType: 'learn',
    selectedAnswer: selectedAnswer.value,
    shownAt: card.value.shownAt,
  })
  dashboard.value = {
    ...dashboard.value,
    stats: feedback.value.stats,
  }
  submitting.value = false
}
</script>

<template>
  <main class="shell">
    <section class="hero">
      <p class="eyebrow">duolin-gogo</p>
      <h1>Bilingual Git micro-learning</h1>
      <p class="summary">
        Local-first flashcards built from bilingual Markdown notes, with Windows notifications,
        lightweight review sessions, and adaptive repetition.
      </p>
    </section>

    <section class="status-grid">
      <article class="status-card">
        <span class="label">Current focus</span>
        <strong>Git knowledge only</strong>
      </article>
      <article class="status-card">
        <span class="label">Language mode</span>
        <strong>{{ selectedLanguage }} toggle ready</strong>
      </article>
      <article class="status-card">
        <span class="label">Development mode</span>
        <strong>TDD mode</strong>
      </article>
      <article class="status-card">
        <span class="label">Session mode</span>
        <strong>{{ reviewMode ? 'Review session' : 'Learn session' }}</strong>
      </article>
    </section>

    <section v-if="loading" class="study-card">
      <p class="label">Loading</p>
      <strong>Preparing the next Git card...</strong>
    </section>

    <section v-else-if="card" class="study-card">
      <div class="study-header">
        <div>
          <p class="label">{{ reviewMode ? 'Review card' : 'Next card' }}</p>
          <h2>{{ card.title }}</h2>
        </div>
        <div class="language-toggle">
          <button
            :class="{ active: selectedLanguage === 'zh-TW' }"
            type="button"
            @click="selectedLanguage = 'zh-TW'"
          >
            zh-TW
          </button>
          <button
            :class="{ active: selectedLanguage === 'en' }"
            type="button"
            @click="selectedLanguage = 'en'"
          >
            en
          </button>
        </div>
      </div>

      <p v-if="reviewMode" class="review-banner">
        Review session active. Queue size: {{ reviewQueue.length }}
      </p>
      <p class="callout">{{ card.clickbait }}</p>
      <p class="explanation">{{ explanation }}</p>

      <div class="question-block">
        <p class="label">Quick question</p>
        <h3>{{ card.questionText }}</h3>

        <div class="answers">
          <label v-for="choice in card.choices" :key="choice.value" class="answer-option">
            <input v-model="selectedAnswer" type="radio" name="answer" :value="choice.value">
            <span>{{ choice.label }}</span>
          </label>
        </div>

        <button class="submit-button" type="button" :disabled="!selectedAnswer || submitting" @click="handleSubmit">
          {{ submitting ? 'Checking...' : 'Submit answer' }}
        </button>
      </div>

      <div v-if="feedback" class="feedback" :class="{ correct: feedback.isCorrect, wrong: !feedback.isCorrect }">
        <strong>{{ feedback.feedback }}</strong>
        <p>Correct answer: {{ feedback.correctAnswer }}</p>
        <p>{{ feedback.reviewHint }}</p>
      </div>
    </section>

    <section class="status-grid">
      <article class="status-card">
        <span class="label">Studied today</span>
        <strong>{{ stats.studiedToday }}</strong>
      </article>
      <article class="status-card">
        <span class="label">Correct rate</span>
        <strong>{{ formattedCorrectRate }}</strong>
      </article>
      <article class="status-card">
        <span class="label">Review queue</span>
        <strong>{{ reviewQueue.length }}</strong>
      </article>
      <article class="status-card">
        <span class="label">Next review</span>
        <strong>{{ nextReviewText }}</strong>
      </article>
    </section>

    <section class="study-card">
      <div class="study-header">
        <div>
          <p class="label">Weak topics</p>
          <h2>Git concepts to revisit</h2>
        </div>
      </div>

      <div v-if="summary.weakTopics.length" class="weak-topics">
        <article v-for="topic in summary.weakTopics" :key="topic.tag" class="topic-chip">
          <strong>{{ topic.tag }}</strong>
          <span>{{ Math.round(topic.accuracy * 100) }}% accuracy</span>
        </article>
      </div>
      <p v-else class="explanation">No weak topics yet. Keep studying to generate insights.</p>
    </section>
  </main>
</template>
