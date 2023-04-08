package BackendPkg

import(
	"fmt"
)


type User struct{
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	Email string `json:"Email"`
    UserName string `json:"UserName"`
    Password string `json:"Password"`
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

func ValidateUser(currUser User){
	
}