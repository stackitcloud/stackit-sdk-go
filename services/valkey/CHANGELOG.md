## v0.1.0

- New service `valkey`:
  Valkey is a fork of redis under the Linux Foundation. The API in this SDK is mostly compatible with v1.2.0 of the
  `redis` service. Differences as follows:
- no unversioned legacy SDK, just `v1api` and `v2api`
- removal of deprecated `TLS` related attributes and models on model `Instance`
- new attributes: `MinReplicasToWrite` and `ReplBacklogSize` on model `Instance`
- `v1api`
    - `valkey/v1api` changed the return type of the `ListBackups` operation from `*ListbackupsResponse` to `[]Backup`
      when compared to `redis/v1api`. This change was only done for `redis/v2api`.
    - `wait` remove deprecated constants like `wait.INSTANCESTATUS_ACTIVE` when compared to `redis`