package 08_stack

/*
基于数组实现的栈
*/

type ArrayStack struct {
	// 数据
	data []interface{}
	//栈顶指针
	top int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{},0, 32),
		top: -1,
	}
}

