package implement

import (
	"fmt"
)

func makeCompanyIDFilters(companyID string) (filters []string) {
	return []string{
		fmt.Sprintf("id:eq:%s", companyID),
	}
}
