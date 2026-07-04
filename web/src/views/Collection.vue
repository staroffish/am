<template>
  <div>
    <h2>动漫收藏</h2>
    <div class="toolbar">
      <input v-model="keyword" placeholder="搜索中文名/日文名/CAST/类型..." @keyup.enter="search" style="max-width: 360px;" />
      <button @click="search">搜索</button>
      <router-link to="/anime/new" class="btn secondary">+ 新增动漫</router-link>
    </div>
    <table v-if="list.length">
      <thead>
        <tr><th>中文名</th><th>日文名</th><th>类型</th><th>状态</th><th>连载期间</th><th>更新时间</th><th>操作</th></tr>
      </thead>
      <tbody>
        <tr v-for="anime in list" :key="anime.id">
          <td><router-link :to="'/anime/' + anime.id">{{ anime.animenamecn || anime.animenamejp || '(无名称)' }}</router-link></td>
          <td>{{ anime.animenamejp }}</td>
          <td>{{ anime.type }}</td>
          <td>{{ anime.status }}</td>
          <td>{{ anime.serialsduri }}</td>
          <td style="white-space: nowrap;">{{ fmtTime(anime.updated_at) }}</td>
          <td class="flex">
            <router-link :to="'/anime/' + anime.id + '/edit'" class="btn small secondary">编辑</router-link>
            <button class="small danger" @click="del(anime.id)">删除</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-else class="empty">没有找到动漫</div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'

const list = ref([] as any[])
const keyword = ref('')

function fmtTime(t: string) {
  if (!t || t.startsWith('0001')) return ''
  const d = new Date(t)
  return d.toLocaleString('zh-CN', { hour12: false })
}

async function fetchList() {
  const res = await axios.get('/api/v1/anime', { params: { keyword: keyword.value } })
  list.value = res.data || []
}

async function search() { fetchList() }

async function del(id: string) {
  if (!confirm('确定删除？')) return
  await axios.delete(`/api/v1/anime/${id}`)
  fetchList()
}

onMounted(fetchList)
</script>
