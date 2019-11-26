package keyset

func SoloKeySet(k Key) KeySet {
	return oneKeySet{k}
}

type oneKeySet struct {
	key Key
}

func (v oneKeySet) EnumRawKeys(fn func(k Key, exclusive bool) bool) bool {
	return fn(v.key, false)
}

func (v oneKeySet) RawKeyCount() int {
	return 1
}

func (v oneKeySet) IsNothing() bool {
	return false
}

func (v oneKeySet) IsEverything() bool {
	return false
}

func (v oneKeySet) IsExclusive() bool {
	return false
}

func (v oneKeySet) Contains(k Key) bool {
	return v.key == k
}

func (v oneKeySet) ContainsAny(ks KeySet) bool {
	return ks.Contains(v.key)
}

func (v oneKeySet) SupersetOf(ks KeySet) bool {
	if ks.IsExclusive() {
		return false
	}

	switch ks.RawKeyCount() {
	case 0:
		return true
	case 1:
		return ks.Contains(v.key)
	default:
		return false
	}
}

func (v oneKeySet) SubsetOf(ks KeySet) bool {
	return ks.Contains(v.key)
}

func (v oneKeySet) Equal(ks KeySet) bool {
	if ks.IsExclusive() || v.RawKeyCount() != ks.RawKeyCount() {
		return false
	}
	return ks.Contains(v.key)
}

func (v oneKeySet) EqualInverse(ks KeySet) bool {
	if !ks.IsExclusive() || v.RawKeyCount() != ks.RawKeyCount() {
		return false
	}
	return !ks.Contains(v.key)
}

func (v oneKeySet) Inverse() KeySet {
	return exclusiveKeySet{basicKeySet{v.key: {}}}
}

func (v oneKeySet) Union(ks KeySet) KeySet {
	switch {
	case ks.Contains(v.key):
		return ks
	case ks.RawKeyCount() == 0:
		return v
	}
	return inclusiveKeySet{keyUnion(basicKeySet{v.key: {}}, ks)}
}

func (v oneKeySet) Intersect(ks KeySet) KeySet {
	switch {
	case ks.Contains(v.key):
		return v
	case ks.RawKeyCount() == 0:
		return ks
	default:
		return Nothing()
	}
}

func (v oneKeySet) Subtract(ks KeySet) KeySet {
	if ks.Contains(v.key) {
		return Nothing()
	}
	return v
}
