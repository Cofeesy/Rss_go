package matchers
// 搜索rss源的匹配器

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"goinaction/search"
)

type (
	// item 根据 item 字段的标签，将定义的字段
    // 与 rss 文档的字段关联起来 
	item struct {
		 XMLName xml.Name `xml:"item"`
		 PubDate string `xml:"pubDate"`
		 Title string `xml:"title"`
		 Description string `xml:"description"`
		 Link string `xml:"link"`
		 GUID string `xml:"guid"`
		 GeoRssPoint string `xml:"georss:point"`
		}

	image struct{
		XMLName xml.Name `xml:"image"`
		URL string `xml:"url"`
		Title string `xml:"title"`
		Link string `xml:"link"`
	}
	channel struct{
		XMLName xml.Name `xml:"channel"`
		Title string `xml:"title"`
		Description string `xml:"description"`
		Link string `xml:"link"`
		PubDate string `xml:"pubDate"`
		LastBuildDate string `xml:"lastBuildDate"`
		TTL string `xml:"ttl"`
		Language string `xml:"language"`
		ManagingEditor string `xml:"managingEditor"`
		WebMaster string `xml:"webMaster"`
		Image image `xml:"image"`
		Item []item `xml:"item"`
	}
	rssDocument struct{
		XMLName xml.Name `xml:"rss"`
		Channel channel `xml:"channel"`
	}
)

type rssMatcher struct{}

func init()  {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// 当调用者为接口时
// 使用指针作为接收者声明的方法，只能在接口类型的值是一个指针的时候被调用。
// 使用值作为接收者声明的方法，在接口类型的值为值或者指针时，都可以被调用。
func (m rssMatcher) retrive(feed *search.Feed)(*rssDocument, error){
	if feed.URI == " "{
		return nil, errors.New("NO rss feed URL provided")
	}
	// 从网络源获得rss数据源文档
	resp, err := http.Get(feed.URI)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != 200{
		return nil, fmt.Errorf("HTTP Response Error %d", resp.StatusCode)
	}

	// 将rss数据源文档解码到代码定义的结构类型里
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}

// 同一个包（package）里面的所有 .go 文件中的任何顶级（top-level）变量、常量、函数、类型等，
// 都是相互完全可见的，无论它们的文件名是什么，也无论它们的标识符是首字母大写还是小写
func (m rssMatcher) Search(feed * search.Feed, searchTerm string) ([]*search.Result, error){
	var results []*search.Result
	log.Printf("Search Feed Type[%s] Site[%s] For Uri[%s]\n",feed.Type, feed.Name, feed.URI)

	// 获取要搜索的数据
	document, err := m.retrive(feed)
	if err !=nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item{
		// 检查标题部分是否包含搜索项
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil{
			return nil, err
		}

		// 如果找到匹配的项，将其作为结果保存
		if matched {
			results = append(results, &search.Result{
				Field: "Title",
				Content: channelItem.Title,
			})
		}

		// 检查描述部分是否包含搜索项
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}
		// 如果找到匹配的项，将其作为结果保存
		if matched {
			results = append(results, &search.Result{
				Field: "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}
		
