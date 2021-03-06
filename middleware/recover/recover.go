/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-25 17:50:05
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 11:05:11
 */

/**
* @Description: 错误信息处理
* @Author: guoyzh
* @Date: 2019/10/23
 */

package recover

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"sync"

	"github.com/joshua-chen/go-commons/exception"
	"github.com/joshua-chen/go-commons/mvc/context/response"
	"github.com/joshua-chen/go-commons/validator"
	"github.com/kataras/iris/v12/context"

)

type Recover struct {
}

var (
	instance *Recover
	lock     *sync.Mutex = &sync.Mutex{}
)

func Instance() *Recover {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Recover{}
		}
	}
	return instance
}
func (a *Recover) New() context.Handler {

	return New()
}

func New() context.Handler {
	return func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() {
					return
				}

				var stacktrace string
				for i := 1; ; i++ {
					_, f, l, got := runtime.Caller(i)
					if !got {
						break
					}
					stacktrace += fmt.Sprintf("%s:%d\n", f, l)
				}

				excep := exception.Singleton()
				errCode := response.StatusInternalServerError

				if reflect.DeepEqual(excep.Err, err) {
					errCode = excep.Code					
				} else {
					vd := validator.Singleton()
					if reflect.DeepEqual(vd.Err, err) {
						errCode = vd.Code
					}
				}
				
				if len(strconv.Itoa(errCode)) < 5 {
					errCode = errCode * response.StatusCoefficient
				}
				errMsg := fmt.Sprintf("错误信息: %s", err)
				// when stack finishes
				logMessage := fmt.Sprintf("从错误中恢复：('%s')\n", ctx.HandlerName())
				logMessage += errMsg + "\n"
				logMessage += fmt.Sprintf("\n%s", stacktrace)
				// 打印错误日志
				ctx.Application().Logger().Warn(logMessage)
				// 返回错误信息

				result := response.NewErrorResult(errCode, errMsg)
				ctx.JSON(result)
				ctx.StopExecution()
			}
		}()

		ctx.Next()
	}
}
