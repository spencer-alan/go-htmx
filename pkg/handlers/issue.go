package handlers

import (
	"database/sql"
	"fmt"
	"ghost/internal/pages"

	"github.com/labstack/echo/v4"
)

type IssueHandler struct {
	DB *sql.DB
}

func (h IssueHandler) HandleAll(c echo.Context) error {
	mySlice := []string{"Hello", "Is", "This", "On"}

	rows, err := h.DB.Query("SELECT * FROM issues")
	if err != nil {
		return err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	var issues []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return err
		}

		row := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]

			switch val.(type) {
			case []byte:
				v = string(val.([]byte))
			default:
				v = val
			}
			row[col] = v
		}

		issues = append(issues, row)
	}
	for i, issue := range issues {
		fmt.Printf("Issue %d: %v\n", i, issue)
	}

	return render(c, pages.AllIssues(mySlice))
}
