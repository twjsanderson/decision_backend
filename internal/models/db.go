package models

type NewTable struct {
	TableName  string
	PrimaryKey string
	Columns    map[string]string
}
