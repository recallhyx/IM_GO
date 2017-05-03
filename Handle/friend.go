package Handle

import (
	"../database"
)

func GetFriendList(username string)(  friend []string , b bool){
	friend,b=database.GetFriend(username)
	return
}

func AddFriend(username string,friendname string)(b bool){
	b=database.AddFriend(username,friendname)
	return
}

func DelFriend(username string,friendname string)(b bool){
	b=database.DelFriend(username,friendname)
	return
}
