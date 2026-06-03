---
name: clscli
description: Use when querying Tencent Cloud CLS logs with the packaged clscli binary, including listing topics, searching logs, SQL analysis, writing CSV/JSON output, and fetching log context.
metadata:
  {"requires": {"bin": ["bin/clscli-windows-amd64.exe", "bin/clscli-linux-amd64", "bin/clscli-linux-arm64", "bin/clscli-darwin-amd64", "bin/clscli-darwin-arm64", "bin/clscli-windows-arm64.exe"], "env": ["TENCENTCLOUD_SECRET_ID", "TENCENTCLOUD_SECRET_KEY"]}}
---

# CLS CLI Skill

腾讯云日志服务 CLS (Cloud Log Service) 查询技能包。当前技能包内置 `clscli` 二进制，命令面以 `topics`、`query`、`context` 为主。

## Packaged Binaries

```text
SKILL/
  SKILL.md
  bin/
    clscli-windows-amd64.exe
    clscli-windows-arm64.exe
    clscli-linux-amd64
    clscli-linux-arm64
    clscli-darwin-amd64
    clscli-darwin-arm64
```

| Platform | Binary |
|---|---|
| Windows x64 | `bin/clscli-windows-amd64.exe` |
| Windows ARM64 | `bin/clscli-windows-arm64.exe` |
| Linux x64 | `bin/clscli-linux-amd64` |
| Linux ARM64 | `bin/clscli-linux-arm64` |
| macOS Intel | `bin/clscli-darwin-amd64` |
| macOS Apple Silicon | `bin/clscli-darwin-arm64` |

## Configuration

Set Tencent Cloud credentials before calling API-backed commands.

Windows PowerShell:

```powershell
$env:TENCENTCLOUD_SECRET_ID = "xxx"
$env:TENCENTCLOUD_SECRET_KEY = "xxx"
```

macOS/Linux:

```bash
export TENCENTCLOUD_SECRET_ID="xxx"
export TENCENTCLOUD_SECRET_KEY="xxx"
```

Always pass the CLS region explicitly:

```bash
--region ap-shanghai
```

## Quick Start

Windows PowerShell:

```powershell
.\bin\clscli-windows-amd64.exe --help
.\bin\clscli-windows-amd64.exe topics --region ap-shanghai
```

macOS/Linux:

```bash
chmod +x ./bin/clscli-linux-amd64
./bin/clscli-linux-amd64 --help
./bin/clscli-linux-amd64 topics --region ap-shanghai
```

If macOS blocks the downloaded binary:

```bash
xattr -d com.apple.quarantine ./bin/clscli-darwin-arm64
```

## Commands

```text
clscli topics   List log topics
clscli query    Search and analyze logs
clscli context  Get log context by PkgId/PkgLogId
```

Global flags:

```text
--region <region>        CLS region, required for API commands
--format json|csv        Output format, default csv
--output json|csv|path   Output format or file path
-o, --out <path>         Write output to file
```

## List Topics

Use this first when the topic ID or region is unknown.

```bash
clscli topics --region <region> [--topic-name name] [--logset-name name] [--logset-id id] [--limit 20] [--offset 0]
```

Examples:

```bash
clscli topics --region ap-shanghai --topic-name freeswitch --format csv
clscli topics --region ap-shanghai --logset-id <logset_id> --output topics.json
```

CSV columns:

```text
Region, TopicId, TopicName, LogsetId, CreateTime, StorageType
```

## Search Logs

Search a single topic:

```bash
clscli query --region <region> -t <topic_id> -q "UNALLOCATED_NUMBER" --last 1h --format json
```

Search with an absolute Unix millisecond time range:

```bash
clscli query --region <region> -t <topic_id> -q "UNALLOCATED_NUMBER" --from 1780452600000 --to 1780452900000 --limit 100 --format csv
```

Search multiple topics:

```bash
clscli query --region <region> --topics <topic_id_1>,<topic_id_2> -q "ERROR" --last 30m --format json
```

Do not combine `-t/--topic` with `--topics`.

Run SQL analysis:

```bash
clscli query --region <region> -t <topic_id> -q "UNALLOCATED_NUMBER | select count(*) as cnt" --last 1h --format json
```

Write output to files:

```bash
clscli query --region <region> -t <topic_id> -q "ERROR" --last 1h --output result.json
clscli query --region <region> -t <topic_id> -q "ERROR" --last 1h --output result.csv
```

Pagination:

```text
--limit <n>   Logs per request, default 100, max 1000
--max <n>     Positive value enables auto-pagination and caps total rows
--max 0       Single request
```

## Log Context

Use `query` first, then copy `PkgId`, `PkgLogId`, and `Time`.

```bash
clscli context <PkgId> <PkgLogId> --region <region> -t <topic_id> --btime <Time> --prev 10 --next 10 --format json
```

`--btime` is required. It accepts either:

```text
1780452731006
2026-06-03 10:12:11.006
```

Filter context logs:

```bash
clscli context <PkgId> <PkgLogId> --region <region> -t <topic_id> --btime 1780452731006 -q "UNALLOCATED_NUMBER" --format csv
```

## Query Syntax

Use CLS CQL for search conditions:

```text
ERROR
level:ERROR
UNALLOCATED_NUMBER
uuid:"68fccc99-8994-4d21-86f8-48828809beeb"
level:ERROR AND service:fs
```

Append SQL after `|` for analysis:

```text
* | select count(*) as cnt
UNALLOCATED_NUMBER | select count(*) as cnt
```

## Troubleshooting

Missing credentials:

```text
credentials required: set TENCENTCLOUD_SECRET_ID and TENCENTCLOUD_SECRET_KEY
```

Set both environment variables in the current shell.

Missing region:

```text
region required: use --region
```

Pass the correct CLS region for the topic.

Topic not found:

```text
ResourceNotFound.TopicNotExist
```

The topic is usually in another region. Run `topics` in likely regions or search by topic name.

Missing `BTime` for context:

```text
MissingParameter: BTime
```

Pass `--btime` from the query result `Time` field.

## Verification

The packaged Windows amd64 binary was verified on 2026-06-03 with:

```text
SKILL\bin\clscli-windows-amd64.exe
```

See `CLSCLI_TEST_REPORT.md` in the repository root for the full test matrix and outputs.
