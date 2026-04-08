<script setup lang="ts">
import { computed, ref } from 'vue'
import type { APIResponse } from '../types'

const props = defineProps<{
  response: APIResponse | null
  loading: boolean
}>()

const copied = ref(false)

const displayText = computed((): string => {
  if (!props.response) return ''
  if (props.response.data !== undefined) {
    return JSON.stringify(props.response.data, null, 2)
  }
  if (props.response.raw) return props.response.raw
  if (props.response.error) return props.response.error
  return JSON.stringify(props.response, null, 2)
})

const statusClass = computed(() => {
  if (!props.response) return ''
  return props.response.success ? 'success' : 'failure'
})

async function copyToClipboard() {
  if (!displayText.value) return
  await navigator.clipboard.writeText(displayText.value)
  copied.value = true
  setTimeout(() => {
    copied.value = false
  }, 1500)
}
</script>

<template>
  <div class="panel">
    <div class="panel-header">
      <span class="title">Ответ</span>
      <span v-if="response" :class="['badge', statusClass]">
        {{ response.success ? 'success' : 'error' }}
      </span>
      <button
        v-if="displayText"
        class="copy-btn"
        @click="copyToClipboard"
        :title="copied ? 'Copied!' : 'Copy to clipboard'"
      >
        {{ copied ? '✓ Copied' : 'Copy' }}
      </button>
    </div>
    <div class="response-body">
      <div v-if="loading" class="loading">
        <span class="spinner" />
        Loading…
      </div>
      <pre v-else-if="displayText" :class="['response-text', statusClass]">{{ displayText }}</pre>
      <span v-else class="placeholder">Response will appear here</span>
    </div>
  </div>
</template>

<style scoped>
.panel {
  display: flex;
  flex-direction: column;
  background: #fff;
  border: 1px solid #d0d7de;
  border-radius: 8px;
  overflow: hidden;
  align-self: start;
  min-height: 480px;
}

.panel-header {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 12px 16px;
  border-bottom: 1px solid #d0d7de;
  background: #f6f8fa;
}

.title {
  font-size: 14px;
  font-weight: 600;
  color: #24292e;
}

.badge {
  font-size: 11px;
  font-weight: 600;
  padding: 2px 8px;
  border-radius: 12px;
  text-transform: uppercase;
  letter-spacing: 0.03em;
}

.badge.success {
  background: #dafbe1;
  color: #1a7f37;
}

.badge.failure {
  background: #ffebe9;
  color: #cf222e;
}

.copy-btn {
  margin-left: auto;
  padding: 4px 10px;
  font-size: 12px;
  font-weight: 500;
  border: 1px solid #d0d7de;
  border-radius: 6px;
  background: #fff;
  cursor: pointer;
  transition: background 0.15s;
}

.copy-btn:hover {
  background: #f3f4f6;
}

.response-body {
  flex: 1;
  padding: 16px;
  overflow-y: auto;
  min-height: 400px;
}

.response-text {
  font-family: 'SFMono-Regular', Consolas, 'Liberation Mono', Menlo, monospace;
  font-size: 13px;
  line-height: 1.6;
  white-space: pre-wrap;
  word-break: break-all;
  margin: 0;
  color: #24292e;
}

.response-text.failure {
  color: #cf222e;
}

.placeholder {
  font-size: 13px;
  color: #8c959f;
}

.loading {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 13px;
  color: #57606a;
}

.spinner {
  display: inline-block;
  width: 14px;
  height: 14px;
  border: 2px solid #d0d7de;
  border-top-color: #25d366;
  border-radius: 50%;
  animation: spin 0.6s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
</style>
