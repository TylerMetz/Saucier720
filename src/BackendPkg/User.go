package BackendPkg

import(
	_ "fmt"
)


type User struct{
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	Email string `json:"Email"`
    UserName string `json:"UserName"`
    Password string `json:"Password"`
	UserPantry Pantry 
	UserList List 
}


func ValidateUser(currUser User) string{
	passwordDb := Database{
		Name: "func pwdb",
	}

	var returnCookie string 

	// checks if password read in is equal to db password
	returnPassword, _ := passwordDb.GetUserPassword(currUser.UserName)
	if returnPassword == currUser.Password{
		// stores new cookie in database
		passwordDb.StoreCookie(currUser.UserName,GenerateCookie(currUser.UserName))
		returnCookie, _ = passwordDb.ReadCookie(currUser.UserName)

	} 
	
	return returnCookie
}

func GenerateCookie(username string) string{
	// cookie generation function
	var returnCookie string = username + "720"

	return returnCookie
}