package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/just1689/pb-go-server/io/ws"
	"github.com/just1689/pb-go-server/model/db/views"
	"github.com/just1689/pb-go-server/model/incoming"
	"log"
	"strings"
)

func HandleTermSearch(client *ws.Client, raw []byte) {

	var messageTermSearch incoming.MessageTermSearch
	if err := json.Unmarshal(raw, &messageTermSearch); err != nil {
		log.Println(fmt.Sprintf("Error unmarshaling, %s", err))
	}

	clauses := queriesToClause(convertMapToQuery(messageTermSearch.Payload.Query[0].Data))
	query := `
		SELECT
			w.wid,
			w._verse_node AS tree_node,
			rid_range_by_tree_node.lower_rid,
			rid_range_by_tree_node.upper_rid
		FROM
			word_features w,
			rid_range_by_tree_node
		WHERE
			CLAUSES
			AND
			rid_range_by_tree_node.tree_node = w._verse_node
		ORDER BY
			w._verse_node;
	`
	sqlQuery := strings.Replace(query, "CLAUSES", clauses, 1)

	//Execute sql
	db, err := sql.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/parabible_test")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Prepare statement for reading data
	stmtOut, err := db.Prepare(sqlQuery)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {

		var termSearchResult views.TermSearchResult
		if err := rows.Scan(&termSearchResult.Wid, &termSearchResult.TreeNode, &termSearchResult.UpperRid, &termSearchResult.LowerRid); err != nil {
			log.Fatal(err)
		}
		log.Println("Row read: %i, %i, %i, %i", termSearchResult.Wid, termSearchResult.TreeNode, termSearchResult.LowerRid, termSearchResult.UpperRid)
	}

	//???
	if !rows.NextResultSet() {
		log.Fatal("expected more result sets", rows.Err())
	}

}

func queriesToClause(queries []string) string {
	s := len(queries)
	if s == 0 {
		return ""
	}
	return strings.Join(queries[:], " AND ")
}

func convertMapToQuery(m map[string]string) []string {
	var results []string
	for k, v := range m {
		results = append(results, fieldValueToQuery(k, v))
	}
	return results
}

func fieldValueToQuery(key, value string) string {
	return fmt.Sprintf("_%s='%s'", key, value)
}
