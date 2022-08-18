package list

// List 接口
type List interface {
	Size() int
	IsEmpty() bool
	Contains(interface{}) bool
	Get(index int) (interface{}, error)
	Set(index int, value interface{}) error
	Insert(index int, value interface{}) error
	Append(value interface{}) error
	Remove(index int) error
	IndexOf(value interface{}) int
	Clear()
	String() string
}
