<script setup lang="ts">
import { ref } from 'vue'

const props = defineProps<{
  idInstance: string
  apiTokenInstance: string
  loading: boolean
}>()

const emit = defineEmits<{
  'update:idInstance': [value: string]
  'update:apiTokenInstance': [value: string]
  getSettings: []
  getStateInstance: []
}>()

const errors = ref({ idInstance: '', apiTokenInstance: '' })

function validate(): boolean {
  errors.value.idInstance = props.idInstance.trim() ? '' : 'Required'
  errors.value.apiTokenInstance = props.apiTokenInstance.trim() ? '' : 'Required'
  return !errors.value.idInstance && !errors.value.apiTokenInstance
}

function onGetSettings() {
  if (validate()) emit('getSettings')
}

function onGetStateInstance() {
  if (validate()) emit('getStateInstance')
}
</script>

<template>
  <div class="section">
    <div class="field">
      <input
        type="text"
        placeholder="idInstance"
        :value="idInstance"
        @input="emit('update:idInstance', ($event.target as HTMLInputElement).value)"
        :class="{ error: errors.idInstance }"
        autocomplete="off"
      />
      <span v-if="errors.idInstance" class="error-msg">{{ errors.idInstance }}</span>
    </div>
    <div class="field">
      <input
        type="text"
        placeholder="apiTokenInstance"
        :value="apiTokenInstance"
        @input="emit('update:apiTokenInstance', ($event.target as HTMLInputElement).value)"
        :class="{ error: errors.apiTokenInstance }"
        autocomplete="off"
      />
      <span v-if="errors.apiTokenInstance" class="error-msg">{{ errors.apiTokenInstance }}</span>
    </div>
    <button @click="onGetSettings" :disabled="loading">getSettings</button>
    <button @click="onGetStateInstance" :disabled="loading">getStateInstance</button>
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
