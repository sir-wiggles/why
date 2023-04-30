package clients

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sir-wiggles/why/mocks"
	"gotest.tools/assert"
)

func TestFooGet(t *testing.T) {
	ctrl := gomock.NewController(t)

	foo := mocks.NewMockFoo(ctrl)
	foo.EXPECT().Get().Return("foo").Times(1)

	assert.Equal(t, "foo", foo.Get())
}
