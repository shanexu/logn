package common

import "strconv"

type IntSet map[int]struct{}

func MakeIntSetFromStrings(strings ...string) (IntSet, error) {
	if len(strings) == 0 {
		return nil, nil
	}

	set := IntSet{}
	for _, str := range strings {
		i, err := strconv.ParseInt(str, 0, 64)
		if err != nil {
			return nil, err
		}
		set[int(i)] = struct{}{}
	}
	return set, nil
}

func MakeIntSet(ints ...int) IntSet {
	if len(ints) == 0 {
		return nil
	}

	set := IntSet{}
	for _, i := range ints {
		set[i] = struct{}{}
	}
	return set
}

func (set IntSet) Add(i int) {
	set[i] = struct{}{}
}

func (set IntSet) Del(i int) {
	delete(set, i)
}

func (set IntSet) Count() int {
	return len(set)
}

func (set IntSet) Has(s int) (exists bool) {
	if set != nil {
		_, exists = set[s]
	}
	return
}
