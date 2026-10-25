# Xbox 360 LiveInfo Server

OpenPak's Xbox 360 LiveInfo Server.

| | |
| --- | --- |
| Network | Xbox 360 (Live service 25) |
| Role | LiveHive endpoint: serves the `key=value;...` hive the dashboard reads for its service URLs and update revisions |
| Depends on | [`x360-common`](https://github.com/openpak/x360-common) (shared Go library, `replace` + committed `vendor/`) |
| Image | `ghcr.io/openpak/x360-liveinfo` |
| Licence | AGPL-3.0-only |

## What it serves

| Endpoint | What |
| --- | --- |
| `/liveinfo/GetLiveConfig.ashx` | The LiveHive `key=value;...` text the dashboard reads |
| `/liveinfo/getnetworkauthorizationlist.ashx` | Network authorization list, always sent empty |

## Build and test

```sh
TMPDIR=<scratch> X360_KDC_TESTKEYS=<dir with the dev keys> go test -mod=vendor ./...
```

## Docs

- [Protocol](docs/protocol.md)
- [Evidence](docs/evidence.md)
- [Deploy](docs/deploy.md)
- [Notes](docs/notes.md)

## Licence

AGPL-3.0-only. See [LICENSE](LICENSE).
