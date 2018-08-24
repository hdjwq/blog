package db

import "fmt"

func find(cols,table,params string)  string{
	 return fmt.Sprintf("SELECT %s FROM %s WHERE %s",cols,table,params)
}

func insert(table,cols,vals string)  string{
	return fmt.Sprintf("INSERT %s (%s) VALUES (%s)",table,cols,vals)
}
func has(cols,table,params string)  string{
	return fmt.Sprintf("SELECT COUNT(%s) FROM %s WHERE %s",cols,table,params)
}
func update(table,cols,params string)  string{
	 return fmt.Sprintf("UPDATE %s SET %s WHERE %s",table,cols,params)
}
