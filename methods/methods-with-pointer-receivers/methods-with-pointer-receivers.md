# Choosing a value or pointer receiver

There are two reasons for choosing methods with pointer receivers

First is that you can modify value that its receiver points to

Second is avoid copying value on each method call, and it can be more efficient in a large struct, for example
