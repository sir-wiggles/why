package clients

//go:generate mockgen -destination=../mocks/mock_foo.go -package=mocks . Foo
type Foo interface {
	Get() string
}

type foo struct{}

func NewFoo() *foo {
	return new(foo)
}

func (f *foo) Get() string {
	return "foo.Get()"
}
