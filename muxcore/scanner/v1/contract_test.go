package scannerv1

import (
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestScannerServiceRPCNamesStable(t *testing.T) {
	want := map[string]bool{
		"Scan":                 true,
		"ScanLibraryRoots":     true,
		"GetStats":             true,
		"AddWatchDir":          true,
		"UpdateWatchDir":       true,
		"SetWatchDirEnabled":   true,
		"RemoveWatchDir":       true,
		"ListWatchDirs":        true,
		"ListImported":         true,
		"ImportPath":           true,
		"ListImportCandidates": true,
	}
	got := make(map[string]bool, len(ScannerService_ServiceDesc.Methods))
	for _, m := range ScannerService_ServiceDesc.Methods {
		got[m.MethodName] = true
	}
	for name := range want {
		if !got[name] {
			t.Errorf("missing RPC %s", name)
		}
	}
}

func TestWatchDirFieldNumbers(t *testing.T) {
	want := map[string]protoreflect.FieldNumber{
		"id":                 1,
		"path":               2,
		"media_type":         3,
		"library_path":       4,
		"enabled":            5,
		"created_at":         6,
		"tv_library_path":    7,
		"music_library_path": 8,
	}
	assertFieldNumbers(t, (&WatchDir{}).ProtoReflect().Descriptor(), want)
}

func TestImportCandidateFieldNumbers(t *testing.T) {
	want := map[string]protoreflect.FieldNumber{
		"path":           1,
		"name":           2,
		"size":           3,
		"watch_dir_id":   4,
		"title":          5,
		"year":           6,
		"media_type":     7,
		"season_number":  8,
		"episode_number": 9,
		"artist":         10,
		"album":          11,
	}
	assertFieldNumbers(t, (&ImportCandidate{}).ProtoReflect().Descriptor(), want)
}

func TestImportPathResultFieldNumbers(t *testing.T) {
	want := map[string]protoreflect.FieldNumber{
		"original_path":    1,
		"destination_path": 2,
		"status":           3,
		"error":            4,
	}
	assertFieldNumbers(t, (&ImportPathResult{}).ProtoReflect().Descriptor(), want)
}

func TestImportedFileFieldNumbers(t *testing.T) {
	want := map[string]protoreflect.FieldNumber{
		"id":               1,
		"original_path":    2,
		"destination_path": 3,
		"file_name":        4,
		"media_type":       5,
		"title":            6,
		"year":             7,
		"season_number":    8,
		"episode_number":   9,
		"quality":          10,
		"tmdb_id":          11,
		"imported_at":      12,
		"status":           13,
		"error_message":    14,
		"artist":           15,
		"album":            16,
		"track_number":     17,
	}
	assertFieldNumbers(t, (&ImportedFile{}).ProtoReflect().Descriptor(), want)
}

func TestGetStatsResponseFieldNumbers(t *testing.T) {
	want := map[string]protoreflect.FieldNumber{
		"total_imported":            1,
		"watch_dirs":                2,
		"last_scan_at":              3,
		"last_scan_status":          4,
		"last_scan_files_found":     5,
		"last_scan_files_imported":  6,
		"last_scan_files_skipped":   7,
		"last_error":                8,
	}
	assertFieldNumbers(t, (&GetStatsResponse{}).ProtoReflect().Descriptor(), want)
}

func TestMediaTypeVocabulary(t *testing.T) {
	cases := []MediaType{
		MediaType_MEDIA_TYPE_MOVIE,
		MediaType_MEDIA_TYPE_TV,
		MediaType_MEDIA_TYPE_BOTH,
		MediaType_MEDIA_TYPE_MUSIC,
	}
	for _, mt := range cases {
		if mt == MediaType_MEDIA_TYPE_UNSPECIFIED {
			t.Errorf("unexpected unspecified in vocabulary list")
		}
		if mt.String() == "" {
			t.Errorf("media type %v has empty String()", mt)
		}
	}
}

func TestImportPath(t *testing.T) {
	var _ grpc.ServiceDesc = ScannerService_ServiceDesc
}

func TestCapabilityIDs(t *testing.T) {
	const consumerCapability = "scanner"
	const contractCapability = "contracts.scanner"
	if consumerCapability == contractCapability {
		t.Fatal("consumer and contract capability ids must differ")
	}
}

func assertFieldNumbers(t *testing.T, desc protoreflect.MessageDescriptor, want map[string]protoreflect.FieldNumber) {
	t.Helper()
	for name, num := range want {
		fd := desc.Fields().ByName(protoreflect.Name(name))
		if fd == nil {
			t.Errorf("missing field %q", name)
			continue
		}
		if fd.Number() != num {
			t.Errorf("field %q number = %d, want %d", name, fd.Number(), num)
		}
	}
}
