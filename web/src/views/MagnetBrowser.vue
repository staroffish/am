<template>
  <div class="modal-mask" @click.self="$emit('close')">
    <div class="modal-box" style="max-width:720px;max-height:80vh;overflow-y:auto;">
      <h3>从爬虫数据新增规则</h3>

      <label><span>日期</span>
        <select v-model="selectedDate" @change="fetchMagnets">
          <option v-for="d in dates" :key="d" :value="d">{{ d }}</option>
        </select>
      </label>

      <label><span>过滤关键词</span>
        <input v-model="keyword" placeholder="输入关键词过滤..." />
      </label>

      <table v-if="filteredMagnets.length">
        <thead><tr><th>磁链名称</th><th style="width:100px;">操作</th></tr></thead>
        <tbody>
          <tr v-for="m in filteredMagnets" :key="m.name">
            <td style="max-width:480px;overflow:hidden;text-overflow:ellipsis;word-break:break-all;" :title="m.name">{{ m.name }}</td>
            <td>
              <button class="small" :disabled="generating === m.name"
                      @click="generateRule(m)">
                {{ generating === m.name ? '生成中...' : '新增规则' }}
              </button>
            </td>
          </tr>
        </tbody>
      </table>
      <div v-else class="empty" style="padding:24px;">暂无磁链数据</div>

      <div class="flex gap-2 flex-end" style="margin-top:16px;">
        <button class="secondary" @click="$emit('close')">关闭</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

defineEmits(['close'])
const router = useRouter()

interface AnimeMagnet {
  name: string
  magnet_link: string
}

const dates = ref<string[]>([])
const selectedDate = ref('')
const magnets = ref<AnimeMagnet[]>([])
const keyword = ref('')
const generating = ref<string | null>(null)

const filteredMagnets = computed(() => {
  if (!keyword.value) return magnets.value
  const kw = keyword.value.toLowerCase()
  return magnets.value.filter(m => m.name.toLowerCase().includes(kw))
})

function getToday() {
  const d = new Date()
  return d.getFullYear() + '-' +
    String(d.getMonth() + 1).padStart(2, '0') + '-' +
    String(d.getDate()).padStart(2, '0')
}

function generateObjectId() {
  const timestamp = Math.floor(Date.now() / 1000).toString(16).padStart(8, '0')
  const random = Array.from({ length: 16 }, () => Math.floor(Math.random() * 16).toString(16)).join('')
  return `${timestamp}${random}`
}

async function fetchDates() {
  try {
    const res = await axios.get('/api/v1/magnets/dates')
    dates.value = res.data || []
    if (dates.value.length > 0) {
      // Prefer today, fallback to newest
      const today = getToday()
      selectedDate.value = dates.value.includes(today) ? today : dates.value[0]
    } else {
      selectedDate.value = getToday()
    }
  } catch (_) {
    selectedDate.value = getToday()
  }
}

async function fetchMagnets() {
  if (!selectedDate.value) return
  try {
    const res = await axios.get('/api/v1/magnets', { params: { date: selectedDate.value } })
    magnets.value = res.data || []
  } catch (_) {
    magnets.value = []
  }
}

async function generateRule(magnet: AnimeMagnet) {
  generating.value = magnet.name
  try {
    const res = await axios.post('/api/v1/ai/generate-rule', {
      magnet_name: magnet.name,
    })
    const data = res.data
    if (data.is_invalid) {
      alert('AI 生成的正则表达式可能无效，请手动检查修改。\n\n原始响应: ' + (data.raw_response || ''))
    }
    router.push({
      path: '/tasks/new',
      query: {
        name: data.japanese_name || magnet.name,
        regexp: data.regex || '',
        anime_id: generateObjectId(),
      },
    })
  } catch (e: any) {
    const msg = e.response?.data?.error || e.message || '未知错误'
    alert('AI 生成失败: ' + msg + '\n将跳转到手动填写页面。')
    router.push({
      path: '/tasks/new',
      query: {
        name: magnet.name,
        anime_id: generateObjectId(),
      },
    })
  } finally {
    generating.value = null
  }
}

onMounted(async () => {
  await fetchDates()
  await fetchMagnets()
})
</script>
