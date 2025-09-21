package search

// 用于支持不同匹配器的接口
import (
	"fmt"
	"log"
)

// 保存搜索的结果
type Result struct{
	Field string
	Content string
}

// 如果接口类型只包含一个方法，那么这
// 个类型的名字以 er 结尾（接口命名规范）
type Matcher interface{
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}


// 并发的执行搜索
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result){
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil{
		log.Println("match.Match:", err)
		return
	}
	
	// 将结果写进通道
	for _, result := range searchResults{
		results <- result
	}

}

func Display(results chan *Result){
	// 通道会一直被阻塞，直到有结果写入
	// 一旦通道被关闭，for 循环就会终止
	for result := range results{
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}