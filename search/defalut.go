package search
// 搜索数据用的默认匹配器


// defalutMatcher实现了默认路由器
type defalutMatcher struct{

}

func init(){
	var matcher defalutMatcher
	Register("default", matcher)
}

// Search 实现了默认匹配器的行为
// 这里创建的非指针，是因为defalutMatcher 没有
// 内存消耗，指针接受者
func (m defalutMatcher) Search(feed *Feed, searchTerm string) ([]*Result,error){
	return nil, nil
}