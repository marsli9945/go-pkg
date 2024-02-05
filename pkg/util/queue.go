package util

/*
*
扩展类型有两种方式
定义别名
*/
type Queue[T interface{}] []T

func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Pop() T {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}
