<template>
  <div>
    <h2>{{ isNew ? '新增动漫' : '编辑动漫' }}</h2>
    <form @submit.prevent="save">
      <label><span>中文名</span><input v-model="form.animenamecn" /></label>
      <label><span>日文名</span><input v-model="form.animenamejp" /></label>
      <label><span>CAST</span><input v-model="form.cast" /></label>
      <label><span>类型</span><input v-model="form.type" /></label>
      <label><span>状态</span>
        <select v-model="form.status">
          <option>连载中</option>
          <option>已完结</option>
        </select>
      </label>
      <label><span>连载期间</span><input v-model="form.serialsduri" placeholder="2013/04~2013/10" /></label>
      <label><span>存储路径</span><input v-model="form.stordir" /></label>
      <label><span>封面URL</span><input v-model="form.image_url" placeholder="https://..." /></label>
      <div class="flex gap-2">
        <button type="submit">保存</button>
        <button type="button" class="btn secondary" @click="$router.back()">取消</button>
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
  animenamecn: '', animenamejp: '', cast: '', type: '', status: '连载中',
  serialsduri: '', stordir: '', image_url: '',
})

onMounted(async () => {
  if (id) {
    const res = await axios.get(`/api/v1/anime/${id}`)
    if (res.data.anime) {
      const a = res.data.anime
      Object.assign(form.value, {
        animenamecn: a.animenamecn || '',
        animenamejp: a.animenamejp || '',
        cast: a.cast || '',
        type: a.type || '',
        status: a.status || '连载中',
        serialsduri: a.serialsduri || '',
        stordir: a.stordir || '',
        image_url: '',
      })
    }
  }
})

async function save() {
  if (isNew.value) {
    await axios.post('/api/v1/anime', form.value)
  } else {
    await axios.put(`/api/v1/anime/${id}`, form.value)
  }
  router.push(id ? '/anime/' + id : '/')
}
</script>
