type Foo struct {
    second chan struct{}
    third chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
        second: make(chan struct{}),
        third: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
    f.second <- struct{}{}
}

func (f *Foo) Second(printSecond func()) {
    <-f.second
	/// Do not change this line
	printSecond()
    f.third <- struct{}{}
}

func (f *Foo) Third(printThird func()) {
	// Do not change this line
    <-f.third
	printThird()
}
