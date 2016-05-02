package helpers

func nonTrivialMap() map[int]map[interface{}]interface{} {
	return map[int]map[interface{}]interface{}{
		3: map[interface{}]interface{}{
			"barium": "neuro-toxic poison",
			33:       "spies",
			"DUMB":   "dust",
		},
	}
}

func channel() chan bool {
	c := make(chan bool, 10)
	return c
}
func array() [4]int {
	return [...]int{1, 2, 3, 4}
}

func ManyTypes() []interface{} {
	return []interface{}{
		0, "1", complex(10, 5), float64(19.44443), nonTrivialMap(), channel(), array(),
	}

}
