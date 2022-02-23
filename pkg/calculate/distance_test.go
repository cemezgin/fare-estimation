package calculate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDistanceCalculatorInterface struct {
	mock.Mock
}

func (m *MockDistanceCalculatorInterface) Distance(origin Coordinates, destination Coordinates) float64 {

	args := m.Called(origin, destination)
	return args.Get(0).(float64)

}

func TestDistanceResult(t *testing.T) {

	testObj := new(MockDistanceCalculatorInterface)

	testObj.On("Distance", Coordinates{37.966660, 23.728308}, Coordinates{37.966627, 23.728263}).Return(0.005387608950290441)
	assert.Equal(t, NewDistanceCalculator().Distance(Coordinates{37.966660, 23.728308}, Coordinates{37.966627, 23.728263}),
		testObj.Distance(Coordinates{37.966660, 23.728308}, Coordinates{37.966627, 23.728263}))
}
