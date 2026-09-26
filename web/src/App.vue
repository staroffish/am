<template>
  <div id="am-app">
    <header class="app-header">
      <h1 class="app-title">动漫管理</h1>
      <nav class="app-nav">
        <router-link to="/">收藏</router-link>
        <router-link to="/tasks">下载规则</router-link>
        <router-link to="/downloads">下载任务</router-link>
      </nav>
    </header>
    <main class="app-main">
      <router-view />
    </main>
  </div>
</template>

<script setup lang="ts">
</script>

<style>
  :root {
    --bg: #f5f6fa;
    --surface: #fff;
    --border: #e2e5ed;
    --primary: #4f6ef7;
    --text: #2c3e50;
    --text-dim: #8896a8;
    --accent-green: #22b07d;
    --accent-red: #e74c3c;
    --accent-yellow: #f39c12;
    --radius: 8px;
    --shadow: 0 1px 3px rgba(0,0,0,0.04), 0 1px 2px rgba(0,0,0,0.06);
  }

  * { margin: 0; padding: 0; box-sizing: border-box; }
  body {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'PingFang SC', 'Hiragino Sans GB', sans-serif;
    background: var(--bg);
    color: var(--text);
    line-height: 1.6;
    min-height: 100vh;
  }

  .app-header {
    background: var(--surface);
    border-bottom: 1px solid var(--border);
    padding: 0 24px;
    height: 56px;
    display: flex;
    align-items: center;
    gap: 32px;
    position: sticky;
    top: 0;
    z-index: 100;
    box-shadow: var(--shadow);
  }

  .app-title {
    font-size: 20px;
    font-weight: 700;
    color: var(--primary);
    letter-spacing: 2px;
  }

  .app-nav {
    display: flex;
    gap: 4px;
  }

  .app-nav a {
    color: var(--text-dim);
    text-decoration: none;
    padding: 8px 16px;
    border-radius: var(--radius);
    font-size: 14px;
    transition: all 0.15s;
  }

  .app-nav a:hover { color: var(--text); background: var(--bg); }
  .app-nav a.router-link-active { color: var(--primary); background: rgba(79,110,247,0.08); font-weight: 600; }

  .app-main {
    max-width: 1800px;
    margin: 0 auto;
    padding: 24px;
  }

  table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 16px;
    background: var(--surface);
    border-radius: var(--radius);
    overflow: hidden;
    box-shadow: var(--shadow);
    table-layout: auto;
  }
  th, td {
    text-align: left;
    padding: 10px 12px;
    border-bottom: 1px solid var(--border);
    font-size: 13px;
    vertical-align: middle;
    overflow: hidden;
  }
  td.col-narrow { white-space: nowrap; width: 1%; }
  td.col-path { white-space: normal; word-break: break-all; font-size: 12px; color: var(--text-dim); }
  td.col-actions { white-space: nowrap; width: 1%; }
  th { color: var(--text-dim); font-weight: 600; font-size: 12px; text-transform: uppercase; letter-spacing: 0.5px; background: var(--bg); }
  tr:last-child td { border-bottom: none; }
  tr:hover td { background: var(--bg); }

  button, a.btn, .btn {
    background: var(--primary);
    color: #fff;
    border: none;
    padding: 6px 12px;
    border-radius: var(--radius);
    font-size: 13px;
    line-height: 1.4;
    cursor: pointer;
    text-decoration: none;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    gap: 4px;
    transition: all 0.15s;
    font-weight: 500;
    vertical-align: middle;
    font-family: inherit;
    height: 30px;
    box-sizing: border-box;
  }
  button.small, a.btn.small, .btn.small {
    padding: 4px 10px;
    font-size: 12px;
    height: 26px;
  }
  button:hover, a.btn:hover, .btn:hover { opacity: 0.85; transform: translateY(-1px); }
  button:disabled { opacity: 0.5; cursor: not-allowed; transform: none; }
  button.danger { background: var(--accent-red); }
  button.secondary, a.btn.secondary { background: var(--bg); color: var(--text); border: 1px solid var(--border); }
  button.small { padding: 4px 10px; font-size: 12px; }

  input, select {
    background: var(--surface);
    border: 1px solid var(--border);
    border-radius: var(--radius);
    padding: 9px 12px;
    color: var(--text);
    font-size: 14px;
    width: 100%;
    transition: border-color 0.15s;
  }
  input:focus { outline: none; border-color: var(--primary); box-shadow: 0 0 0 3px rgba(79,110,247,0.1); }

  label { display: block; margin-bottom: 14px; font-size: 13px; color: var(--text-dim); }
  label span { display: block; margin-bottom: 4px; font-weight: 500; color: var(--text); }

  form {
    background: var(--surface);
    border: 1px solid var(--border);
    border-radius: var(--radius);
    padding: 24px;
    max-width: 560px;
    box-shadow: var(--shadow);
  }

  h2 { font-size: 20px; margin-bottom: 16px; font-weight: 700; }
  h3 { font-size: 16px; margin-bottom: 12px; color: var(--text-dim); }

  .toolbar {
    display: flex;
    gap: 10px;
    margin-bottom: 20px;
    align-items: center;
    flex-wrap: wrap;
  }

  .status-badge {
    display: inline-block;
    padding: 2px 10px;
    border-radius: 12px;
    font-size: 12px;
    font-weight: 500;
  }
  .status-downloading { background: rgba(79,110,247,0.1); color: var(--primary); }
  .status-paused { background: rgba(243,156,18,0.1); color: var(--accent-yellow); }
  .status-seeding, .status-finished { background: rgba(34,176,125,0.1); color: var(--accent-green); }
  .status-error { background: rgba(231,76,60,0.1); color: var(--accent-red); }

  .progress-bar { width: 100%; height: 4px; background: var(--border); border-radius: 2px; overflow: hidden; }
  .progress-bar-fill { height: 100%; background: var(--primary); transition: width 0.3s; }

  .card {
    background: var(--surface);
    border: 1px solid var(--border);
    border-radius: var(--radius);
    padding: 20px;
    margin-bottom: 16px;
    box-shadow: var(--shadow);
  }

  .empty { text-align: center; padding: 48px; color: var(--text-dim); }

  .modal-mask { position: fixed; inset: 0; background: rgba(0,0,0,0.3); z-index: 200; display: flex; align-items: center; justify-content: center; }
  .modal-box { background: var(--surface); border-radius: var(--radius); padding: 24px; max-width: 480px; width: 90%; box-shadow: 0 4px 24px rgba(0,0,0,0.12); }
  .modal-box h3 { margin-bottom: 16px; }

  .full-path { word-break: break-all; font-size: 12px; color: var(--text-dim); max-width: 300px; }

  .flex { display: flex; gap: 8px; align-items: center; }
  .flex-end { justify-content: flex-end; }
  .gap-2 { gap: 16px; }
  .mt-2 { margin-top: 16px; }
  .mb-2 { margin-bottom: 16px; }
</style>
