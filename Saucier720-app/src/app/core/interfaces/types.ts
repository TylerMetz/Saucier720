import { Ingredient, Pantry, List} from "./ingredient";

export interface LoginRequest {
    UserName: string;
    Password: string;
}

export interface LogoutRequest {
    UserName: string;
}
export interface GetPantryRequest {
    UserName: string;
    Pantry: Pantry;
}

export interface GetPantryResponse {
    Pantry: Pantry;
}

export interface UpdatePantryRequest {
    UserName: string;
    Pantry: Pantry;
    ItemsToDelete: Ingredient[];
}

export interface PostPantryRequest { 
    UserName: string;
    Ingredient: Ingredient;
}

export interface SignupRequest {
    UserName:   string;
    FirstName:  string; 
    LastName:   string; 
    Email:      string; 
    Password:   string; 
}

export interface GetListRequest {
    UserName:   string; 
}

export interface PostListRequest {
    UserName: string;
    Ingredient: Ingredient;
}

export interface UpdateListRequest{
    UserName: string;
    List: List; 
    ItemsToDelete: Ingredient[];
}
