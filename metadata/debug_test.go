package main

import (
	"fmt"
	"testing"

	"github.com/dsoprea/go-exif/v3"
	common "github.com/dsoprea/go-exif/v3/common"
)

func TestDebugCollect(t *testing.T) {
	tiff := buildExifTiff()
	for i, byt := range tiff[:12] {
		fmt.Printf("  tiff[%d]=0x%02x\n", i, byt)
	}
	im, err1 := common.NewIfdMappingWithStandard()
	ti := exif.NewTagIndex()
	eh, index, err2 := exif.Collect(im, ti, tiff)
	fmt.Printf("DEBUG err1=%v err2=%v FirstIfdOffset=%d\n", err1, err2, eh.FirstIfdOffset)
	if index.RootIfd != nil {
		for _, e := range index.RootIfd.Entries() {
			fmt.Printf("DEBUG entry tag=0x%04x name=%q\n", e.TagId(), e.TagName())
		}
		if v := valorTag(index.RootIfd, 0x010F); v != "" {
			fmt.Printf("DEBUG Make=%q\n", v)
		} else {
			fmt.Printf("DEBUG Make vacio\n")
		}
	} else {
		fmt.Printf("DEBUG RootIfd es nil\n")
	}
}
