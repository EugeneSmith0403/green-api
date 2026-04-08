<script setup lang="ts">
import { ref } from 'vue'

defineProps<{ loading: boolean }>()

const emit = defineEmits<{
  sendFile: [payload: { chatId: string; urlFile: string; fileName?: string; caption?: string }]
}>()

const chatId = ref('')
const urlFile = ref('')
const fileName = ref('')
const caption = ref('')
const errors = ref({ chatId: '', urlFile: '' })

function submit() {
  errors.value.chatId = chatId.value.trim() ? '' : 'Required'
  errors.value.urlFile = urlFile.value.trim() ? '' : 'Required'
  if (errors.value.chatId || errors.value.urlFile) return
  emit('sendFile', {
    chatId: chatId.value.trim(),
    urlFile: urlFile.value.trim(),
    fileName: fileName.value.trim() || undefined,
    caption: caption.value.trim() || undefined,
  })
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
      <input
        v-model="urlFile"
        type="text"
        placeholder="urlFile"
        :class="{ error: errors.urlFile }"
        autocomplete="off"
      />
      <span v-if="errors.urlFile" class="error-msg">{{ errors.urlFile }}</span>
    </div>
    <input v-model="fileName" type="text" placeholder="fileName (optional)" autocomplete="off" />
    <input v-model="caption" type="text" placeholder="caption (optional)" autocomplete="off" />
    <button @click="submit" :disabled="loading">sendFileByUrl</button>
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

input {
  width: 100%;
  padding: 8px 10px;
  border: 1px solid #d0d7de;
  border-radius: 6px;
  font-size: 14px;
  outline: none;
  transition: border-color 0.15s;
}

input:focus {
  border-color: #25d366;
}

input.error {
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
