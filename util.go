package gocommon

import "time"

// 创建一个定时器，到时 f 会被调用。
// 返回一个chan,如果需要取消定时器，则可调用 close(c)
func TimeOut(d time.Duration, f func()) chan int {
	t := time.NewTimer(d)
	over := make(chan int, 1)
	go func() {
		for {
			select {
			case <-t.C:
				{
					f()
				}
			case <-over:
				{
					t.Stop()
					return
				}
			}
		}
	}()
	return over
}
