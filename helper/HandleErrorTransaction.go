package helper

import "log"

func HandleErrorTransaction(err error) {
	if err != nil {
		log.Printf("failed to commit transaction: %v", err)
	}
}
