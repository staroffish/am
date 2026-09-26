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
        <tr><th>日文名</th><th>中文名</th><th>类型</th><th>状态</th><th>连载期间</th><th>存储路径</th><th @click="toggleSort" style="cursor:pointer;">更新时间 {{ sortAsc ? '↑' : '↓' }}</th><th>操作</th></tr>
      </thead>
      <tbody>
        <tr v-for="anime in list" :key="anime.id">
          <td><router-link :to="'/anime/' + anime.id">{{ anime.animenamejp || '(无名称)' }}</router-link></td>
          <td>{{ anime.animenamecn }}</td>
          <td>{{ anime.type }}</td>
          <td class="col-narrow">{{ anime.status }}</td>
          <td>{{ anime.serialsduri }}</td>
          <td class="col-path">{{ anime.stordir }}</td>
          <td class="col-narrow">{{ fmtTime(anime.updated_at) }}</td>
          <td class="col-actions">
            <router-link :to="'/anime/' + anime.id + '/edit'" class="btn small secondary">编辑</router-link>
            <button class="small danger" @click="del(anime.id)">删除</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-else class="empty">暂无收藏动漫</div>

    <div v-if="totalPages > 1" class="pagination">
      <button :disabled="page <= 1" @click="goPage(page - 1)">上一页</button>
      <span class="page-info">{{ page }} / {{ totalPages }} (共 {{ total }} 条)</span>
      <button :disabled="page >= totalPages" @click="goPage(page + 1)">下一页</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import axios from 'axios'

const router = useRouter()
const route = useRoute()

const list = ref([] as any[])
const keyword = ref((route.query.keyword as string) || '')
const page = ref(parseInt(route.query.page as string) || 1)
const total = ref(0)
const pageSize = ref(50)
const sortAsc = ref(route.query.sort === 'asc')

const totalPages = computed(() => Math.ceil(total.value / pageSize.value))

function fmtTime(t: string) {
  if (!t || t.startsWith('0001')) return ''
  const d = new Date(t)
  return d.toLocaleString('zh-CN', { hour12: false })
}

async function fetchList() {
  const skip = (page.value - 1) * pageSize.value
  const res = await axios.get('/api/v1/anime', {
    params: { keyword: keyword.value, skip, sort: sortAsc.value ? 'asc' : 'desc' }
  })
  list.value = res.data.items || []
  total.value = res.data.total || 0
  if (res.data.limit) pageSize.value = res.data.limit
}

function toggleSort() {
  sortAsc.value = !sortAsc.value
  page.value = 1
  updateURL()
  fetchList()
}

function updateURL() {
  const q: any = {}
  if (keyword.value) q.keyword = keyword.value
  if (page.value > 1) q.page = page.value
  if (sortAsc.value) q.sort = 'asc'
  router.replace({ query: q })
}

async function search() {
  page.value = 1
  updateURL()
  fetchList()
}

async function goPage(p: number) {
  page.value = p
  updateURL()
  fetchList()
}

async function del(id: string) {
  if (!confirm('确定删除？')) return
  await axios.delete(`/api/v1/anime/${id}`)
  fetchList()
}

onMounted(fetchList)
</script>

<style scoped>
.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 16px;
  margin-top: 20px;
  padding: 12px;
}
.page-info {
  font-size: 13px;
  color: var(--text-dim);
}
</style>
