package utils

import (
	"fmt"
	"strings"
)

func GetSearchValues(text string) []string {
	// Splitting the data into categories
	categories := make(map[string][]string)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		if strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			category := parts[0]
			terms := strings.Split(parts[1], ",")
			categories[category] = terms
		}
	}

	// Prepare to store generated search queries
	var searchQueries []string

	// Iterate over each category (except ExperienceLevel)
	for category, terms := range categories {
		if category != "ExperienceLevel" {
			// Group terms into batches of 10 with OR
			for i := 0; i < len(terms); i += 10 {
				end := i + 10
				if end > len(terms) {
					end = len(terms)
				}
				group := terms[i:end]

				// Wrap each term in intext:""
				for j := range group {
					group[j] = fmt.Sprintf("intext:\"%s\"", group[j])
				}
				orGroup := strings.Join(group, " OR ")
				searchQueries = append(searchQueries, fmt.Sprintf("(%s)", orGroup))
			}
		}
	}

	var result []string
	// Add ExperienceLevel combinations
	experienceLevels := categories["ExperienceLevel"]
	for _, level := range experienceLevels {
		for _, query := range searchQueries {
			searchQuery := fmt.Sprintf("%s intext:\"%s\" intext:\"job\"", query, level)
			result = append(result, searchQuery)
		}
	}

	return result
}
