/*
 * Unit tests for library package
 */
package library

import (
	"testing"
	"time"
)

func TestWriteLibraryFiles0(t *testing.T) {
	emptyLoader := EmptyLibraryLoader{"Empty"}
	dateUpdated := time.Now().Format("2006-01-02")
	lib := Library{
		Title: "Library",
		Summary: "Top level collection in the Library",
		DateUpdated: dateUpdated,
		TargetStatus: "public",
		Loader: emptyLoader,
	}
	WriteLibraryFiles(lib)
}

func TestWriteLibraryFiles1(t *testing.T) {
	mockLoader := MockLibraryLoader{"Mock"}
	dateUpdated := time.Now().Format("2006-01-02")
	lib := Library{
		Title: "Library",
		Summary: "Top level collection in the Library",
		DateUpdated: dateUpdated,
		TargetStatus: "public",
		Loader: mockLoader,
	}
	WriteLibraryFiles(lib)
}