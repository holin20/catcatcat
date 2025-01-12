package ezgo

import (
	"cmp"
	"sort"
)

// MapStableForEach iterates over a map in increasing key order and applies the provided function to each key-value pair.
func MapStableForEach[K cmp.Ordered, V any](m map[K]V, f func(K, V)) {
	// Extract keys from the map
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// Sort the keys
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// Iterate over the sorted keys and apply the function
	for _, k := range keys {
		f(k, m[k])
	}
}
