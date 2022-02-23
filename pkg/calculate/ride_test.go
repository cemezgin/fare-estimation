package calculate

import (
	"fare-estimation/internal/ride"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLessThanHundredKM(t *testing.T) {

	ridesList := []ride.Ride{}

	timestamp1, err := strconv.ParseInt("1405587697", 10, 64)
	if err != nil {
		assert.Error(t, err)
	}
	tm1 := time.Unix(timestamp1, 0)

	timestamp2, err := strconv.ParseInt("1405587707", 10, 64)
	if err != nil {
		assert.Error(t, err)
	}
	tm2 := time.Unix(timestamp2, 0)

	timestamp3, err := strconv.ParseInt("1405587717", 10, 64)
	if err != nil {
		assert.Error(t, err)
	}
	tm3 := time.Unix(timestamp3, 0)

	timestamp4, err := strconv.ParseInt("1405587727", 10, 64)
	if err != nil {
		assert.Error(t, err)
	}
	tm4 := time.Unix(timestamp4, 0)

	ridesList = append(ridesList,
		ride.Ride{1, 37.953066, 23.735606, tm1},
		ride.Ride{1, 37.953009, 23.735593, tm2},
		ride.Ride{1, 37.953195, 23.736224, tm3},
		ride.Ride{1, 37.953433, 23.736926, tm4},
	)

	tuples := []ride.RideTuples{}

	tuples = append(tuples, ride.RideTuples{
		0.006439786548918445,
		0.002777777777777778,
		2.31832315761064,
	},
		ride.RideTuples{
			0.05906477219481011,
			0.002777777777777778,
			21.263317990131636,
		},
		ride.RideTuples{
			0.06699857255258379,
			0.002777777777777778,
			24.119486118930162,
		},
	)

	assert.Equal(t, tuples, Filter(ridesList).rideList)
}

func TestMoreThanHundredKM(t *testing.T) {
	//@todo fix after found way to filter
	ridesList := []ride.Ride{}

	timestamp5, err := strconv.ParseInt("1405587758", 10, 64)
	if err != nil {
		assert.Error(t, err)
	}
	tm5 := time.Unix(timestamp5, 0)

	timestamp6, err := strconv.ParseInt("1405587768", 10, 64)
	if err != nil {
		assert.Error(t, err)
	}
	tm6 := time.Unix(timestamp6, 0)

	ridesList = append(ridesList,
		ride.Ride{8, 37.963705, 23.732530, tm5},
		ride.Ride{8, 37.968053, 23.730544, tm6},
	)

	tuples := []ride.RideTuples{}

	tuples = append(tuples,
		ride.RideTuples{
			0.5138670483785056,
			0.002777777777777778,
			184.992137416262,
		},
	)

	assert.Equal(t, tuples, Filter(ridesList).rideList)
}

func TestRide_FareAmountMinAmount(t *testing.T) {
	tuples := []ride.RideTuples{}

	timestamp1, err := strconv.ParseInt("1405587697", 10, 64)
	if err != nil {
		assert.Error(t, err)
	}
	tm1 := time.Unix(timestamp1, 0)

	tuples = append(tuples,
		ride.RideTuples{
			0.006439786548918445,
			0.002777777777777778,
			2.31832315761064,
		},
		ride.RideTuples{
			0.05906477219481011,
			0.002777777777777778,
			21.263317990131636,
		},
		ride.RideTuples{
			0.06699857255258379,
			0.002777777777777778,
			24.119486118930162,
		},
	)

	rc := RideCalculation{tuples, tm1}

	assert.Equal(t, 3.47, rc.FareAmount())
}

func TestRide_FareAmountDay(t *testing.T) {
	tuples := []ride.RideTuples{}

	timestamp1, err := strconv.ParseInt("1405587697", 10, 64)
	if err != nil {
		assert.Error(t, err)
	}
	tm1 := time.Unix(timestamp1, 0)

	tuples = append(tuples,
		ride.RideTuples{
			11.006439786548918445,
			1.002777777777777778,
			0.31832315761064,
		},
		ride.RideTuples{
			12.05906477219481011,
			1.002777777777777778,
			1.263317990131636,
		},
		ride.RideTuples{
			400.06699857255258379,
			1.002777777777777778,
			24.119486118930162,
		},
	)

	rc := RideCalculation{tuples, tm1}

	assert.Equal(t, 25.90816666666667, rc.FareAmount())
}

func TestRide_FareAmountNight(t *testing.T) {
	tuples := []ride.RideTuples{}

	timestamp1, err := strconv.ParseInt("1645567200", 10, 64)
	if err != nil {
		assert.Error(t, err)
	}
	tm1 := time.Unix(timestamp1, 0)

	tuples = append(tuples,
		ride.RideTuples{
			11.006439786548918445,
			1.002777777777777778,
			0.31832315761064,
		},
		ride.RideTuples{
			12.05906477219481011,
			1.002777777777777778,
			1.263317990131636,
		},
		ride.RideTuples{
			400.06699857255258379,
			1.002777777777777778,
			24.119486118930162,
		},
	)

	rc := RideCalculation{tuples, tm1}

	assert.Equal(t, 26.469722222222224, rc.FareAmount())
}
