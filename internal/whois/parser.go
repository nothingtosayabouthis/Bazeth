package whois

import (
	"strings"

	"github.com/openrdap/rdap"
)

func bestRegistrar(entities []rdap.Entity) *rdap.Entity {
	for i := range entities {
		for _, role := range entities[i].Roles {
			if strings.EqualFold(role, "registrar") {
				return &entities[i]
			}
		}
	}

	return nil
}
