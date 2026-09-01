# contracts-scanner

Protobuf/gRPC contract interfaces for media library scanner modules in MuxCore.

## Proto Service

- **`ScannerService`** (`proto/muxcore/scanner/v1/scanner.proto`)
  - `Scan` / `ScanLibraryRoots` — scan watch dirs or library roots (`dry_run`, optional `path`, `media_type` filter)
  - `GetStats` — scanner statistics, last scan status, and last-scan file counts
  - `AddWatchDir` / `UpdateWatchDir` / `SetWatchDirEnabled` / `RemoveWatchDir` / `ListWatchDirs` — manage import watch directories
  - `ListImported` — paginated imported file history (`status` filter, `error_message` on failures)
  - `ImportPath` — import a specific path on demand (local path or `storage://` URI); per-file `results`
  - `ListImportCandidates` — list pending import candidates with parse-preview metadata

### Dual library roots

Watch dirs may specify separate destinations:

- **`library_path`** — movies / default root
- **`tv_library_path`** — TV library when `media_type` is `tv` or `both`
- **`music_library_path`** — music library when `media_type` is `music` or `both`

### Media types

Canonical tokens: `movie`, `tv`, `both`, `music` (see `MediaType` enum). Music imports use `artist`, `album`, `track_number` on `ImportPathRequest` and `ImportedFile`.

Proto source was migrated from `media-scanner/proto/scannerv1/`.

## Version

**v0.2.0** — see [CHANGELOG.md](CHANGELOG.md) and [COMPATIBILITY.md](COMPATIBILITY.md).

**v0.1.0** — initial `ScannerService` with watch-dir CRUD, import history, and candidate listing.

## Go import

Generated stubs: `github.com/Muxcore-Media/contracts-scanner/muxcore/scanner/v1` (checked in; run `make proto` after `.proto` edits).

## Implementing Modules

- [media-scanner](https://github.com/Muxcore-Media/media-scanner) — file watcher and library import

## Regenerate protos

```bash
make proto
```

With Nix:

```bash
nix-shell -p protobuf go --run 'make proto'
```

CI runs `make proto` and fails if generated stubs drift.
