import { Injectable } from '@angular/core';
import { HttpClient, HttpParams, HttpHeaders } from '@angular/common/http';
import { Recipe } from 'src/app/core/interfaces/recipe';
import { GetPantryRequest, GetRecipesRequest, PostRecipeRequest } from '../../interfaces/types';
import MealDealzRoutes from '../../interfaces/routes';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RecipeService {
  private userRecipeUrl = 'http://localhost:8082/api/UserRecipesSelect';
  private favoriteRecipeUrl = 'http://localhost:8082/api/FavoriteRecipesSelect';
  private recommendedRecipeUrl = 'http://localhost:8082/api/RecommendedRecipesSelect';
  private newRecipeUrl = 'http://localhost:8082/api/NewUserRecipe';
  private addFavoriteRecipeUrl = 'http://localhost:8082/api/AddFavoriteRecipe';
  private removeFavoriteRecipeUrl = 'http://localhost:8082/api/RemoveFavoriteRecipe';
  private filtersUrl = 'http://localhost:8082/api/RecommendedRecipesFilters';
  private deleteUserRecipeUrl = 'http://localhost:8082/api/DeleteUserRecipe';

  constructor(private http: HttpClient) { }

  getRecipes(request: GetRecipesRequest): Observable<any> {
    console.log('recipe req username ',request.UserName)
    console.log('recipe req filter', request.RecipeFilter)

    const options = (request.UserName) ?
    { params: new HttpParams().set('username', request.UserName) } : {};

    options.params = options.params?.append('self', request.RecipeFilter.SelfCreatedRecipes)
    options.params = options.params?.append('mdRecipes', request.RecipeFilter.MealDealzRecipes)
    options.params = options.params?.append('others', request.RecipeFilter.UserCreatedRecipes)
    console.log('recipes get options', options)
    return this.http.get<any>(MealDealzRoutes.getRecipesUrl, options);
  }

  postNewRecipe(request: PostRecipeRequest){
    console.log('LoginRequest', request)
    return this.http.post(MealDealzRoutes.postRecipesUrl, request, { observe: 'response', responseType: 'json', withCredentials: true });
  }

  postFavoriteRecipe(recipeID: number){
    const headers = new HttpHeaders({ 
      'Content-Type': 'application/json', 
    });
    return this.http.post<any>(this.addFavoriteRecipeUrl, recipeID, { headers, withCredentials: true });
  }

  postRemoveFavoriteRecipe(recipeID: number){ //delete favorite recipe 
    const headers = new HttpHeaders({ 
      'Content-Type': 'application/json', 
    });
    return this.http.post<any>(this.removeFavoriteRecipeUrl, recipeID, { headers, withCredentials: true });
  }

  postFilterValues(filterValues: any){
    const headers = new HttpHeaders({ 
      'Content-Type': 'application/json', 
    });
    return this.http.post<any>(this.filtersUrl, filterValues, { headers, withCredentials: true });
  }

  postDeleteUserRecipe(recipeID: number){
    const headers = new HttpHeaders({ 
      'Content-Type': 'application/json', 
    });
    return this.http.post<any>(this.deleteUserRecipeUrl, recipeID, { headers, withCredentials: true });
  }
}

