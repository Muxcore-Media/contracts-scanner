# Changelog

## [0.2.0] — 2026-08-31

### Added
- `MediaType` enum (`movie`, `tv`, `both`, `music`) and music fields on `ImportedFile` / `ImportPathRequest`.
- Dual library roots on watch dirs: `tv_library_path`, `music_library_path` on `WatchDir` / `AddWatchDirRequest`.
- `UpdateWatchDir` and `SetWatchDirEnabled` RPCs.
- Parse-preview fields on `ImportCandidate` (`title`, `year`, `media_type`, season/episode, `artist`, `album`).
- Per-file `ImportPathResult` on `ImportPathResponse` (`original_path`, `destination_path`, `status`, `error`).
- `ScanRequest` / `ScanLibraryRootsRequest`: `dry_run`, optional `path`, `media_type` filter.
- `ListImportedRequest.status` filter; `ImportedFile.error_message` for failed imports.
- Last-scan counters and `last_error` on `GetStatsResponse`.
- `contract_test.go` freezing RPC names and field numbers.
- Forgejo CI `proto-check` job.

### Fixed
- Makefile `clean` no longer deletes committed stubs under `muxcore/scanner/v1/`.

### Compatibility
- Additive **v1** proto changes; implementers and callers may adopt new fields incrementally.
- Capability **`scanner`** (implementers) vs **`contracts.scanner`** (catalog).

## [0.1.0] — 2026-08-08

### Added
- Initial `ScannerService` contract (`Scan`, `ScanLibraryRoots`, `GetStats`, watch-dir CRUD, `ListImported`, `ImportPath`, `ListImportCandidates`).
- Generated Go stubs under `muxcore/scanner/v1`.

### Compatibility
- Contract interface version **v1**; implementers advertise capability `scanner`.
- Consumers: `media-scanner`, `media-automation`, `admin-ui`.
