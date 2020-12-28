package oracle

import (
	"strings"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func IsReservedWord(v string) bool {
	return contains(ReservedWordsList, strings.ToUpper(v))
}

var ReservedWordsList = []string{
	"AGGREGATE", "AGGREGATES", "ALL", "ALLOW", "ANALYZE", "ANCESTOR", "AND", "ANY", "AS", "ASC", "AT", "AVG", "BETWEEN",
	"BINARY_DOUBLE", "BINARY_FLOAT", "BLOB", "BRANCH", "BUILD", "BY", "BYTE", "CASE", "CAST", "CHAR", "CHILD", "CLEAR",
	"CLOB", "COMMIT", "COMPILE", "CONSIDER", "COUNT", "DATATYPE", "DATE", "DATE_MEASURE", "DAY", "DECIMAL", "DELETE",
	"DESC", "DESCENDANT", "DIMENSION", "DISALLOW", "DIVISION", "DML", "ELSE", "END", "ESCAPE", "EXECUTE", "FIRST",
	"FLOAT", "FOR", "FROM", "HIERARCHIES", "HIERARCHY", "HOUR", "IGNORE", "IN", "INFINITE", "INSERT", "INTEGER",
	"INTERVAL", "INTO", "IS", "LAST", "LEAF_DESCENDANT", "LEAVES", "LEVEL", "LIKE", "LIKEC", "LIKE2", "LIKE4", "LOAD",
	"LOCAL", "LOG_SPEC", "LONG", "MAINTAIN", "MAX", "MEASURE", "MEASURES", "MEMBER", "MEMBERS", "MERGE", "MLSLABEL",
	"MIN", "MINUTE", "MODEL", "MONTH", "NAN", "NCHAR", "NCLOB", "NO", "NONE", "NOT", "NULL", "NULLS", "NUMBER",
	"NVARCHAR2", "OF", "OLAP", "OLAP_DML_EXPRESSION", "ON", "ONLY", "OPERATOR", "OR", "ORDER", "OVER", "OVERFLOW",
	"PARALLEL", "PARENT", "PLSQL", "PRUNE", "RAW", "RELATIVE", "ROOT_ANCESTOR", "ROWID", "SCN", "SECOND", "SELF",
	"SERIAL", "SET", "SOLVE", "SOME", "SORT", "SPEC", "SUM", "SYNCH", "TEXT_MEASURE", "THEN", "TIME", "TIMESTAMP",
	"TO", "UNBRANCH", "UPDATE", "USING", "VALIDATE", "VALUES", "VARCHAR2", "WHEN", "WHERE", "WITHIN", "WITH", "YEAR",
	"ZERO", "ZONE",
}
