# x360-liveinfo

The Xbox 360 LiveHive endpoint: `/liveinfo/GetLiveConfig.ashx` (Live service 25, TCP 25), the `key=value;...` hive the dashboard reads
for its service URLs and update revisions.

Part of OpenPak's Xbox 360 stack: one repository and one image per service ([`x360-kdc`](../x360-kdc),
[`x360-gateway`](../x360-gateway), [`x360-xuacs`](../x360-xuacs), [`x360-liveinfo`](../x360-liveinfo),
[`x360-notice`](../x360-notice), [`x360-xexds`](../x360-xexds)) around the shared Go library
[`x360-common`](../x360-common) (`replace` + committed `vendor/`; after changing x360-common run
`go mod vendor` here and commit).


## Configuration

Environment `X360_LIVEINFO_*` (see `deploy/primary/openpak-x360-liveinfo.container` and `x360-common/config`): `ADDR`, `SMREVISION`, `POLL_MINUTES`, `EXTRA`, `DEBUG`.

## Build and test

`TMPDIR=<scratch> X360_KDC_TESTKEYS=<dir with the dev keys> go test -mod=vendor ./...`

Tag `vX.Y.Z` and CI publishes `ghcr.io/openpak/x360-liveinfo:latest` (amd64 + arm64); hosts auto-update `:latest`.
`deploy/` holds the live quadlet and helpers. History before 2026-10-18 is the x360-kdc monorepo's, filtered to
this service. Licensed AGPL-3.0, see LICENSE.
