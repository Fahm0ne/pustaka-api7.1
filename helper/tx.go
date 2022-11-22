package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
	errorRolback :=	tx.Rollback()
		PanicIfError(errorRolback)
	} else {
	errorCommit :=	tx.Commit()
	PanicIfError(errorCommit)
	}
}