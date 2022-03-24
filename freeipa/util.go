package freeipa

import (
	"strings"
)

func splitID(id string) (a, b, c string) {
	resource_map := make([]string, 3)
	copy(resource_map, strings.Split(id, "/"))
	return resource_map[0], resource_map[1], resource_map[2]
}

func sliceStrings(itemsRaw []interface{}) (slice []string) {
	items := make([]string, len(itemsRaw))
	for i, raw := range itemsRaw {
		items[i] = raw.(string)
	}
	return items
}
