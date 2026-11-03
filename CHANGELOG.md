# Changelog

- 2026-11-03: v0.9.2: Every request logs one stderr line, `http <method> <path> -> <status> <duration>ms` (x360-common v0.5.1 `httplog`): no query, headers, cookies, tokens or bodies; `/healthz` skipped.

- 2026-11-02: v0.9.1: deployment advertises UpdateUriRoot1 and SOMRevision (x360-som online manifest); no code change.

- 2026-10-19: configure the native profile image hostname and port.

- 2026-10-19: configure the dashboard Home content root and locale-aware manifest path.

- 2026-10-19: deployment advertises HTTPS progress/profile URLs; XAM omits REST Authorization
  when LiveHive advertises HTTP.

- 2026-10-18: v0.9.0 split out of x360-kdc v0.8.14 into its own repository and image (`ghcr.io/openpak/x360-liveinfo`),
  history kept; shared packages now come from `x360-common`. Earlier entries: the x360-kdc CHANGELOG.
