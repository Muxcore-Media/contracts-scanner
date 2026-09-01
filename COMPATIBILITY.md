# Compatibility

| Component | Requirement |
|-----------|-------------|
| Contract interface | `ScannerService` **v1** |
| Go module path | `github.com/Muxcore-Media/contracts-scanner` |
| Generated package | `github.com/Muxcore-Media/contracts-scanner/muxcore/scanner/v1` |
| MuxCore core | ≥ 0.4.0 (consumers) |
| Capability ID (contract repo) | `contracts.scanner` |
| Capability ID (implementers / callers) | `scanner` |

Implementers register capability **`scanner`** in `muxcore.json` and implement gRPC service **`ScannerService`**. The contract repository capability id **`contracts.scanner`** identifies this proto module in the umbrella catalog.

## Media type vocabulary

`media_type` string fields accept lowercase tokens (see `MediaType` enum in proto):

| Token | Meaning |
|-------|---------|
| `movie` | Feature film library |
| `tv` | TV series library |
| `both` | Movies and TV from one watch dir |
| `music` | Audio library |

## Dual library roots

Watch directories may target separate destination roots:

| Field | Use |
|-------|-----|
| `library_path` | Default / movie destination |
| `tv_library_path` | TV destination when `media_type` is `tv` or `both` |
| `music_library_path` | Music destination when `media_type` is `music` or `both` |

Implementers should persist all three; callers use `UpdateWatchDir` to change destinations without re-adding.

## Import path results

`ImportPathResponse.results` carries per-file outcomes. Aggregate counts (`files_found`, `files_imported`, `files_skipped`) remain for backward-compatible summaries.

| `status` | Meaning |
|----------|---------|
| `imported` | File moved/copied/hardlinked to library |
| `skipped` | Already imported, sample/extra, or policy skip |
| `failed` | Error; see `error` |

Paths may be local filesystem paths or `storage://` URIs when the implementer supports mesh storage.

## Imported file status

| `status` | Meaning |
|----------|---------|
| `imported` | Successfully organized |
| `failed` | Import attempted but failed; see `error_message` |
| `skipped` | Recorded skip (optional implementer use) |

`ListImportedRequest.status` filters on this column.

## Scan options

| Field | Behavior |
|-------|----------|
| `dry_run` | Parse and classify without writing files |
| `path` | Restrict scan to one watch dir or file |
| `media_type` | Filter by kind during scan |

## Stats / scan log

`GetStatsResponse` exposes the latest `scan_log` row: `last_scan_status`, `last_scan_files_*`, and `last_error` when the implementer records failures.

## Domain events

Implementers **should** publish `file.imported` and `import.failed` via `contracts-media/events` after successful or failed imports. RPC success alone is not sufficient for automation (`media-automation`).

Breaking proto changes require a new major contract version and coordinated consumer updates.
