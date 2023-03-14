package gedis

// Iterator 迭代器
type Iterator struct {
	Data  []interface{}
	Index int
}

// NewIterator 构造函数
func NewIterator(data []interface{}) *Iterator {
	return &Iterator{Data: data}
}

// HasNext 判断是否还有下一个
func (this *Iterator) HasNext() bool {
	if this.Data == nil || len(this.Data) == 0 {
		return false
	}
	return true
}

// Next 获取下一个
func (this *Iterator) Next() (ret interface{}) {
	ret = this.Data[this.Index]
	this.Index += 1
	return
}
