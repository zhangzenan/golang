package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

// 👉 全局错误监控器
type PanicMonitor struct {
	mu      sync.Mutex
	panicCh chan interface{}
	done    chan struct{}
}

var GlobalPanicMonitor = &PanicMonitor{
	panicCh: make(chan interface{}, 10),
	done:    make(chan struct{}),
}

// 启动全局监控
func (pm *PanicMonitor) Start() {
	go func() {
		for {
			select {
			case p := <-pm.panicCh:
				log.Printf("💥 全局捕获 panic: %v", p)
			case <-pm.done:
				return
			}
		}
	}()
}

// 停止监控
func (pm *PanicMonitor) Stop() {
	close(pm.done)
}

// 注册 panic
func (pm *PanicMonitor) Report(panicVal interface{}) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	select {
	case pm.panicCh <- panicVal:
	default:
		log.Printf("⚠️ Panic channel full, dropped: %v", panicVal)
	}
}

type NodeComponent struct {
	name  string
	delay time.Duration
}

func (d *NodeComponent) Name() string {
	return d.name
}

func (c *NodeComponent) Execute(ctx context.Context, input []int) ([]int, error) {
	fmt.Println("NodeComponent:", c.name)

	data, err := c.doProcess(input)
	return data, err
}

// 👇 新增：纯业务逻辑方法
func (c *NodeComponent) doProcess(input []int) ([]int, error) {
	result := []int{}
	for _, v := range input {
		result = append(result, v*2)
	}
	if c.Name() == "D" {
		time.Sleep(time.Second * 3)
		/*a := 0
		b := 5 / a
		fmt.Println(b)*/
	}
	fmt.Println(c.name, "执行完成")
	return result, nil
}

type Result struct {
	Name string
	Data []int
	Err  error
}

func RunComponents(parentCtx context.Context, components []NodeComponent, input []int, timeout time.Duration) []Result {
	var wg sync.WaitGroup
	resultCh := make(chan Result, len(components))

	for _, cmp := range components {

		wg.Add(1)

		go func(c NodeComponent) {
			defer wg.Done()

			// 👉 每个组件独立超时（关键🔥）
			ctx, cancel := context.WithTimeout(parentCtx, timeout)
			defer cancel()

			doneCh := make(chan struct {
				data []int
				err  error
			}, 1)

			go func() {
				data, err := c.Execute(ctx, input)
				doneCh <- struct {
					data []int
					err  error
				}{data, err}
			}()

			var res Result
			// 👉 不阻塞写入（防止泄漏）
			select {
			case <-ctx.Done():
				fmt.Println(c.Name(), "超时/取消")
				res = Result{
					Name: c.Name(),
					Data: nil,
					Err:  ctx.Err(),
				}
			case result := <-doneCh:
				res = Result{
					Name: c.Name(),
					Data: result.data,
					Err:  result.err,
				}
			}
			// 👉 统一发送，只需一次 select
			select {
			case resultCh <- res:
			case <-parentCtx.Done():
			}

		}(cmp)
	}
	// 关闭 channel
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// ===== 收集结果 =====
	var results []Result
	for res := range resultCh {
		results = append(results, res)
	}

	return results
}

func main() {
	// 👉 启动全局 panic 监控
	GlobalPanicMonitor.Start()
	defer GlobalPanicMonitor.Stop()

	defer func() {
		if r := recover(); r != nil {
			log.Printf("💥 Main函数最终防线：%v", r)
			os.Exit(1) // 如果 main 崩溃，退出程序
		}
	}()
	components := []NodeComponent{
		{name: "A", delay: 2 * time.Second},
		{name: "B", delay: 4 * time.Second},
		{name: "C", delay: 4 * time.Second},
		{name: "D", delay: 2 * time.Second},
		{name: "E", delay: 2 * time.Second},
	}

	result := RunComponents(context.Background(), components, []int{1, 2, 3}, 2*time.Second)
	for a, i := range result {
		fmt.Println(a, i.Name, i.Data, i.Err)
	}

	for {
		time.Sleep(1 * time.Second)
	}
}
