package db

import (
	"fmt"
)

func find(col,table,params string)  string{
	 return fmt.Sprintf("SELECT %s FROM %s WHERE %s",col,table,params)
}
func insert(cols,vals,table string)  string{
	return fmt.Sprintf("INSERT %s(%s) VALUES (%s)",table,cols,vals)
}
func update(table,sets,params string)  string{
	return fmt.Sprintf("UPDATE %s SET %s WHERE %s",table,sets,params)
}