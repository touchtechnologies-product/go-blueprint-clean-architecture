package implement

import (
	"fmt"
)

func makeCompanyIDFilters(companyID string) (filters []string) {
	return []string{
		fmt.Sprintf("_id:eq:%s", companyID),
	}
}
