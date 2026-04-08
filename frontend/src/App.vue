<script setup lang="ts">
import { ref, watch } from 'vue'
import ConnectionForm from './components/ConnectionForm.vue'
import MessageForm from './components/MessageForm.vue'
import FileForm from './components/FileForm.vue'
import ResponsePanel from './components/ResponsePanel.vue'
import * as api from './api/greenapi'
import type { APIResponse } from './types'

const idInstance = ref(localStorage.getItem('idInstance') ?? '')
const apiTokenInstance = ref(localStorage.getItem('apiTokenInstance') ?? '')

watch(idInstance, (v) => localStorage.setItem('idInstance', v))
watch(apiTokenInstance, (v) => localStorage.setItem('apiTokenInstance', v))

const loading = ref(false)
const response = ref<APIResponse | null>(null)

async function call(fn: () => Promise<APIResponse>) {
  loading.value = true
  try {
    response.value = await fn()
  } catch (e) {
    response.value = { success: false, action: '', error: String(e) }
  } finally {
    loading.value = false
  }
}

function handleGetSettings() {
  call(() => api.getSettings(idInstance.value, apiTokenInstance.value))
}

function handleGetStateInstance() {
  call(() => api.getStateInstance(idInstance.value, apiTokenInstance.value))
}

function handleSendMessage(payload: { chatId: string; message: string }) {
  call(() =>
    api.sendMessage(idInstance.value, apiTokenInstance.value, payload.chatId, payload.message),
  )
}

function handleSendFile(payload: {
  chatId: string
  urlFile: string
  fileName?: string
  caption?: string
}) {
  call(() =>
    api.sendFileByUrl(
      idInstance.value,
      apiTokenInstance.value,
      payload.chatId,
      payload.urlFile,
      payload.fileName,
      payload.caption,
    ),
  )
}
</script>

<template>
  <div class="app">
    <header class="app-header">
      <span class="logo">GREEN-API</span>
      <span class="subtitle">WhatsApp Integration Demo</span>
    </header>
    <main class="layout">
      <div class="left-panel">
        <div class="card">
          <ConnectionForm
            v-model:id-instance="idInstance"
            v-model:api-token-instance="apiTokenInstance"
            :loading="loading"
            @get-settings="handleGetSettings"
            @get-state-instance="handleGetStateInstance"
          />
        </div>
        <div class="card">
          <MessageForm :loading="loading" @send-message="handleSendMessage" />
        </div>
        <div class="card">
          <FileForm :loading="loading" @send-file="handleSendFile" />
        </div>
      </div>
      <ResponsePanel :response="response" :loading="loading" />
    </main>
  </div>
</template>

<style>
*,
*::before,
*::after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  background: #f6f8fa;
  color: #24292e;
  min-height: 100vh;
}
</style>

<style scoped>
.app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.app-header {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 14px 24px;
  background: #fff;
  border-bottom: 1px solid #d0d7de;
}

.logo {
  font-size: 16px;
  font-weight: 700;
  color: #25d366;
  letter-spacing: 0.02em;
}

.subtitle {
  font-size: 13px;
  color: #57606a;
}

.layout {
  display: grid;
  grid-template-columns: 320px 1fr;
  gap: 24px;
  padding: 24px;
  flex: 1;
  align-items: start;
}

.left-panel {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.card {
  background: #fff;
  border: 1px solid #d0d7de;
  border-radius: 8px;
  padding: 16px;
}

@media (max-width: 720px) {
  .layout {
    grid-template-columns: 1fr;
  }
}
</style>
