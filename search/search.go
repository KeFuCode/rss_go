package search

import (
	"log"	
	"sync"
)


//matchers: 用 map 保存用于搜索的的已经注册的匹配器
var matchers = make(map[string]Matcher)

//Run 执行 search 逻辑
func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result)

	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个 goroutine 来监控是否所有的工作都做完了
	go func()  {
		// 等待所有任务完成
		waitGroup.Wait()

		// 用关闭通道的方式，通知Display函数
		// 可以退出程序了
		close(results)
	}()
	
	// 启动函数，显示返回的结果
	// 并且在最后一个结果显示完后返回
	Display(results)
}

//调用 register 来注册匹配器供程序使用。
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}