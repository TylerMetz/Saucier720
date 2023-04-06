package BackendPkg

import(
	"fmt"
)


type User struct{
	FirstName string
	LastName string
	Email string
    UserName string `json:"username"`
    Password string `json:"password"`
	UserPantry Pantry
}

func (u* User) PrintUserInfo(){
	fmt.Println("First Name: " + u.FirstName)
	fmt.Println("Last Name: " + u.LastName)
	fmt.Println("Email: " + u.Email)
	fmt.Println("UserName: " + u.UserName)
	fmt.Println("Password: " + u.Password)
	u.UserPantry.DisplayPantry()
}

