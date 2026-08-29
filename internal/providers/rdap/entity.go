package rdap

import (
	"strings"

	"github.com/openrdap/rdap"
)

// bestEntity returns the most relevant entity based on RDAP roles.
func bestEntity(entities []rdap.Entity) *rdap.Entity {
	if len(entities) == 0 {
		return nil
	}

	priority := []string{
		"registrant",
		"technical",
		"administrative",
		"abuse",
	}

	for _, wanted := range priority {
		for i := range entities {
			for _, role := range entities[i].Roles {
				if strings.EqualFold(role, wanted) {
					return &entities[i]
				}
			}
		}
	}

	return &entities[0]
}
