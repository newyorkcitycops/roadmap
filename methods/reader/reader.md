# Reader

io package specifies Reader interface

Reader interface has a Read method

Read method returns a byte slice with data, its len and an error value, and returns an EOF error when the stream ends

You can actually define how many bytes it will read per iteration, when denoting byte slice length
