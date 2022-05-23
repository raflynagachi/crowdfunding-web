package helpers

import "strings"

func PerksToSlice(campaignPerks string) []string {
	var perks []string
	for _, perk := range strings.Split(campaignPerks, ",") {
		perks = append(perks, strings.Trim(perk, " "))
	}
	return perks
}
