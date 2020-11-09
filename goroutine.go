package gokit

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"sync"
	"time"
)

// GoSafely wraps a `go func()` with recover()
func GoSafely(wg *sync.WaitGroup, ignoreRecover bool, handler func(), catchFunc func(r interface{})) {
	if wg != nil {
		wg.Add(1)
	}
	go func() {
		defer func() {
			if r := recover(); r != nil {
				if ignoreRecover {
					panicLog(fmt.Sprintf("goroutine panic: %v\n%s\n",
						r, string(debug.Stack())))
					os.Exit(1)
				}
				if catchFunc != nil {
					if wg != nil {
						wg.Add(1)
					}
					go func() {
						defer func() {
							if p := recover(); p != nil {
								if ignoreRecover {
									panicLog(fmt.Sprintf("recover goroutine panic:%v\n%s\n", p, string(debug.Stack())))
									os.Exit(1)
								}

							}

							if wg != nil {
								wg.Done()
							}
						}()
						catchFunc(r)
					}()
				}
			}
			if wg != nil {
				wg.Done()
			}
		}()
		handler()
	}()
}

func panicLog(info string) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return
	}
	now := time.Now()
	// todo 按日期分文件，避免过大
	logPath := filepath.Join(dir, "crash.log")
	file, er := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if er == nil {
		defer file.Close()
		file.WriteString(now.Format("2006-01-02 15:04:05") + "\r\n")
		file.WriteString(fmt.Sprintf("%s\r\n", info))
		file.WriteString("========\r\n")
	}
}
