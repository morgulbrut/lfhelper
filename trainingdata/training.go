package trainingdata

import "time"

type TrainingUnit struct {
	ProgName  string
	RunTime   time.Duration
	Cals      int
	Distance  float64
	RiseDist  int
	AvSpeed   float64
	AvPace    time.Duration
	AvPower   int
	AvHeartR  int
	Product   string
	Timestamp time.Time
}

type Training struct {
	Units []TrainingUnit
}
