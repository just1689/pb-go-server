package incoming

type TermSearch struct {
	Query        []Query  `json:"query"`
	SearchRange  string   `json:"search_range"`
	SearchFilter []string `json:"search_filter"`
	Texts        []string `json:"texts"`
}

type Query struct {
	UID      string            `json:"uid"`
	Inverted bool              `json:"inverted"`
	Data     map[string]string `json:"data"`
}

/*
{
	"query":[
			{
				"uid":"1538779288476",
				"inverted":false,
				"data":{
					"sdbh":"Liquids","nu":"pl"
                }
		     }
     ],
     "search_range":"phrase",
	 "search_filter":[],
     "texts":["sbl","net","wlc","lxx"]
}
*/
