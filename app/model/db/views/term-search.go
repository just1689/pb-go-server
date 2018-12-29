package views

import "database/sql"

type TermSearchResult struct {
	Wid      sql.NullInt64
	TreeNode sql.NullInt64
	LowerRid sql.NullInt64
	UpperRid sql.NullInt64
}
