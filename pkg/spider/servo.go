package spider

type Servo struct {
	pin             byte
	minVal          uint16
	maxVal          uint16
	ninetyDegMicros int16
	reversed        bool
}

func GetServos() [12]*Servo {
	return [12]*Servo{
		//
		// Front right leg.
		//
		// Hip.
		{
			pin:             0,
			minVal:          1500,
			maxVal:          2500,
			ninetyDegMicros: 1700,
			reversed:        true,
		},
		// Coxa.
		{
			pin:             1,
			minVal:          1200,
			maxVal:          2600,
			ninetyDegMicros: 2111,
			reversed:        false,
		},
		// Tibia.
		{
			pin:             2,
			minVal:          1400,
			maxVal:          2500,
			ninetyDegMicros: 1900,
			reversed:        false,
		},
		//
		// Front left leg.
		//
		// Hip.
		{
			pin:             4,
			minVal:          700,
			maxVal:          1700,
			ninetyDegMicros: 1611,
			reversed:        false,
		},
		// Coxa.
		{
			pin:             5,
			minVal:          500,
			maxVal:          1900,
			ninetyDegMicros: 1045,
			reversed:        true,
		},
		// Tibia.
		{
			pin:             6,
			minVal:          1300,
			maxVal:          2400,
			ninetyDegMicros: 1900,
			reversed:        true,
		},

		// 	0     1500     2200     2500
		// 	4      700     1200     1700
		// 	8      700     1200     1700
		//    12     1400     1900     2400

		// 	1     1200     2000     2600 (smaller == upwards)
		// 	5      500     1100     1900 (bigger == upwards)
		// 	9      700     1300     2100 (bigger == upwards)
		// 	13 --- BAD PIN --- (middle = 1800?) (smaller == upwards?)

		// 	2     1400     1900     2500 (bigger == outwards)
		// 	6     1300     1900     2400 (smaller == outwards)
		//    10     1500     2100     2500 (smaller == outwards)
		//   14     1100     1600     2200 (bigger == outwards)

	}
}

func (s Servo) Pin() byte {
	return s.pin
}

func (s *Servo) DegreesToMicros(deg int16) uint16 {
	if s.reversed {
		deg = 90 - deg
	} else {
		deg = deg - 90
	}
	micros := uint16(deg*100/9 + s.ninetyDegMicros)
	if micros > s.maxVal {
		return s.maxVal
	}
	if micros < s.minVal {
		return s.minVal
	}
	return micros
}
