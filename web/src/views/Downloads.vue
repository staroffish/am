<template>
  <div>
    <h2>下载任务</h2>
    <div class="toolbar">
      <button @click="showAdd = true">+ 添加任务</button>
    </div>

    <div v-if="showAdd" class="modal-mask" @click.self="showAdd = false">
      <div class="modal-box">
        <h3>添加下载任务</h3>
        <form @submit.prevent="addTask">
          <label><span>磁链/URL</span><input v-model="addForm.link" placeholder="magnet:?xt=..." /></label>
          <label><span>存储路径</span><input v-model="addForm.store_path" placeholder="/usb/TV/..." /></label>
          <div class="flex gap-2 flex-end">
            <button type="button" class="secondary" @click="showAdd = false">取消</button>
            <button type="submit">添加</button>
          </div>
        </form>
      </div>
    </div>

    <table v-if="list.length">
      <thead>
        <tr><th>名称</th><th>大小</th><th>进度</th><th>状态</th><th>存储</th><th>创建时间</th><th>操作</th></tr>
      </thead>
      <tbody>
        <tr v-for="t in list" :key="t.hash">
          <td>{{ t.name }}</td>
          <td class="col-narrow">{{ formatSize(t.size) }}</td>
          <td style="min-width:100px;">
            <div class="progress-bar"><div class="progress-bar-fill" :style="{ width: (t.progress * 100).toFixed(1) + '%' }"></div></div>
            <span style="font-size:12px;color:var(--text-dim);">{{ (t.progress * 100).toFixed(1) }}%</span>
          </td>
          <td class="col-narrow"><span class="status-badge" :class="statusClass(t.status)">{{ t.status }}</span></td>
          <td class="col-path">{{ t.store_path }}</td>
          <td class="col-narrow">{{ t.created_time }}</td>
          <td class="col-actions">
            <button v-if="t.status !== 'paused'" class="small secondary" @click="pause(t.hash)">停止</button>
            <button v-if="t.status === 'paused'" class="small secondary" @click="resume(t.hash)">继续</button>
            <button class="small danger" @click="del(t.hash)">删除</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-else class="empty">暂无下载任务</div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import axios from 'axios'

const list = ref([] as any[])
const showAdd = ref(false)
const addForm = ref({ link: '', store_path: '/usb/TV/' })
let timer: any = null

function statusClass(s: string) {
  if (s === 'downloading') return 'status-downloading'
  if (s === 'paused') return 'status-paused'
  if (s === 'seeding' || s === 'finished') return 'status-seeding'
  return 'status-error'
}

function formatSize(s: number) {
  if (s < 1024) return s + ' B'
  if (s < 1048576) return (s / 1024).toFixed(1) + ' KB'
  if (s < 1073741824) return (s / 1048576).toFixed(1) + ' MB'
  return (s / 1073741824).toFixed(1) + ' GB'
}

async function fetch() {
  const res = await axios.get('/api/v1/downloads')
  const data = res.data || []
  data.sort((a: any, b: any) => b.created_time.localeCompare(a.created_time))
  list.value = data
}

async function addTask() {
  await axios.post('/api/v1/downloads', addForm.value)
  addForm.value = { link: '', store_path: '/usb/TV/' }
  showAdd.value = false
  fetch()
}

async function pause(hash: string) { await axios.post(`/api/v1/downloads/${hash}/pause`); fetch() }
async function resume(hash: string) { await axios.post(`/api/v1/downloads/${hash}/resume`); fetch() }
async function del(hash: string) { if (confirm('确定删除？')) { await axios.delete(`/api/v1/downloads/${hash}`); fetch() } }

onMounted(() => { fetch(); timer = setInterval(fetch, 5000) })
onUnmounted(() => { if (timer) clearInterval(timer) })
</script>
