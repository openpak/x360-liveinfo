# Deploy: x360-liveinfo

## Configuration

Environment `X360_LIVEINFO_*` (see `deploy/primary/openpak-x360-liveinfo.container` and `x360-common/config`): `ADDR`, `SMREVISION`, `POLL_MINUTES`, `EXTRA`, `DEBUG`.

## Build and test

`TMPDIR=<scratch> X360_KDC_TESTKEYS=<dir with the dev keys> go test -mod=vendor ./...`

Tag `vX.Y.Z` and CI publishes `ghcr.io/openpak/x360-liveinfo:latest` (amd64 + arm64); hosts auto-update `:latest`.
`deploy/` holds the live quadlet and helpers. History before 2026-10-18 is the x360-kdc monorepo's, filtered to
this service. Licensed AGPL-3.0, see LICENSE.
