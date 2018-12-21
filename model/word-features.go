package model

type WordFeature struct {
	Wid           int32  `json:"wid"`
	GWordUtf8     string `json:"_g_word_utf8"`
	Tricons       string `json:"_tricons"`
	Lex           string `json:"_lex"`
	VocUtf8       string `json:"_voc_utf8"`
	Sp            string `json:"_sp"`
	IsDefinite    string `json:"_is_definite"`
	GConsUtf8     string `json:"_g_cons_utf8"`
	HasSuffix     string `json:"_has_suffix"`
	Gloss         string `json:"_gloss"`
	TrailerUtf8   string `json:"_trailer_utf8"`
	Sdbh          string `json:"_sdbh"`
	Lxxlexeme     string `json:"_lxxlexeme"`
	Nu            string `json:"_nu"`
	Gn            string `json:"_gn"`
	St            string `json:"_st"`
	Accent        string `json:"_accent"`
	AccentQuality string `json:"_accent_quality"`
	Ps            string `json:"_ps"`
	Vt            string `json:"_vt"`
	Vs            string `json:"_vs"`
	GPrsUtf8      string `json:"_g_prs_utf8"`
	PrsNu         string `json:"_prs_nu"`
	PrsGn         string `json:"_prs_gn"`
	PrsPs         string `json:"_prs_ps"`
	GUvfUtf8      string `json:"_g_uvf_utf8"`
	Pos           string `json:"_pos"`
	Lexeme        string `json:"_lexeme"`
	GkGloss       string `json:"_gk_gloss"`
	Case          string `json:"_case"`
	Number        string `json:"_number"`
	Gender        string `json:"_gender"`
	Person        string `json:"_person"`
	Tense         string `json:"_tense"`
	Voice         string `json:"_voice"`
	Mood          string `json:"_mood"`
	TcNote        string `json:"_tc_note"`
	CfWid         int32  `json:"_cf_wid"`
	Degree        string `json:"_degree"`
	Declension    string `json:"_declension"`
	PhraseNode    int32  `json:"_phrase_node"`
	ClauseNode    int32  `json:"_clause_node"`
	SentenceNode  int32  `json:"_sentence_node"`
	VerseNode     int32  `json:"_verse_node"`
}
