package mysql

const (
	// Find
	QueryFindQuestions = `SELECT id, no, question, answer FROM questions`

	// Insert
	QueryInsertQuestion = `INSERT INTO questions (no, question, answer) VALUES (:no, :question, :answer)`

	// Update
	QueryUpdateQuestion = `UPDATE questions SET no=:no, question=:question, answer=:answer WHERE no=:no`

	// delete
	QueryDeleteQuestion = `DELETE FROM questions WHERE no=:no`
)
