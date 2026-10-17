# AGENTS.md — am 项目工作约定

本文件由 DSH 的 agent-instructions 机制自动注入到每个会话，属于本仓库的硬性约定。

## Git 硬性规则（任何情况下都不可违反）

1. **禁止 force push**
   - 不使用 `git push --force`、`--force-with-lease`、`-f`，包括对自己刚创建的分支。
   - 推送被拒绝时先 `git fetch` 查清原因；需要改历史时改用新分支名，或追加新提交，而不是覆盖远端。

2. **禁止直接向主分支 `v1.0` 提交或推送**
   - `v1.0` 是主分支，只能通过 GitHub 上的 PR 合并来更新。
   - 所有改动先建特性分支（`feature/*`、`fix/*`、`chore/*`），推送分支后开 PR。
   - 不要因为"改动很小"就直接提交到 `v1.0`。

3. **推送前自检**
   - `git branch --show-current` 不能是 `v1.0`。
   - 当前分支的上游不能是 `origin/v1.0`：`git switch -c <branch> origin/v1.0` 会把上游继承为 `origin/v1.0`，必须先 `git branch --unset-upstream`，再显式 `git push -u origin <branch>`。
   - push 命令里显式写明目标分支，避免无参数 `git push` 推到主分支。

4. **提交范围**
   - 只提交本次改动相关的文件，禁止 `git add -A` / `git add .`。
   - 纯文件权限变更（644 ↔ 755）和构建产物不要提交。

## 其他约定

- **不要提交含真实密钥的文件**：`configs/config.yaml`、`configs/config_test.yaml` 里有真实 API key，保持未跟踪状态。
- 本机文件系统不支持创建 symlink，`tests/node_modules/.bin/*` 会一直显示类型变更，忽略即可（历史上已把整个 `tests/` 从仓库移除）。
- 提交信息使用中文 + conventional 前缀（`feat:` / `fix:` / `chore:` / `refactor:`）。
- 改完 Go 代码至少跑 `go build ./...` 和 `go test ./...` 再提交；注意本机 `GOPATH`/`GOCACHE` 可能只读，必要时用 `GOPATH=/tmp/... GOCACHE=/tmp/...`。
