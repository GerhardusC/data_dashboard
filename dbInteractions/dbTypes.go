package dbInteractions

type DBRowMeasurement[T string | float64] struct {
	Timestamp int
	Topic string
	Value T
}

