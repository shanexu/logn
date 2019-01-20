package common

import "strconv"

type FloatSet map[float64]struct{}

func MakeFloatSetFromStrings(strings ...string) (FloatSet, error) {
	if len(strings) == 0 {
		return nil, nil
	}

	set := FloatSet{}
	for _, str := range strings {
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		set[f] = struct{}{}
	}
	return set, nil
}

func MakeFloatSet(floats ...float64) FloatSet {
	if len(floats) == 0 {
		return nil
	}

	set := FloatSet{}
	for _, f := range floats {
		set[f] = struct{}{}
	}
	return set
}

func (set FloatSet) Add(f float64) {
	set[f] = struct{}{}
}

func (set FloatSet) Del(f float64) {
	delete(set, f)
}

func (set FloatSet) Count() int {
	return len(set)
}

func (set FloatSet) Has(f float64) (exists bool) {
	if set != nil {
		_, exists = set[f]
	}
	return
}
