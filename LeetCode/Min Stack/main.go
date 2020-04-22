type MinStack struct {
    stack []int
    minStack []int // stores current min value
}

/** initialize your data structure here. */
func Constructor() MinStack {
    var minStack = new(MinStack)
    minStack.stack = []int{}
    minStack.minStack = []int{}
    return *minStack
}

// push value to the top of the stack
// if the value is smaller than top value of minStack (or minStack is empty) then add this value to minStack 
func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack, x)
    if(len(this.minStack) == 0) {
        this.minStack = append(this.minStack, x)
    } else if x <= this.minStack[len(this.minStack) - 1] {
        this.minStack = append(this.minStack, x)
    }
}

// pops value from the top of the stack
// if this value is at the top of the minStack (current min value) then pop this value from minStack
func (this *MinStack) Pop()  {
    if this.stack[len(this.stack) - 1] == this.minStack[len(this.minStack) - 1] {
        this.minStack = this.minStack[:len(this.minStack) - 1]
    }
    this.stack = this.stack[:len(this.stack) - 1]
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack) - 1]
}

// return the top value from minStack
func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack) - 1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
