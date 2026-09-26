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

## 从旧的 MongoDB/Redis 版本迁移（离线）

导入工具在独立子模块 `tools/legacyimport` 里，**直连 MongoDB + Redis**，
不需要旧服务进程在跑（只要求两个库网络可达）。它不会把 mongo/redis 依赖带进主模块。

```bash
make migrate

# 先看看能读到多少数据，不写库
bin/migrate -dry-run -mongo-pass 'xxx' -redis-pass 'xxx'

# 正式导入（-db 指向本机磁盘）
bin/migrate -db /本机磁盘/am.db -mongo-pass 'xxx' -redis-pass 'xxx'
```

密码也可以用环境变量 `MONGO_PASS` / `REDIS_PASS` 传。常用参数：

| 参数 | 默认值 | 说明 |
|---|---|---|
| `-mongo` / `-mongo-db` / `-mongo-user` | `mongodb://mongo.holygrail.com:27017` / `am` / `am` | Mongo 连接 |
| `-redis` / `-redis-db` | `redis.holygrail.com:6379` / `2` | Redis 连接 |
| `-db` | `data/am.db` | 目标 SQLite 文件 |
| `-skip-images` / `-skip-magnets` | false | 跳过封面图 / 磁力缓存 |
| `-dry-run` | false | 只统计 |

导入是幂等的，可以重复跑。迁完之后 `tools/legacyimport` 就可以删掉了。
