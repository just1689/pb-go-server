package api

import (
	"fmt"
	"github.com/just1689/pb-go-server/io"
	"net/http"
)
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func HandleIncoming(client *io.Client, msg []byte) {

}

func HandleTermSearch(w http.ResponseWriter, r *http.Request) {

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
			word0.${treeNode};`

	//Execute sql
	db, err := sql.Open("mysql", "user:password@/dbname")
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

	var squareNum int // we "scan" the result in here

	// Query the square-number of 13
	err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 13 is: %d", squareNum)

	// Query another number.. 1 maybe?
	err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Printf("The square number of 1 is: %d", squareNum)

}

func HandleWordLookup(w http.ResponseWriter, r *http.Request) {

}

func HandleChapterText(w http.ResponseWriter, r *http.Request) {

}

func HandleNodeText(w http.ResponseWriter, r *http.Request) {

}
