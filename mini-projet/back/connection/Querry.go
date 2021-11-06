package connection

import (
	"dataacces"
)

type notSecuredQuery struct {
	info ConnetionInfo
}
type securedQuery struct {
	id string
}

func (query securedQuery) getResult() *ConnectedUser {
	rows, _ := dataacces.ExecQuery("SELECT ID, USER, PASSWORD FROM USER WHERE PASSWORD = ?", query.id)
	if !rows.Next() {
		return nil
	}
	var user ConnectedUser
	_ = rows.Scan(&user.Id, &user.User, &user.Password)
	return &user
}
