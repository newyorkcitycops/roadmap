# Channels

Channels allow send and receive data through goroutines

Make function is used for create channels and you must specify channel data type

Channels only send and receive data unless the other side is ready, so goroutines can be synchronized without locks or conditions
