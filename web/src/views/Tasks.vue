<template>
  <div>
    <h2>下载规则</h2>
    <div class="toolbar">
      <router-link to="/tasks/new" class="btn">+ 新增规则</router-link>
      <button @click="scanDownload">扫描并下载</button>
      <button class="small secondary" @click="refresh" style="border-radius:50%;width:28px;height:28px;padding:0;" title="刷新">↻</button>
    </div>
    <div class="toolbar">
      <span style="font-size:13px;color:var(--text-dim);">手动触发爬虫:</span>
      <button v-for="s in spiders" :key="s.name" class="small secondary" @click="crawl(s.name)" :disabled="s.running">
        {{ s.running ? '⏳' : '' }} {{ s.name }}
      </button>
    </div>
    <table v-if="tasks.length">
      <thead>
        <tr><th>名称</th><th>最新集数</th><th>存储路径</th><th>动漫ID</th><th>更新时间</th><th>操作</th></tr>
      </thead>
      <tbody>
        <tr v-for="t in tasks" :key="t.id">
          <td style="max-width:240px;overflow:hidden;text-overflow:ellipsis;">{{ t.name }}</td>
          <td class="col-narrow" style="min-width:80px;">{{ t.latest_chapter }}</td>
          <td class="col-path" style="max-width:180px;">{{ t.store_path }}</td>
          <td class="col-narrow" style="max-width:80px;overflow:hidden;text-overflow:ellipsis;" :title="t.anime_id">{{ t.anime_id }}</td>
          <td class="col-narrow">{{ fmtTime(t.updated_at) }}</td>
          <td class="col-actions">
            <router-link :to="'/tasks/' + t.id + '/edit'" class="btn small secondary">编辑</router-link>
            <button class="small danger" @click="del(t.id)">删除</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-else class="empty">暂无下载规则</div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'

const tasks = ref([] as any[])
const spiders = ref([] as any[])

function fmtTime(t: string) {
  if (!t || t.startsWith('0001')) return ''
  const d = new Date(t)
  return d.toLocaleString('zh-CN', { hour12: false })
}

async function fetchTasks() {
  const res = await axios.get('/api/v1/tasks')
  tasks.value = res.data || []
}

async function fetchSpiders() {
  const res = await axios.get('/api/v1/spiders')
  spiders.value = res.data || []
}

async function crawl(name: string) {
  try {
    await axios.post(`/api/v1/spiders/${name}/crawl`)
    alert(`爬虫 ${name} 已触发`)
  } catch (e: any) {
    alert(`爬虫 ${name} 失败: ${e.response?.data?.error || e.message}`)
  }
}

async function scanDownload() {
  const res = await axios.post('/api/v1/tasks/scan-and-download')
  alert(`扫描并下载完成，创建 ${res.data.length} 个任务`)
}

async function refresh() { fetchTasks(); fetchSpiders() }

async function del(id: number) {
  if (!confirm('确定删除？')) return
  await axios.delete(`/api/v1/tasks/${id}`)
  fetchTasks()
}

onMounted(() => { fetchTasks(); fetchSpiders() })
</script>
