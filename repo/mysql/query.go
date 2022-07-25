package mysql

const (
	// Find
	QueryFindQuestions = `SELECT id, no, question, answer FROM questions`

	// Insert
	QueryInsertQuestion = `INSERT INTO questions (id, no, question, answer) VALUES (:id, :no, :question, :answer)`

	// Update
	QueryUpdateQuestion = `UPDATE questions SET no=:no, question=:question, answer=:answer WHERE no=:no`

	// delete
	QueryDeleteQuestion = `DELETE FROM questions WHERE no=:no`
)
