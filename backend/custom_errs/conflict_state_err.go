package custom_errs

import "net/http"

type ConflictStateErr struct {
}

func (c *ConflictStateErr) Error() string {
	return "conflict state custom_errs"
}

func (c *ConflictStateErr) GetHttpStatus() int {
	return http.StatusConflict
}
