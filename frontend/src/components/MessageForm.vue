<script setup lang="ts">
import { ref } from 'vue'

defineProps<{ loading: boolean }>()

const emit = defineEmits<{
  sendMessage: [payload: { chatId: string; message: string }]
}>()

const chatId = ref('')
const message = ref('')
const errors = ref({ chatId: '', message: '' })

function submit() {
  errors.value.chatId = chatId.value.trim() ? '' : 'Required'
  errors.value.message = message.value.trim() ? '' : 'Required'
  if (errors.value.chatId || errors.value.message) return
  emit('sendMessage', { chatId: chatId.value.trim(), message: message.value.trim() })
}
</script>

<template>
  <div class="section">
    <div class="field">
      <input
        v-model="chatId"
        type="text"
        placeholder="chatId"
        :class="{ error: errors.chatId }"
        autocomplete="off"
      />
      <span v-if="errors.chatId" class="error-msg">{{ errors.chatId }}</span>
    </div>
    <div class="field">
      <textarea
        v-model="message"
        placeholder="message"
        rows="3"
        :class="{ error: errors.message }"
      />
      <span v-if="errors.message" class="error-msg">{{ errors.message }}</span>
    </div>
    <button @click="submit" :disabled="loading">sendMessage</button>
  </div>
</template>

<style scoped>
.section {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.field {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

input,
textarea {
  width: 100%;
  padding: 8px 10px;
  border: 1px solid #d0d7de;
  border-radius: 6px;
  font-size: 14px;
  outline: none;
  resize: vertical;
  font-family: inherit;
  transition: border-color 0.15s;
}

input:focus,
textarea:focus {
  border-color: #25d366;
}

input.error,
textarea.error {
  border-color: #d1242f;
}

.error-msg {
  font-size: 12px;
  color: #d1242f;
}

button {
  width: 100%;
  padding: 8px 12px;
  border: none;
  border-radius: 6px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  background: #25d366;
  color: #fff;
  transition: background 0.15s;
}

button:hover:not(:disabled) {
  background: #1da851;
}

button:disabled {
  opacity: 0.55;
  cursor: not-allowed;
}
</style>
