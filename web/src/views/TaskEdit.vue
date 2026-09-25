<template>
  <div>
    <h2>{{ isNew ? '新增规则' : '编辑规则' }}</h2>
    <form @submit.prevent="save">
      <label v-if="isNew"><span>名称(日文)</span><input v-model="form.name" /></label>
      <label><span>正则表达式</span><input v-model="form.regexp" placeholder="【字幕组】\\[...](%02d)..." /></label>
      <label><span>最新集数</span><input v-model.number="form.latest_chapter" type="number" /></label>
      <label v-if="isNew"><span>存储路径</span><input v-model="form.store_path" /></label>
      <label><span>动漫ID</span><input v-model="form.anime_id" /></label>
      <div class="flex gap-2">
        <button type="submit">保存</button>
        <router-link v-if="!isNew && form.anime_id" :to="'/anime/' + form.anime_id" class="btn secondary">动漫详情</router-link>
        <router-link to="/tasks" class="btn secondary">取消</router-link>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import axios from 'axios'

const route = useRoute()
const router = useRouter()
const id = route.params.id as string | undefined
const isNew = computed(() => !id)

const form = ref({
  name: '', regexp: '', latest_chapter: 0, store_path: '', anime_id: '',
})

function getDefaultStorePath() {
  const now = new Date()
  let y = now.getFullYear()
  const m = now.getMonth() + 1
  // 取离当前月份最近的季度月份(1/4/7/10)，12 月归到次年 1 月
  let season: number
  if (m === 12) { y += 1; season = 1 }
  else if (m >= 9) season = 10
  else if (m >= 6) season = 7
  else if (m >= 3) season = 4
  else season = 1
  const prefix = localStorage.getItem('store_dir_prefix') || '/usb/TV/'
  return `${prefix}${y}${String(season).padStart(2,'0')}/`
}

function generateObjectId() {
  const timestamp = Math.floor(Date.now() / 1000).toString(16).padStart(8, '0')
  const random = Array.from({length: 16}, () => Math.floor(Math.random() * 16).toString(16)).join('')
  return `${timestamp}${random}`
}

onMounted(async () => {
  // Pre-fill from query params (from magnet browser AI generation)
  const q = route.query
  if (q.name) form.value.name = q.name as string
  if (q.regexp) form.value.regexp = q.regexp as string
  if (q.anime_id) form.value.anime_id = q.anime_id as string

  try {
    const cfg = await axios.get('/api/v1/config')
    if (cfg.data?.default_store_dir_prefix) {
      localStorage.setItem('store_dir_prefix', cfg.data.default_store_dir_prefix)
    }
  } catch (_) {}

  if (id) {
    const res = await axios.get(`/api/v1/tasks/${id}`)
    const t = res.data
    Object.assign(form.value, {
      regexp: t.regexp || '',
      latest_chapter: t.latest_chapter || 0,
      anime_id: t.anime_id || '',
      name: t.name || '',
      store_path: t.store_path || '',
    })
  } else {
    form.value.store_path = getDefaultStorePath()
    if (!form.value.anime_id) {
      form.value.anime_id = generateObjectId()
    }
  }
})

async function save() {
  const payload: any = {
    regexp: form.value.regexp,
    latest_chapter: form.value.latest_chapter,
    anime_id: form.value.anime_id,
  }
  if (isNew.value) {
    payload.name = form.value.name
    payload.store_path = form.value.store_path
    await axios.post('/api/v1/tasks', payload)
  } else {
    await axios.put(`/api/v1/tasks/${id}`, payload)
  }
  router.push('/tasks')
}
</script>
