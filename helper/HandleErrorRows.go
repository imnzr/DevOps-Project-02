package helper

import (
	"fmt"

	"github.com/imnzr/DevOps-Project-02/models/domain"
)

func HandleErrorRows(err error) (domain.User, error) {
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to scan row: %w", err)
	}

	return domain.User{}, nil
}
