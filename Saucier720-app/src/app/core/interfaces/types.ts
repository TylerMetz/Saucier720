export interface LoginRequest {
    UserName: string;
    Password: string;
}

export interface LogoutRequest {
    UserName: string;
}
export interface GetPantryRequest {
    UserName: string;
}

export interface SignupRequest {
    UserName:   string;
    FirstName:  string; 
    LastName:   string; 
    Email:      string; 
    Password:   string; 
}
