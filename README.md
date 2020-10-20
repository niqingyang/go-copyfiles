## go-copyfiles

用 Go 实现的复制文件，可用于 SourceTree 自定义操作导出选中的文件，支持通过文件选择器来选中导出文件的目标目录。

## Usage

```bash
copyfiles -h
Usage:
  copyfiles [OPTIONS]

Application Options:
      --dir=     the path of the copy files parent dir
      --files=   the files for copy
      --dest=    the destination directory of the copy files
  -v, --version  the application version info

Help Options:
  -h, --help     Show this help message
```

## SourceTree

![](https://raw.githubusercontent.com/niqingyang/go-copyfiles/master/images/01.png)