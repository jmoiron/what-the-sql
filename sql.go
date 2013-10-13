package main

// sql execution support

import (
	"errors"

	"github.com/jmoiron/modl"
	"github.com/jmoiron/termtable"
)

type Question struct {
	Number int                        `json:"number"`
	Text   string                     `json:"text"`
	Check  func([][]interface{}) bool `json:"-"`
}

// compare to lists of strings as though they were a set
func setCompare(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m := map[string]uint8{}
	for _, s := range s1 {
		m[s] = 0
	}
	for _, s := range s2 {
		if _, ok := m[s]; !ok {
			return false
		}
	}
	return true
}

func stringResults(results [][]interface{}) ([]string, error) {
	sres := make([]string, 0, len(results))
	for _, res := range results {
		if res == nil {
			return sres, errors.New("Null encountered where string expected.")
		}
		sres = append(sres, res[0].(string))
	}
	return sres, nil
}

func NoopCheck(results [][]interface{}) bool {
	return false
}

func CompareResults(results [][]interface{}, answers []string) bool {
	if len(results) != len(answers) {
		return false
	}
	if len(results[0]) != 1 {
		return false
	}
	stringRes, err := stringResults(results)
	if err != nil {
		return false
	}
	return setCompare(stringRes, answers)
}

func Question1(results [][]interface{}) bool {
	return CompareResults(results, Question1Answer)
}

func Question2(results [][]interface{}) bool {
	return CompareResults(results, Question2Answer)
}

func Question3(results [][]interface{}) bool {
	return CompareResults(results, Question3Answer)
}

func Question4(results [][]interface{}) bool {
	return CompareResults(results, Question4Answer)
}

func Question5(results [][]interface{}) bool {
	return CompareResults(results, Question5Answer)
}

func Question6(results [][]interface{}) bool {
	if len(results) != len(Question6Answer) {
		return false
	}
	if len(results[0]) != 2 {
		return false
	}

	org1 := make([]string, 0, len(results))
	org2 := make([]string, 0, len(results))

	for _, res := range results {
		if res[0] == nil || res[1] == nil {
			return false
		}
		s1, s2 := res[0].(string), res[1].(string)
		org1 = append(org1, s1+"|"+s2)
		org2 = append(org2, s2+"|"+s1)
	}
	return setCompare(org1, Question6Answer) || setCompare(org2, Question6Answer)

}

var Questions = []Question{
	{1, "List the top 10 employee names by salary.", Question1},
	{2, "List employee names who have a bigger salary than their boss.", Question2},
	{3, "List employee names who have the biggest salary in their departments.", Question3},
	{4, "List department names that have less than 10 people in it.", Question4},
	{5, "List employees who have a boss in a different department from them.", Question5},
	{6, "List all department names along with the total salary there.", Question6},
}

type SqlResult struct {
	Results         [][]interface{}
	Columns         []string
	CurrentQuestion int
	Answered        bool
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

	status := Status{}
	dbm.Dbx.Get(&status, "SELECT * FROM _status;")
	result.CurrentQuestion = status.CurrentQuestion
	result.Answered = Questions[result.CurrentQuestion-1].Check(result.Results)

	return result, nil
}
