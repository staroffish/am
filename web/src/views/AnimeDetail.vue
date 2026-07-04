<template>
  <div>
    <div v-if="anime">
      <div class="toolbar">
        <button class="btn secondary" @click="$router.back()">← 返回</button>
        <router-link :to="'/anime/' + anime.id + '/edit'" class="btn secondary">编辑</router-link>
      </div>
      <div class="card">
        <div class="flex gap-2" style="align-items:flex-start;">
          <img v-if="animeHasImage" :src="imageSrc" style="max-width:200px;max-height:280px;border-radius:var(--radius);" />
          <div style="flex:1">
            <h2>{{ anime.animenamejp }} <small style="color:var(--text-dim);font-size:0.8em;">{{ anime.animenamecn }}</small></h2>
            <p><strong>ID:</strong> {{ anime.id }}</p>
            <p><strong>类型:</strong> {{ anime.type || '-' }}</p>
            <p><strong>状态:</strong> {{ anime.status || '-' }}</p>
            <p><strong>CAST:</strong> {{ anime.cast || '-' }}</p>
            <p><strong>连载期间:</strong> {{ anime.serialsduri || '-' }}</p>
            <p><strong>存储路径:</strong> {{ anime.stordir }}</p>
            <p><strong>更新时间:</strong> {{ fmtTime(anime.updated_at) }}</p>
          </div>
        </div>
      </div>
      <div v-if="files.length" class="mt-2">
        <h3>文件列表 ({{ files.length }})</h3>
        <div class="card">
          <div v-for="f in files" :key="f.name" style="padding: 6px 0;border-bottom:1px solid var(--border);word-break:break-all;">
            {{ f.name }}
          </div>
        </div>
      </div>
      <div v-else class="empty">存储目录为空或无法访问</div>
    </div>
    <div v-else class="empty">加载中...</div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import axios from 'axios'

const route = useRoute()
const anime = ref(null as any)
const files = ref([] as any[])
const imageSrc = ref('')

const animeHasImage = computed(() => !!imageSrc.value)

function fmtTime(t: string) {
  if (!t || t.startsWith('0001')) return ''
  const d = new Date(t)
  return d.toLocaleString('zh-CN', { hour12: false })
}

onMounted(async () => {
  const res = await axios.get(`/api/v1/anime/${route.params.id}`)
  anime.value = res.data.anime
  files.value = res.data.files || []

  if (res.data.has_image) {
    const imgResp = await axios.get(`/api/v1/anime/${route.params.id}/image`, { responseType: 'arraybuffer' })
    if (imgResp.data && imgResp.data.byteLength > 100) {
      const blob = new Blob([imgResp.data], { type: imgResp.headers['content-type'] || 'image/jpeg' })
      imageSrc.value = URL.createObjectURL(blob)
    }
  }
})
</script>
