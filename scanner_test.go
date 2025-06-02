package goZbarScanner

import (
	"testing"
)

func TestScanFile(t *testing.T) {
	res, err := ScanFile("qr.png")
	if err != nil {
		t.Fatal(err)
	}
	if res != "https://mdc.html5.qq.com/directdown?app=qqbrowser&channel=1100125022" {
		t.Fatalf("Result not match:%v\n", res)
	}
}
