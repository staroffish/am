# am

动漫下载管理器：定时爬取磁力 → 按规则匹配 → 推给 qBittorrent，并维护一份番剧库。

## 存储

单实例部署，持久化只用 **一个 SQLite 文件**（GORM + `github.com/glebarez/sqlite`，纯 Go，无需 CGO）。

- `anime`：番剧（含封面图 BLOB）
- `download_tasks`：下载规则
- `magnets`：爬虫抓到的磁力，按日期归档（替代原来的 Redis key）

> ⚠️ `sqlite.path` 必须指向**本机磁盘**。SQLite 放在 NFS/CIFS 网络共享盘上会有锁失效甚至损坏的风险。

原来的 MongoDB 与 Redis 已完全移除。

## 构建

```bash
make build        # 前端 + 服务端，产物 bin/am
make build-server # 只编服务端
make migrate      # 编旧数据导入工具，产物 bin/migrate（一次性，可删）
```

## 运行

```bash
bin/am -config configs/config.yaml
```

配置见 `configs/config.yaml`（该文件不入库），关键项：

```yaml
sqlite:
  path: "./data/am.db"   # 本机磁盘路径
```

## 从旧的 MongoDB/Redis 版本迁移

旧服务还在运行时执行（导入工具走旧服务的 HTTP API，不依赖 mongo/redis 驱动）：

```bash
make migrate
bin/migrate -old http://192.168.3.5:8222 -db ./data/am.db
# 可选：-skip-images / -skip-magnets / -dry-run
```

然后停掉旧服务，用新二进制启动即可。导入是一次性的，迁移完成后 `cmd/migrate` 可以删掉。
