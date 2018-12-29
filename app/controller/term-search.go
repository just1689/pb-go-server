package controller

import (
	"database/sql"
	"github.com/just1689/pb-go-server/io/ws"
	"github.com/just1689/pb-go-server/model/db/views"
	"github.com/just1689/pb-go-server/model/incoming"
	"log"
)

func HandleTermSearch(client *ws.Client, message incoming.Message) {

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
