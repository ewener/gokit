package gokit

import (
	bytes2 "bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"dpf.blueguard/logger"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	st "go.bug.st/serial"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)


func Rs(data interface{}, err error, code ...string) FrontendResponse {
	return RsWith(new(Response), data, err, code...)
}

func RsWith(res *Response, data interface{}, err error, code ...string) FrontendResponse {
	if err == nil {
		res.Code = Success
		res.Data = data
	} else {
		res.Data = data
		res.Err = err.Error()
		if len(code) > 0 {
			res.Code = code[0]
		} else {
			res.Code = SystemError
		}
	}
	return StructToMap(res)
}

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

func LogPanic() {
	if err := recover(); err != nil {
		file, er := os.OpenFile(CrashLog, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if er == nil {
			defer file.Close()
			file.WriteString(time.Now().Format("2006-01-02 15:04:05") + "\r\n")
			file.WriteString(fmt.Sprintf("%v\r\n", err))
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			file.WriteString(fmt.Sprintf("%s\r\n", string(buf[:n])))
			file.WriteString("========\r\n")
		}
		os.Exit(1)
	}
}

// 记录goroutine异常日志
func SafetyRun(do func()) {
	go func() {
		defer LogPanic()
		do()
	}()
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

func DeepCopy(dst, src interface{}) error {
	var buf bytes2.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes2.NewBuffer(buf.Bytes())).Decode(dst)
}

func UserHomeDir() string {
	dir, _ := os.UserHomeDir()
	return dir
}

// 解决默认浮点类型json marshal之后，不会保留为0的小数位，如：23.00->23
type Float64 float64

func (n Float64) MarshalJSON() ([]byte, error) {
	// 保留两位小数
	return []byte(fmt.Sprintf("%.2f", n)), nil
}

type Exchanger interface {
	Sign() Exchanger   // 签名
	Exchange() error   // 交互
	VerifySign() error // 验签
}

var Dispatcher dispatcher

type Msg struct {
	// 消息类型
	MsgType MsgType
	// 消息数据
	Data interface{}
	// 发送目的地
	Dst interface{}
}

type dispatcher struct {
	processors map[MsgType]Processor
	sync.RWMutex
}

func (dis *dispatcher) Add(processor Processor) {
	if processor == nil {
		return
	}
	dis.Lock()
	defer dis.Unlock()
	if dis.processors == nil {
		dis.processors = make(map[MsgType]Processor)
	}
	dis.processors[processor.MsgType()] = processor
}

func (dis *dispatcher) Remove(msgType MsgType) {
	dis.Lock()
	defer dis.Unlock()
	delete(dis.processors, msgType)
}

func (dis *dispatcher) Dispatch(msg Msg) {
	if dis.processors != nil {
		dis.RLock()
		defer dis.RUnlock()
		if dis.processors != nil {
			processor := dis.processors[msg.MsgType]
			if processor != nil {
				processor.Process(msg)
			}
		}
	}
}

type MsgType uint8

const (
	Websocket MsgType = iota // websocket
)

type Processor interface {
	MsgType() MsgType
	Process(Msg)
}

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var RandomCodeLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var RequestIdLetters = []rune("0123456789")
var RequestIdLen = 6

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

// 生成清洗随机字符串
func GenCleanRandom(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = RandomCodeLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	rand.Seed(time.Now().Unix())
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

func GetSerialPorts() ([]string, error) {
	ports, err := st.GetPortsList()
	if err != nil {
		return nil, err
	}
	return ports, err
}

// 获取当前时间戳
func CurrentTime() int64 {
	timeStr := time.Now()
	timeNumber := timeStr.Unix()
	return timeNumber
}


func IsFileExist(fileName string) (error, bool) {
	_, err := os.Stat(fileName)
	if err == nil {
		return nil, true
	}
	if os.IsNotExist(err) {
		return nil, false
	}
	return err, false
}
