package implement

import (
	"fmt"
)

func makeStaffIDFilters(staffID string) (filters []string) {
	return []string{
		fmt.Sprintf("id:eq:%s", staffID),
	}
}
