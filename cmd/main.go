package main

import (
	"flag"
	"fmt"
	"github.com/chenxifun/auto-generate-code/auto"
	"github.com/chenxifun/auto-generate-code/common"
)

var inFilePath string
var outFilePath string
var dateFmt string
var versionPath string
var versionL int

func init() {
	flag.StringVar(&inFilePath, "i", "", "模板文件地址")
	flag.StringVar(&outFilePath, "o", "", "输出文件地址")
	flag.StringVar(&dateFmt, "d", "20060102150405", "时间格式")
	flag.StringVar(&versionPath, "v", "version.conf", "版本")
	flag.IntVar(&versionL, "l", 99, "版本自动增加列")
}

func main() {
	flag.Parse()

	if inFilePath == "" || outFilePath == "" {
		fmt.Println("can not be emtry")
		return
	}

	newVersion, err := auto.AutoVersion(versionL, versionPath)
	if err != nil {
		fmt.Println("auto version err ,", err)
		return
	}

	data := auto.MapData(dateFmt, newVersion)

	tmp, err := common.ReadFile(inFilePath)
	if err != nil {
		fmt.Println("read tmp err ,", err)
		return
	}

	code := auto.ConvertTmp(string(tmp), data)

	err = common.WriteFile([]byte(code), outFilePath, true)
	if err != nil {
		fmt.Println("write code err ,", err)
		return
	}

	err = common.WriteFile([]byte(newVersion), versionPath, true)
	if err != nil {
		fmt.Println("write version err ,", err)
		return
	}

}
