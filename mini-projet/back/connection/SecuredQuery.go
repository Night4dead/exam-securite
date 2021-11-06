package connection

import (
	"dataacces"
)

func (query securedQuery) GetResult() *ConnectedUser {
	rows, _ := dataacces.ExecPreparedQuery("SELECT ID, USER, PASSWORD FROM USER WHERE PASSWORD = ?;", query.id)
	if !rows.Next() {
		return nil
	}
	var user ConnectedUser
	_ = rows.Scan(&user.Id, &user.User, &user.Password)
	return &user
}
