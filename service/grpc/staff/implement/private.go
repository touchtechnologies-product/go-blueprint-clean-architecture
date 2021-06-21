package implement

import (
	"fmt"
)

func makeStaffIDFilters(staffID string) (filters []string) {
	return []string{
		fmt.Sprintf("_id:eq:%s", staffID),
	}
}
