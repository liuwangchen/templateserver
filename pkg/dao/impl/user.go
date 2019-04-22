package impl

import "fmt"

type UserDao struct {
}

func (*UserDao) GetUserById(id string) string {
	fmt.Println(id)
	return fmt.Sprintf("get %v", id)
}
