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

func ValidateUser(currUser User) string{
	passwordDb := Database{
		Name: "func pwdb",
	}

	var returnCookie string 

	// checks if password read in is equal to db password
	returnPassword := passwordDb.GetUserPassword(currUser.UserName)
	if returnPassword == currUser.Password{
		// stores new cookie in database
		passwordDb.StoreCookie(currUser.UserName,GenerateCookie(currUser.UserName))
		returnCookie = passwordDb.ReadCookie(currUser.UserName)

	} 
	
	return returnCookie
}

func GenerateCookie(username string) string{
	// cookie generation function
	var returnCookie string = username + "720"

	return returnCookie
}