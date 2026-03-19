package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type NodeComponent struct {
	name  string
	delay time.Duration
}

func (d *NodeComponent) Name() string {
	return d.name
}

func (c *NodeComponent) Execute(ctx context.Context, input []int) ([]int, error) {
	fmt.Println("NodeComponent:", c.name)
	// 👉 使用 goroutine 包装不可中断的操作
	resultCh := make(chan struct {
		data []int
		err  error
	}, 1)
	go func() {
		// 👉 添加 recover 保护
		defer func() {
			if r := recover(); r != nil {
				resultCh <- struct {
					data []int
					err  error
				}{nil, fmt.Errorf("panic recovered: %v", r)}
			}
		}()
		data, err := c.doProcess(input)
		resultCh <- struct {
			data []int
			err  error
		}{data, err}
	}()
	select {
	case <-ctx.Done():
		fmt.Println(c.name, "超时/取消")
		return nil, ctx.Err()
	case result := <-resultCh:
		return result.data, result.err
	}
}

// 👇 新增：纯业务逻辑方法
func (c *NodeComponent) doProcess(input []int) ([]int, error) {
	result := []int{}
	for _, v := range input {
		result = append(result, v*2)
	}
	if c.Name() == "D" {
		//time.Sleep(time.Second * 1)
		a := 0
		b := 5 / a
		fmt.Println(b)
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

			data, err := c.Execute(ctx, input)

			// 👉 不阻塞写入（防止泄漏）
			select {
			case resultCh <- Result{
				Name: c.Name(),
				Data: data,
				Err:  err,
			}:
			case <-parentCtx.Done():
				return
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
