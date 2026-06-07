package printinorder

type Foo struct {
	channel1 chan struct{}
	channel2 chan struct{}
}

// NewFoo initializes the struct with unbuffered channels for synchronization.
// Time Complexity O(1). All channel operations (send and receive) and print calls execute in constant time.
// Space Complexity: O(1). The struct only allocates two unbuffered channels of empty structs, consuming a fixed amount of memory.
func NewFoo() *Foo {
	return &Foo{
		channel1: make(chan struct{}),
		channel2: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	printFirst()
	f.channel1 <- struct{}{}
}

func (f *Foo) Second(printSecond func()) {
	<-f.channel1
	printSecond()
	f.channel2 <- struct{}{}
}

func (f *Foo) Third(printThird func()) {
	<-f.channel2
	printThird()
}
