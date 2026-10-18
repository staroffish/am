<template>
  <div>
    <h2>Clash 中转</h2>

    <!-- 订阅源地址 -->
    <div class="card">
      <div class="row">
        <b>订阅源地址</b>
        <code v-if="settings.airport_url">{{ settings.airport_url }}</code>
        <span v-else class="dim">未配置，请点「编辑」填写机场订阅地址</span>
      </div>
      <div class="row">
        <button class="small" :disabled="pulling" @click="pull">
          {{ pulling ? '拉取中…' : '拉取并生成' }}
        </button>
        <button class="small secondary" @click="showSourceEditor = !showSourceEditor">
          {{ showSourceEditor ? '收起' : '编辑' }}
        </button>
      </div>
    </div>

    <!-- 编辑订阅源（默认隐藏） -->
    <div v-if="showSourceEditor" class="card">
      <label>
        <span>机场订阅 URL</span>
        <input v-model="settings.airport_url" placeholder="https://..." />
      </label>
      <label>
        <span>User-Agent</span>
        <input v-model="settings.airport_ua" placeholder="clash.meta" />
      </label>
      <div class="toolbar" style="margin-bottom:0;">
        <button @click="saveSource">保存订阅源</button>
        <button class="secondary" @click="showSourceEditor = false">取消</button>
      </div>
    </div>

    <!-- 版本信息 -->
    <div class="card">
      <div class="row">
        <span>当前版本：</span>
        <span v-if="status.current">
          #{{ status.current.id }} · {{ fmtTime(status.current.created_at) }} ·
          {{ sourceLabel(status.current.source) }}
        </span>
        <span v-else class="dim">暂无，请先「拉取并生成」</span>
      </div>
      <div class="row">
        <span>上次拉取：</span>
        <span v-if="status.settings.last_pull_at">
          {{ fmtTime(status.settings.last_pull_at) }} ·
          {{ pullStatusLabel(status.settings.last_pull_status) }}
          <span v-if="status.settings.last_pull_error" class="err-text">
            （{{ status.settings.last_pull_error }}）
          </span>
        </span>
        <span v-else class="dim">从未拉取</span>
      </div>
      <div class="row">
        <b>OpenClash 订阅地址</b>
        <code class="copyable">{{ subUrl }}</code>
        <button class="small secondary" @click="copySub">{{ copied ? '已复制' : '复制' }}</button>
      </div>
    </div>

    <p v-if="pullMessage" class="msg">{{ pullMessage }}</p>

    <!-- 版本列表 -->
    <div class="toolbar">
      <button class="small secondary" @click="showTransform = true">编辑转换设置</button>
      <button class="small secondary" @click="loadAll">刷新</button>
      <span class="hint">保留 {{ status.version_count }} / {{ settings.keep_versions }} 个版本</span>
    </div>
    <table>
      <thead>
        <tr><th>版本</th><th>时间</th><th>来源</th><th>机场节点数</th><th>状态</th><th>操作</th></tr>
      </thead>
      <tbody>
        <tr v-if="!versions.length">
          <td colspan="6" class="empty-cell">还没有版本，点「拉取并生成」创建第一个版本</td>
        </tr>
        <tr v-for="v in versions" :key="v.id">
          <td class="col-narrow">#{{ v.id }}</td>
          <td class="col-narrow">{{ fmtTime(v.created_at) }}</td>
          <td class="col-narrow">{{ sourceLabel(v.source) }}</td>
          <td class="col-narrow">{{ v.node_count }}</td>
          <td class="col-narrow">
            <span v-if="v.is_current" class="status-badge status-finished">当前</span>
            <span v-else class="dim">—</span>
          </td>
          <td class="col-actions">
            <button class="small secondary" @click="openVersion(v.id)">编辑</button>
            <button v-if="!v.is_current" class="small" @click="activate(v.id)">设为当前</button>
            <button v-if="!v.is_current" class="small danger" @click="removeVersion(v.id)">删除</button>
            <button class="small secondary" @click="download(v.id)">下载</button>
          </td>
        </tr>
      </tbody>
    </table>

    <!-- YAML 编辑器 -->
    <div v-if="editor.open" class="modal-mask" @click.self="closeEditor">
      <div class="modal-box modal-wide">
        <h3>{{ editor.title }}</h3>
        <div class="editor-wrap">
          <pre ref="highlightRef" class="yaml-highlight" aria-hidden="true" v-html="highlightedContent"></pre>
          <textarea
            v-model="editor.content"
            class="yaml-input"
            :readonly="editor.readonly"
            spellcheck="false"
            wrap="off"
            @scroll="onEditorScroll"
          ></textarea>
        </div>
        <ul v-if="editor.errors.length" class="err-list">
          <li v-for="(e, i) in editor.errors" :key="i">{{ e }}</li>
        </ul>
        <div class="toolbar" style="justify-content:flex-end;margin:12px 0 0;">
          <button class="secondary" @click="closeEditor">关闭</button>
          <button v-if="!editor.readonly" :disabled="editor.saving" @click="saveVersion">
            {{ editor.saving ? '保存中…' : '保存为新版本' }}
          </button>
        </div>
      </div>
    </div>

    <!-- 转换设置弹窗 -->
    <div v-if="showTransform" class="modal-mask" @click.self="showTransform = false">
      <div class="modal-box modal-settings">
        <h3>转换设置</h3>
        <div class="form-grid">
          <label><span>VPS 地址</span><input v-model="settings.vps_host" placeholder="vps.example.com" /></label>
          <label><span>VPS 端口</span><input v-model.number="settings.vps_port" type="number" /></label>
          <label><span>SS 加密方式</span><input v-model="settings.vps_cipher" placeholder="aes-256-gcm" /></label>
          <label><span>SS 密码</span><input v-model="settings.vps_password" type="password" /></label>
          <label><span>VPS 节点名</span><input v-model="settings.vps_node_name" placeholder="VPS-SS" /></label>
          <label><span>VPS 所在 proxy-groups 名称</span><input v-model="settings.vps_group_name" placeholder="PROXY" /></label>
          <label><span>保留版本数</span><input v-model.number="settings.keep_versions" type="number" min="1" /></label>
        </div>
        <label>
          <span>rules（整段替换；支持 # 注释，注释会保留到生成的配置里）</span>
          <textarea v-model="settings.rules_text" class="yaml-editor" style="min-height:220px;" spellcheck="false" placeholder="# 国内直连&#10;- GEOIP,CN,DIRECT&#10;# 其余走代理&#10;- MATCH,PROXY"></textarea>
        </label>
        <label class="checkbox-row">
          <input type="checkbox" v-model="settings.rule_providers_enabled" />
          <span>启用自定义 rule-anchor / rule-providers（会覆盖机场自带的）</span>
        </label>
        <label>
          <span>rule-anchor / rule-providers（启用后写入配置；可同时贴 rule-anchor 和 rule-providers）</span>
          <textarea
            v-model="settings.rule_providers_text"
            class="yaml-editor"
            style="min-height:160px;"
            spellcheck="false"
            :disabled="!settings.rule_providers_enabled"
            placeholder="rule-anchor:&#10;  ip: &amp;ip {type: http, interval: 86400, behavior: ipcidr, format: mrs}&#10;  domain: &amp;domain {type: http, interval: 86400, behavior: domain, format: mrs}&#10;rule-providers:&#10;  cn_domain:&#10;    &lt;&lt;: *domain&#10;    url: https://..."></textarea>
        </label>
        <div class="toolbar" style="justify-content:flex-end;margin:12px 0 0;">
          <button class="secondary" @click="showTransform = false">取消</button>
          <button @click="saveTransform">保存设置</button>
        </div>
        <p class="hint">保存设置不会重新生成版本；修改后请关闭弹窗并点「拉取并生成」。</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import axios from 'axios'

const subUrl = `${location.origin}/api/v1/clash-relay/sub`

const defaultSettings = {
  airport_url: '',
  airport_ua: 'clash.meta',
  vps_host: '',
  vps_port: 8443,
  vps_cipher: 'aes-256-gcm',
  vps_password: '',
  vps_node_name: 'VPS-SS',
  vps_group_name: 'PROXY',
  keep_versions: 10,
  rules_text: '',
  rule_providers_enabled: false,
  rule_providers_text: '',
  last_pull_at: null as string | null,
  last_pull_status: '',
  last_pull_error: '',
}

const settings = reactive({ ...defaultSettings })
const versions = ref<any[]>([])
const status = reactive({
  current: null as any,
  version_count: 0,
  settings: { last_pull_at: null as string | null, last_pull_status: '', last_pull_error: '' },
})

const pulling = ref(false)
const pullMessage = ref('')
const showSourceEditor = ref(false)
const showTransform = ref(false)
const pendingSource = ref<'pull' | 'manual'>('manual')
const pendingDraftContent = ref('')
const pendingAirportHash = ref('')

const editor = reactive({
  open: false,
  readonly: false,
  content: '',
  parentId: null as number | null,
  title: '',
  errors: [] as string[],
  saving: false,
})

function fmtTime(t: string | null) {
  if (!t) return ''
  const d = new Date(t)
  if (isNaN(d.getTime())) return t
  return d.toLocaleString('zh-CN', { hour12: false })
}

function sourceLabel(s: string) {
  return s === 'manual' ? '手动编辑' : s === 'pull' ? '拉取生成' : s || '—'
}

function pullStatusLabel(s: string) {
  switch (s) {
    case 'ok': return '成功'
    case 'unchanged': return '无变化'
    case 'draft': return '已拉取（未保存）'
    case 'error': return '失败'
    default: return '—'
  }
}

async function loadSettings() {
  const res = await axios.get('/api/v1/clash-relay/settings')
  Object.assign(settings, defaultSettings, res.data || {})
}

async function loadVersions() {
  const res = await axios.get('/api/v1/clash-relay/versions')
  versions.value = res.data || []
}

async function loadStatus() {
  const res = await axios.get('/api/v1/clash-relay/status')
  const data = res.data || {}
  status.current = data.current || null
  status.version_count = data.version_count || 0
  status.settings = {
    last_pull_at: data.settings?.last_pull_at ?? null,
    last_pull_status: data.settings?.last_pull_status ?? '',
    last_pull_error: data.settings?.last_pull_error ?? '',
  }
}

async function loadAll() {
  await Promise.all([loadSettings(), loadVersions(), loadStatus()])
}

async function saveSource() {
  try {
    const res = await axios.put('/api/v1/clash-relay/settings', {
      airport_url: settings.airport_url,
      airport_ua: settings.airport_ua,
    })
    Object.assign(settings, defaultSettings, res.data || {})
    pullMessage.value = '订阅源已保存'
    showSourceEditor.value = false
    await loadStatus()
  } catch (e: any) {
    pullMessage.value = '保存失败: ' + (e.response?.data?.error || e.message)
  }
}

async function saveTransform() {
  try {
    const res = await axios.put('/api/v1/clash-relay/settings', {
      vps_host: settings.vps_host,
      vps_port: Number(settings.vps_port) || 0,
      vps_cipher: settings.vps_cipher,
      vps_password: settings.vps_password,
      vps_node_name: settings.vps_node_name,
      vps_group_name: settings.vps_group_name,
      keep_versions: Number(settings.keep_versions) || 10,
      rules_text: settings.rules_text,
      rule_providers_enabled: !!settings.rule_providers_enabled,
      rule_providers_text: settings.rule_providers_text,
    })
    Object.assign(settings, defaultSettings, res.data || {})
    pullMessage.value = '转换设置已保存'
    showTransform.value = false
    await loadStatus()
  } catch (e: any) {
    pullMessage.value = '保存失败: ' + (e.response?.data?.error || e.message)
  }
}

async function pull() {
  pulling.value = true
  pullMessage.value = ''
  try {
    const res = await axios.post('/api/v1/clash-relay/pull')
    const data = res.data || {}
    pullMessage.value = data.message || '已生成草稿'
    await loadAll()
    if (data.draft?.content) {
      pendingSource.value = 'pull'
      pendingDraftContent.value = data.draft.content
      pendingAirportHash.value = data.draft.airport_hash || ''
      openEditor(data.draft.content, null, '拉取草稿（未保存）· 保存后生成版本')
    }
  } catch (e: any) {
    pullMessage.value = '拉取失败: ' + (e.response?.data?.error || e.message)
    await loadStatus()
  } finally {
    pulling.value = false
  }
}

async function openVersion(id: number) {
  try {
    const res = await axios.get(`/api/v1/clash-relay/versions/${id}`)
    const v = res.data
    pendingSource.value = 'manual'
    pendingDraftContent.value = ''
    pendingAirportHash.value = ''
    openEditor(v.content || '', v.id, `编辑 · 基于版本 #${v.id}`)
  } catch (e: any) {
    pullMessage.value = '读取版本失败: ' + (e.response?.data?.error || e.message)
  }
}

async function removeVersion(id: number) {
  if (!confirm(`确定删除版本 #${id}？删除后不可恢复。`)) return
  try {
    await axios.delete(`/api/v1/clash-relay/versions/${id}`)
    pullMessage.value = `已删除版本 #${id}`
    await loadAll()
  } catch (e: any) {
    pullMessage.value = '删除失败: ' + (e.response?.data?.error || e.message)
  }
}

function openEditor(content: string, parentId: number | null, title: string) {
  editor.open = true
  editor.readonly = false
  editor.content = content
  editor.parentId = parentId
  editor.title = title
  editor.errors = []
  editor.saving = false
}

function closeEditor() {
  editor.open = false
}

async function saveVersion() {
  editor.saving = true
  editor.errors = []
  try {
    const editedFromPull = pendingSource.value === 'pull' && editor.content !== pendingDraftContent.value
    const res = await axios.post('/api/v1/clash-relay/versions', {
      content: editor.content,
      parent_id: editor.parentId,
      source: editedFromPull ? 'manual' : pendingSource.value,
      airport_hash: pendingAirportHash.value,
    })
    editor.open = false
    pendingSource.value = 'manual'
    pendingDraftContent.value = ''
    pendingAirportHash.value = ''
    pullMessage.value = res.data?.unchanged
      ? '内容与当前版本一致，未生成新版本'
      : `已保存为新版本 #${res.data?.version?.id}`
    await loadAll()
  } catch (e: any) {
    const msg = e.response?.data?.error || e.message
    editor.errors = String(msg).split('；').filter(Boolean)
  } finally {
    editor.saving = false
  }
}

async function activate(id: number) {
  try {
    await axios.post(`/api/v1/clash-relay/versions/${id}/activate`)
    await loadAll()
  } catch (e: any) {
    pullMessage.value = '切换失败: ' + (e.response?.data?.error || e.message)
  }
}

async function download(id: number) {
  try {
    const res = await axios.get(`/api/v1/clash-relay/versions/${id}`)
    const blob = new Blob([res.data?.content || ''], { type: 'text/yaml;charset=utf-8' })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = `clash-relay-v${id}.yaml`
    a.click()
    URL.revokeObjectURL(url)
  } catch (e: any) {
    pullMessage.value = '下载失败: ' + (e.response?.data?.error || e.message)
  }
}

const copied = ref(false)

// HTTP（非安全上下文）下 navigator.clipboard 不可用，需要回退到 execCommand。
// 兼容 http://内网IP 直接访问的场景。
async function copyText(text: string): Promise<boolean> {
  if (window.isSecureContext && navigator.clipboard && navigator.clipboard.writeText) {
    try {
      await navigator.clipboard.writeText(text)
      return true
    } catch {
      // 继续走 fallback
    }
  }

  const ta = document.createElement('textarea')
  ta.value = text
  ta.setAttribute('readonly', '')
  ta.style.position = 'fixed'
  ta.style.top = '0'
  ta.style.left = '0'
  ta.style.width = '1px'
  ta.style.height = '1px'
  ta.style.padding = '0'
  ta.style.border = 'none'
  ta.style.outline = 'none'
  ta.style.opacity = '0'
  document.body.appendChild(ta)

  const prev = document.activeElement as HTMLElement | null
  ta.select()
  ta.setSelectionRange(0, text.length)

  let ok = false
  try {
    ok = document.execCommand('copy')
  } catch {
    ok = false
  }
  document.body.removeChild(ta)
  if (prev && typeof prev.focus === 'function') prev.focus()
  return ok
}

async function copySub() {
  const ok = await copyText(subUrl)
  if (ok) {
    copied.value = true
    pullMessage.value = '订阅地址已复制到剪贴板'
    window.setTimeout(() => { copied.value = false }, 1500)
  } else {
    pullMessage.value = '自动复制失败，请手动复制下面的地址'
    window.prompt('自动复制失败，请手动复制：', subUrl)
  }
}

const highlightRef = ref<HTMLElement | null>(null)
const highlightedContent = computed(() => highlightYaml(editor.content))

function onEditorScroll(e: Event) {
  const ta = e.target as HTMLTextAreaElement
  const pre = highlightRef.value
  if (pre) {
    pre.scrollTop = ta.scrollTop
    pre.scrollLeft = ta.scrollLeft
  }
}

function escHtml(s: string): string {
  return s.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;')
}

function splitComment(line: string): { code: string; comment: string } {
  let inS = false
  let inD = false
  for (let i = 0; i < line.length; i++) {
    const c = line[i]
    if (c === "'" && !inD) inS = !inS
    else if (c === '"' && !inS) inD = !inD
    else if (c === '#' && !inS && !inD && (i === 0 || /\s/.test(line[i - 1]))) {
      return { code: line.slice(0, i), comment: line.slice(i) }
    }
  }
  return { code: line, comment: '' }
}

function highlightValue(v: string): string {
  const e = escHtml(v)
  return e.replace(/(&quot;.*?&quot;|&#39;.*?&#39;)/g, '<span class="y-str">$1</span>')
}

function highlightLine(line: string): string {
  const { code, comment } = splitComment(line)
  let out = ''
  const km = code.match(/^(\s*)(-\s+)?([^:#]+?)(:)(\s?)/)
  if (km) {
    out += escHtml(km[1])
    if (km[2]) out += '<span class="y-dash">-</span> '
    out += `<span class="y-key">${escHtml(km[3])}</span>:`
    if (km[5]) out += ' '
    out += highlightValue(code.slice(km[0].length))
  } else {
    const lm = code.match(/^(\s*)(-)(\s+)(.*)$/)
    if (lm) {
      out += escHtml(lm[1]) + '<span class="y-dash">-</span>' + escHtml(lm[3]) + highlightValue(lm[4])
    } else {
      out += highlightValue(code)
    }
  }
  if (comment) out += `<span class="y-comment">${escHtml(comment)}</span>`
  return out
}

function highlightYaml(src: string): string {
  return src.split('\n').map(highlightLine).join('\n')
}

onMounted(loadAll)
</script>

<style scoped>
.row { display: flex; align-items: center; gap: 8px; flex-wrap: wrap; margin-bottom: 6px; }
.row:last-child { margin-bottom: 0; }
code {
  background: var(--bg);
  padding: 3px 8px;
  border-radius: 4px;
  word-break: break-all;
  font-size: 12px;
}
.copyable { user-select: all; }
.dim { color: var(--text-dim); }
.err-text { color: var(--accent-red); }
.msg { margin: 0 0 16px; font-size: 13px; }
.hint { font-size: 12px; color: var(--text-dim); margin-top: 10px; }

.checkbox-row { display: flex; align-items: center; gap: 8px; margin-bottom: 10px; }
.checkbox-row input { width: auto; }
.checkbox-row span { display: inline; margin: 0; }

.form-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 0 16px;
}

textarea.yaml-editor {
  background: var(--bg);
  border: 1px solid var(--border);
  border-radius: var(--radius);
  padding: 12px;
  color: var(--text);
  width: 100%;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 12px;
  line-height: 1.5;
  resize: vertical;
}
textarea.yaml-editor:focus { outline: none; border-color: var(--primary); box-shadow: 0 0 0 3px rgba(79,110,247,0.1); }

.modal-wide {
  max-width: 1100px;
  width: 94vw;
  max-height: 90vh;
  display: flex;
  flex-direction: column;
}
.modal-wide .editor-wrap { flex: 1; }

.editor-wrap {
  position: relative;
  flex: 1;
  min-height: 55vh;
  border: 1px solid var(--border);
  border-radius: var(--radius);
  background: var(--bg);
  overflow: hidden;
}
.yaml-highlight,
.yaml-input {
  position: absolute;
  inset: 0;
  margin: 0;
  padding: 12px;
  border: 0;
  width: 100%;
  height: 100%;
  font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
  font-size: 12px;
  line-height: 1.5;
  white-space: pre;
  tab-size: 2;
}
.yaml-highlight {
  pointer-events: none;
  overflow: hidden;
  color: var(--text);
}
.yaml-input {
  background: transparent;
  color: transparent;
  caret-color: var(--text);
  resize: none;
  outline: none;
  overflow: auto;
}
.yaml-input::selection { background: rgba(79,110,247,0.25); }
.yaml-highlight :deep(.y-key) { color: #4f6ef7; }
.yaml-highlight :deep(.y-str) { color: #22b07d; }
.yaml-highlight :deep(.y-comment) { color: var(--text-dim); font-style: italic; }
.yaml-highlight :deep(.y-dash) { color: #e74c3c; }

.modal-settings {
  max-width: 820px;
  width: 94vw;
  max-height: 90vh;
  overflow-y: auto;
}

.empty-cell {
  text-align: center;
  color: var(--text-dim);
  padding: 28px 12px;
}

.err-list {
  color: var(--accent-red);
  font-size: 12px;
  margin: 10px 0 0;
  padding-left: 18px;
}
</style>
