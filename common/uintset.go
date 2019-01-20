package common

import "strconv"

type UintSet map[uint]struct{}

func MakeUintSetFromStrings(strings ...string) (UintSet, error) {
	if len(strings) == 0 {
		return nil, nil
	}

	set := UintSet{}
	for _, str := range strings {
		i, err := strconv.ParseUint(str, 0, 64)
		if err != nil {
			return nil, err
		}
		set[uint(i)] = struct{}{}
	}
	return set, nil
}

func MakeUintSet(uints ...uint) UintSet {
	if len(uints) == 0 {
		return nil
	}

	set := UintSet{}
	for _, i := range uints {
		set[i] = struct{}{}
	}
	return set
}

func (set UintSet) Add(i uint) {
	set[i] = struct{}{}
}

func (set UintSet) Del(i uint) {
	delete(set, i)
}

func (set UintSet) Count() int {
	return len(set)
}

func (set UintSet) Has(s uint) (exists bool) {
	if set != nil {
		_, exists = set[s]
	}
	return
}
