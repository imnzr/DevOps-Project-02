package helper

import "database/sql"

func HandleTx(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			panic(errRollback)
		}
	} else {
		errCommit := tx.Commit()
		if errCommit != nil {
			panic(errCommit)
		}
	}
}
