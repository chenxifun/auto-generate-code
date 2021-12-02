# auto-generate-code

可以根据模版自动生成一个带自增版本号和编译时间的go文件  
可以在代码中配合`go generate`使用  
在代码中增加 
```
//go:generate autoCode -i version.tmp -o version.go -l 3
```
在目录中执行 `go generate` 
生成类似下面的`version.go`  
```
package version

const Version = "2.0.0.3-20211202145947"

```