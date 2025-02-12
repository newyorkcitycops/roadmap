# Mutex

Go implements the concept of mutual exclusion with `sync.Mutex`

`sync.Mutex` provides two methods `Lock` and `Unlock`. A block of code surrounded by these methods perform mutual exclusion

`defer` can be used to ensure an unlocked mutex
