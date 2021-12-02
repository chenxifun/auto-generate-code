package auto

import (
	"fmt"
	"github.com/chenxifun/auto-generate-code/common"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

func ConvertTmp(tmp string, data map[string]string) string {

	for k, v := range data {
		tmp = strings.ReplaceAll(tmp, fmt.Sprintf("{{%s}}", k), v)
	}

	return tmp
}

func MapData(dateFmt string, version string) map[string]string {

	data := make(map[string]string)
	data["date"] = time.Now().Format(dateFmt)
	data["version"] = version

	return data
}

func AutoVersion(l int, path string) (string, error) {

	version, err := common.ReadFile(path)
	if err != nil {
		return "", errors.WithMessage(err, "read version has error")
	}

	return convertVersion(l, string(version)), err
}

func convertVersion(l int, version string) string {

	vs := strings.Split(version, ".")
	if len(vs) <= l {
		l = len(vs) - 1
	}
	vs[l] = autoStr(vs[l])
	return strings.Join(vs, ".")
}

func autoStr(si string) string {
	vb, err := strconv.Atoi(si)
	if err != nil {
		vb = 0
	}

	vb++

	return strconv.Itoa(vb)

}
