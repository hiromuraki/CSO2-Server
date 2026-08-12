package encode

import (
	"bytes"
	"testing"
)

func TestGBKRoundTrip(t *testing.T) {
	InitConverter("gbk")
	src := "你好CSO2"
	// 验证 UTF-8 -> GBK 编码：'你'=C4E3 '好'=BAC3
	gbk, err := Utf8ToLocal(src)
	if err != nil {
		t.Fatal("Utf8ToLocal err:", err)
	}
	want := []byte{0xC4, 0xE3, 0xBA, 0xC3, 'C', 'S', 'O', '2'}
	if !bytes.Equal([]byte(gbk), want) {
		t.Fatalf("Utf8ToLocal mismatch: got % X, want % X", gbk, want)
	}
	// 验证 GBK -> UTF-8 解码
	back, err := LocalToUtf8(gbk)
	if err != nil {
		t.Fatal("LocalToUtf8 err:", err)
	}
	if back != src {
		t.Fatalf("round trip mismatch: got %q, want %q", back, src)
	}
	// 兼容旧函数
	b, err := GbkToUtf8(want)
	if err != nil || string(b) != src {
		t.Fatalf("GbkToUtf8 mismatch: %q %v", b, err)
	}
	b, err = Utf8ToGbk([]byte(src))
	if err != nil || !bytes.Equal(b, want) {
		t.Fatalf("Utf8ToGbk mismatch: % X %v", b, err)
	}
	// 非法字节不应崩溃
	if _, err = LocalToUtf8("\xff\xfe"); err != nil {
		t.Log("invalid input handled:", err)
	}
}

func TestBig5Init(t *testing.T) {
	if !InitConverter("big5") {
		t.Fatal("InitConverter(big5) failed")
	}
	s, err := Utf8ToLocal("中文")
	if err != nil || s == "" {
		t.Fatal("big5 encode err:", err)
	}
}
