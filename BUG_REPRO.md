# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

检察服务站修改骑手诉求标题时，接口允许操作人字段只填空格，审计记录因此无法追溯是谁改了材料。请修复事项修改校验并在审计中保存整理后的身份；现有测试用来验收，这次不得修改测试文件。

## 含 Bug 版本

- 仓库：11DingKing/rider-rights-task-28
- 仓库地址：https://github.com/11DingKing/rider-rights-task-28.git
- parent SHA：9dc556cd838bf6b416f8675f02e3ffa7260fc151

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/rider-rights-task-28.git bug-repro
cd bug-repro
git checkout --detach 9dc556cd838bf6b416f8675f02e3ffa7260fc151
go test ./internal/domain -run "^TestModificationRequiresActor$" -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestModificationRequiresActor$" -count=1
--- FAIL: TestModificationRequiresActor (0.00s)
    task28_test.go:7: blank modification actor was accepted
FAIL
FAIL	riderguard/internal/domain	0.066s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestModificationRequiresActor$" -count=1
--- FAIL: TestModificationRequiresActor (0.00s)
    task28_test.go:7: blank modification actor was accepted
FAIL
FAIL	riderguard/internal/domain	0.002s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

修复后，空白或仅含制表符的修改操作人必须在写入事项前返回校验错误；有效身份去除首尾空格后写入审计，原有标题、材料和类别修改行为保持可用。定向测试、相关包测试及全量回归必须通过，不得删除、跳过或削弱测试。
