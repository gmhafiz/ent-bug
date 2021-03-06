// Code generated by entc, DO NOT EDIT.

package match

const (
	// Label holds the string label denoting the match type in the database.
	Label = "match"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// EdgeHomeTeam holds the string denoting the home_team edge name in mutations.
	EdgeHomeTeam = "home_team"
	// EdgeAwayTeam holds the string denoting the away_team edge name in mutations.
	EdgeAwayTeam = "away_team"
	// Table holds the table name of the match in the database.
	Table = "matches"
	// HomeTeamTable is the table that holds the home_team relation/edge.
	HomeTeamTable = "matches"
	// HomeTeamInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	HomeTeamInverseTable = "teams"
	// HomeTeamColumn is the table column denoting the home_team relation/edge.
	HomeTeamColumn = "match_home_team"
	// AwayTeamTable is the table that holds the away_team relation/edge.
	AwayTeamTable = "matches"
	// AwayTeamInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	AwayTeamInverseTable = "teams"
	// AwayTeamColumn is the table column denoting the away_team relation/edge.
	AwayTeamColumn = "match_away_team"
)

// Columns holds all SQL columns for match fields.
var Columns = []string{
	FieldID,
	FieldStartDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "matches"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"match_home_team",
	"match_away_team",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
