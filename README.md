# 二维码扫描工具

提供一个二维码文件扫描函数，支持扫描 `png`和`jpeg`文件。使用 `zbar` 实现扫描识别。

## `ubuntu24.04`上编译时先安装下面这个支持包：
```
	sudo apt install libzbar-dev
```

## 函数说明：

```
	func ScanFile(name string) (s string, e error)
	参数 name 是文件名；
	返回值 s 是识别出的字符串；
	返回值 e 是错误值。
```