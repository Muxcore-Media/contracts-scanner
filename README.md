# contracts-scanner

Protobuf/gRPC contract interfaces for media library scanner modules in MuxCore.

## Proto Service

- **`ScannerService`** (`proto/muxcore/scanner/v1/scanner.proto`)
  - `Scan` / `ScanLibraryRoots` — scan watch dirs or library roots for new media
  - `GetStats` — scanner statistics and last scan status
  - `AddWatchDir` / `RemoveWatchDir` / `ListWatchDirs` — manage import watch directories
  - `ListImported` — paginated imported file history
  - `ImportPath` — import a specific path on demand
  - `ListImportCandidates` — list pending import candidates

Proto source was migrated from `media-scanner/proto/scannerv1/`.

## Go import

Generated stubs: `github.com/Muxcore-Media/contracts-scanner/muxcore/scanner/v1` (checked in; run `make proto` after `.proto` edits).

## Implementing Modules

- [media-scanner](https://github.com/Muxcore-Media/media-scanner) — file watcher and library import
