package encode

import (
	"bytes"
	"io"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

// 纯 Go 实现的编码转换，不依赖 cgo/iconv
// toLocal 负责 UTF-8 -> 本地编码，toUTF8 负责本地编码 -> UTF-8
var (
	toLocal encoding.Encoding
	toUTF8  encoding.Encoding
)

//InitConverter 初始化本地编码转换器
func InitConverter(local string) bool {
	switch strings.ToLower(local) {
	case "big5", "big-5", "zh-tw":
		toLocal, toUTF8 = traditionalchinese.Big5, traditionalchinese.Big5
	default: //gbk / gb2312 / gb18030 及其他配置一律按 GBK 处理
		toLocal, toUTF8 = simplifiedchinese.GBK, simplifiedchinese.GBK
	}
	return true
}

//GbkToUtf8 转换GBK编码到UTF-8编码
func GbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	return io.ReadAll(r)
}

//Utf8ToGbk 转换UTF-8编码到GBK编码
func Utf8ToGbk(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewEncoder())
	return io.ReadAll(r)
}

//Utf8ToLocal 转换UTF-8编码到本地编码
func Utf8ToLocal(str string) (b string, err error) {
	r := transform.NewReader(strings.NewReader(str), toLocal.NewEncoder())
	buf, err := io.ReadAll(r)
	return string(buf), err
}

//LocalToUtf8 转换本地编码到UTF-8编码
func LocalToUtf8(str string) (b string, err error) {
	r := transform.NewReader(strings.NewReader(str), toUTF8.NewDecoder())
	buf, err := io.ReadAll(r)
	return string(buf), err
}
