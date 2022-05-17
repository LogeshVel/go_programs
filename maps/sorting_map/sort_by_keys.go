package main

import (
	"fmt"
	"sort"
)

func main() {
	population := map[string]int{
		"Australia":  24982688,
		"Qatar":      2781677,
		"Wales":      3139000,
		"Burundi":    11175378,
		"Guinea":     12414318,
		"Niger":      22442948,
		"Brazil":     209469333,
		"Malta":      484630,
		"Peru":       31989256,
		"Yemen":      28498687,
		"Ireland":    4867309,
		"Kenya":      51393010,
		"Montserrat": 5900,
		"Cuba":       11338138,
		"Nicaragua":  6465513,
		"Jordan":     9956011,
		"Gabon":      2119275,
	}

	keys := make([]string, 0, len(population))
	for k := range population {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, population[k])
	}

}
