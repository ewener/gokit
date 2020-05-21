package gokit

import (
	bytes2 "bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Sign(structName interface{}, salt string, excludeNames []string) string {
	j := Sort(structName, excludeNames)
	return SignWithSalt(salt, j)
}

func SignWithSalt(salt string, param string) string {
	return Md5String(param + salt)
}

func Md5String(param string) string {
	return Md5([]byte(param))
}

func Md5(param []byte) string {
	hash := md5.New()
	hash.Write(param)
	sum := strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
	return sum
}
func Sha1(unique string, prefix string) []byte {
	h := sha1.New()
	h.Write([]byte(unique))
	return h.Sum([]byte(prefix))
}
func Sha256(unique string, prefix string) []byte {
	h := sha256.New()
	h.Write([]byte(unique))
	return h.Sum([]byte(prefix))
}
func Sha512(unique string, prefix string) []byte {
	h := sha512.New()
	h.Write([]byte(unique))
	return h.Sum([]byte(prefix))
}

func Buff(p interface{}) (bytes2.Buffer, error) {
	var buff bytes2.Buffer
	bytes, err := json.Marshal(p)
	if err != nil {
		return buff, err
	}
	buff.Write(bytes)
	return buff, nil
}

func ToJson(input interface{}) (string, error) {
	j, e := json.Marshal(input)
	if e != nil {
		return "", e
	}
	return string(j), nil
}

func StructToMap(structs interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(structs)
	d := json.NewDecoder(bytes2.NewBuffer(j))
	// go int64 json序列化精度会丢失，应该使用UseNumber 反序列化
	d.UseNumber()
	_ = d.Decode(&m)
	return m
}

func Sort(structs interface{}, excludeNames []string) string {
	sort.Strings(excludeNames)
	m := StructToMap(structs)
	keys := make([]string, 0, len(m))
	for k := range m {
		i := sort.Search(len(excludeNames), func(i int) bool {
			return excludeNames[i] >= k
		})
		if i < len(excludeNames) && excludeNames[i] == k {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	m2 := make(map[string]interface{}, len(keys))
	for _, k := range keys {
		m2[k] = m[k]
	}
	j, _ := ToJson(m2)
	return j
}

// 解决默认浮点类型json marshal之后，不会保留为0的小数位，如：23.00->23
type Float64 float64

func (n Float64) MarshalJSON() ([]byte, error) {
	// 保留两位小数
	return []byte(fmt.Sprintf("%.2f", n)), nil
}

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// 生成随机字符串
func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// 生成32位md5字串
func GetMd5String(s string, upper bool, half bool) string {
	h := md5.New()
	h.Write([]byte(s))
	result := hex.EncodeToString(h.Sum(nil))
	if upper == true {
		result = strings.ToUpper(result)
	}
	if half == true {
		result = result[8:24]
	}
	return result
}

// 操作系统位数
func Bit() string {
	bit := 32 << (^uint(0) >> 63)
	return strconv.Itoa(bit)
}

func AddEnv(key, value string) error {
	v := value
	if env, b := os.LookupEnv(key); b {
		var delimiter string
		switch runtime.GOOS {
		case "windows":
			delimiter = ";"
		case "linux", "darwin":
			delimiter = ":"
		default:
			return errors.New("illegal delimiter")
		}
		v = env + delimiter + value
	}
	return os.Setenv(key, v)
}

func WithinDays(now, compare time.Time, differ int) (in bool) {
	y, m, d := now.Date()
	y2, m2, d2 := compare.Date()
	if y == y2 && m == m2 && d-d2 <= differ {
		in = true
	}
	return
}

func FloatDataParse(data float64) float64 {
	dataForShow, err := strconv.ParseFloat(strconv.FormatFloat(data, 'f', 2, 32), 64)
	if err != nil {
		return 0.00
	}
	return dataForShow
}
