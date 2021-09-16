package search

// defaultMatcher 实现了默认的匹配器： matcher
// 空结构创建实例时，不会分配任何内存。这种结构适合创建没有任何状态的类型。
// 默认匹配器，不需要维护任何状态，只要实现对应的接口就行。
type defaultMatcher struct{}

// init() 向程序注册默认的匹配器
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search 实现了默认匹配器的方法
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
