package search
// 执行搜索的主控制逻辑

import(
	"log"
	"sync"
	// "fmt"
	// "goinaction/rss"
)

// 注册匹配器映射，类似于注册路由的方式
var matchers = make(map[string]Matcher)

// 执行搜索逻辑
func Run(searchTerm string){
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil{
		log.Fatal(err)
	}

	// 创建一个无缓冲的通道，接受匹配后的结果
	results := make(chan *Result)
	
	// 构造一个waitGroup，以便处理所有的数据源
	var waitGroup sync.WaitGroup

	// 设置需要等待处理
	// 每个数据源的goroutine数量
	waitGroup.Add(len(feeds))
	// fmt.Println(feeds)

	// 为每个数据源启动一个goroutine来查找结果
	for _, feed := range feeds{
		// 获取一个匹配器用于查找
		// 这个架构的关键是使用一个接口类型来匹配并执行具有特定实现的匹配器
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["defalut"]
		}

		// 启动一个goroutine执行搜索
		// 闭包=函数+访问的外部变量
		// 这个闭包的设计思想：
		// 对于这个匿名函数来说，searchTerm, results，waitGroup就是
		// 其访问的外部变量，其改变相当于原变量的操作
		// 因为 matcher 和 feed 变量每次调用时值不相同，
		// 所以并没有使用闭包的方式访问这两个变量
		go func(matcher Matcher, feed *Feed){
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个goroutine来监控所有工作是否做完了
	go func(){
		// 等待所有任务完成
		waitGroup.Wait()

		// 关闭通道，通知Display函数
		// 内置的close函数，作用是关闭通道，会导致协程终止
		close(results)
	}()

	// 启动函数，显示返回的结果
	// 并且在最后一个结果显示完后返回
	// 所有 results 通道里的数据被处理之前，Display 函数不会返回
	Display(results)
}

func Register(feedType string, matcher Matcher){
	if _, exists := matchers[feedType]; exists{
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}

