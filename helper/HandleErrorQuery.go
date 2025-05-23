package helper

import (
	"fmt"

	"github.com/imnzr/DevOps-Project-02/models/domain"
)

func HandleQueryError(err error) (domain.User, error) {
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to execute query: %w", err)
	}
	return domain.User{}, nil
}
