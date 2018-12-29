package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/just1689/pb-go-server/io"
	"github.com/just1689/pb-go-server/model"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

func HandleIncoming(client *io.Client, msg []byte) {
	var message model.Message
	if err := json.Unmarshal(msg, &message); err != nil {
		fmt.Println(fmt.Sprintf("Error unmarshaling, %s", err))
	}
	HandleTermSearch(client, message)

}

func HandleTermSearch(client *io.Client, message model.Message) {

	//Not actually doing the work

	query := `
		SELECT
			${eachWid(queryCount)},
			word0.${treeNode} AS tree_node,
			rid_range_by_tree_node.lower_rid,
			rid_range_by_tree_node.upper_rid
		FROM
			${oneTablePerWord(queryCount)},
			rid_range_by_tree_node
		WHERE
			${whereClauseElements.join("\n\t\t\tAND\n\t\t\t")}
			AND
			rid_range_by_tree_node.tree_node = word0.${treeNode}
		ORDER BY
			word0.${treeNode};
	`

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
	stmtOut, err := db.Prepare(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()

	rows, err := stmtOut.Query("", "", "")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var termSearchResult model.TermSearchResult
		if err := rows.Scan(&termSearchResult.Wid, &termSearchResult.TreeNode, &termSearchResult.UpperRid, &termSearchResult.LowerRid); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Row read: %v, %v, %v, %v", termSearchResult.Wid, termSearchResult.TreeNode, termSearchResult.LowerRid, termSearchResult.UpperRid)
	}

	//???
	if !rows.NextResultSet() {
		log.Fatal("expected more result sets", rows.Err())
	}

}
