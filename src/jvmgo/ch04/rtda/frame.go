package rtda

type Frame struct {
	lower        *Frame        // 指向下面的栈帧
	localVars    LocalVars     // 局部变量表
	operandStack *OperandStack // 操作数栈
}

func NewFrame(maxLocals uint, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
