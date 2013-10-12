package main

// sql execution support

import (
	"github.com/jmoiron/modl"
	"github.com/jmoiron/termtable"
)

type Question struct {
	Number int
	Text   string
	Check  func([][]interface{}) bool
}

func NoopCheck(results [][]interface{}) bool {
	return false
}

var Questions = []Question{
	{1, "List the top 10 employee names by salary.", NoopCheck},
	{2, "List employee names who have a bigger salary than their boss.", NoopCheck},
	{3, "List employee names who have teh biggest salary in their departments.", NoopCheck},
	{4, "List departments that have less than 3 people in it.", NoopCheck},
	{5, "List all departments along with the number of people.", NoopCheck},
	{6, "List employees who have a boss in a different department from them.", NoopCheck},
	{7, "List all departments along with the total salary there.", NoopCheck},
}

type SqlResult struct {
	Results      [][]interface{}
	Columns      []string
	QuestionInfo []bool
}

func (s *SqlResult) String() string {
	if s == nil {
		return ""
	}
	if len(s.Results) == 0 {
		return ""
	}
	return formatResults(s.Results, s.Columns)
}

func formatResults(results [][]interface{}, columns []string) string {
	rows := [][]string{}
	for _, row := range results {
		s := []string{}
		for _, res := range row {
			if res == nil {
				s = append(s, "NULL")
			} else {
				s = append(s, res.(string))
			}
		}
		rows = append(rows, s)
	}
	t := termtable.NewTable(rows, &termtable.TableOptions{Padding: 2, UseSeparator: true})
	t.SetHeader(columns)
	return t.Render()
}

// Execute sql, checking the results against the questions in the quiz, and
// returning a result which can be String()'d to get a tabularized test result
func execSql(dbm *modl.DbMap, sql string) (*SqlResult, error) {

	result := &SqlResult{}
	result.Results = make([][]interface{}, 0, 10)
	result.QuestionInfo = make([]bool, len(Questions))

	rows, err := dbm.Dbx.Queryx(sql)
	if err != nil {
		return result, err
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	result.Columns = columns

	var res []interface{}
	for rows.Next() {
		res, err = rows.SliceScan()
		if err != nil {
			return result, err
		}
		result.Results = append(result.Results, res)
	}

	for i, q := range Questions {
		result.QuestionInfo[i] = q.Check(result.Results)
	}

	return result, nil
}
