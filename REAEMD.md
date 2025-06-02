# 二维码扫描工具

提供一个二维码文件扫描函数，支持扫描 `png`和`jpeg`文件。

```
	func ScanFile(name string) (s string, e error)
	参数 name 是文件名；
	返回值 s 是识别出的字符串；
	返回值 e 是错误值。
```