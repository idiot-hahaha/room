package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	createLiveIDQAGroupTable = "CREATE TABLE IF NOT EXISTS live_qa_group (id SERIAL, room_id INTEGER NOT NULL, group_id INTEGER NOT NULL, created_at TIMESTAMP DEFAULT now()); )"

	// questionGroup
	createGroupTable = "CREATE TABLE IF NOT EXISTS groups (id SERIAL PRIMARY KEY, group_name TEXT NOT NULL, created_at TIMESTAMP DEFAULT now());"
	createQAGroup    = "insert into groups (group_name) values ($1) RETURNING id"
	selectGroup      = ""

	// question
	createQuestionTable              = "CREATE TABLE questions (id SERIAL PRIMARY KEY, group_id INTEGER NOT NULL, content TEXT NOT NULL, embedding VECTOR(384), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);"
	createGrouIDIndex                = "CREATE INDEX idx_questions_group_id ON questions(group_id);"
	selectMostSimilarQuestionByGroup = "Select id, content, embedding <=> $1 as distance from sentence_embeddings WHERE groupID = $2 ORDER BY distance LIMIT 1;"
	insertQuestionByGroup            = "Insert into questions (content, embedding, groupID) values (?, ?, ?);"

	// answer
	createAnswerTable        = "CREATE TABLE answers(id SERIAL PRIMARY KEY, question_id INTEGER NOT NULL, content TEXT NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);"
	selectAnswerByQuestionID = "SELECT id, content FROM answers WHERE question_id = $1;"
	//
	selectQSByGroupID = "SELECT g.name AS group_name, q.id AS question_id, q.question, a.id AS answer_id, a.content AS answer FROM qa_group g JOIN qa_question q ON q.group_id = g.id JOIN qa_answer a ON a.question_id = q.id WHERE g.id = 1 ORDER BY q.id, a.id;"
)

func (d *Dao) SelectMostSimilarQuestionByGroup(groupID int64, embedding []float64) (id int64, content string, distance float64, err error) {
	err = d.pgClient.QueryRow(selectMostSimilarQuestionByGroup, float64Slice2String(embedding), groupID).Scan(&id, &content, &distance)
	if errors.Is(err, sql.ErrNoRows) {
		return -1, "", -1, nil
	}
	if err != nil {
		return -1, "", -1, err
	}
	return
}

func float64Slice2String(floats []float64) (s string) {
	bytes := make([]byte, 0)
	bytes = append(bytes, "["...)
	for i, float := range floats {
		if i == len(floats)-1 {
			bytes = append(bytes, fmt.Sprintf("%f", float)...)
			break
		}
		bytes = append(bytes, fmt.Sprintf("%f,", float)...)
	}
	bytes = append(bytes, "]"...)
	s = string(bytes)
	return
}

func strings2Float64Slice(s string) ([]float64, error) {
	// 去掉首尾中括号
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")

	// 按逗号分割
	parts := strings.Split(s, " ")

	var result []float64
	for _, p := range parts {
		p = strings.TrimSpace(p)
		f, err := strconv.ParseFloat(p, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, f)
	}
	return result, nil
}

func (d *Dao) SelectAnswersByQuestionID(questionID int64) (answers []string, err error) {
	rows, err := d.pgClient.Query(selectAnswerByQuestionID, questionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	answers = make([]string, 0)
	for rows.Next() {
		var (
			id     int64
			answer string
		)
		if err := rows.Scan(&id, &answer); err != nil {
			log.Println(err)
			continue
		}
		answers = append(answers, answer)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (d *Dao) CreateQAGroup(ctx context.Context, groupName string) (id int64, err error) {
	err = d.pgClient.QueryRow(createQAGroup, groupName).Scan(&id)
	if err != nil {
		return -1, err
	}
	return
}
