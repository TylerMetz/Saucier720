import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders } from '@angular/common/http';
import { Recipe } from 'src/app/core/interfaces/recipe';

@Injectable({
  providedIn: 'root'
})
export class RecipeService {
  private recipeUrl = 'http://localhost:8080/api/Recipes';
  private userRecipeUrl = 'http://localhost:8082/api/UserRecipesSelect';
  private favoriteRecipeUrl = 'http://localhost:8082/api/FavoriteRecipesSelect';
  private recommendedRecipeUrl = 'http://localhost:8082/api/RecommendedRecipesSelect';
  private newRecipeUrl = 'http://localhost:8082/api/NewUserRecipe';

  constructor(private http: HttpClient) { }

  getRecipes() {
    const req = new HttpRequest('GET', this.recipeUrl, { 
      reportProgress: true
    });
    
    return this.http.request(req);
  }

  postFavoriteRecipesSelect() {
    const headers = new HttpHeaders({ 
      'Content-Type': 'application/json', 
    });
    return this.http.post<any>(this.favoriteRecipeUrl, "", { headers, withCredentials: true });
  }

  postMyRecipesSelect() {
    const headers = new HttpHeaders({ 
      'Content-Type': 'application/json', 
    });
    return this.http.post<any>(this.userRecipeUrl, "", { headers, withCredentials: true });
  }

  postRecommendedRecipesSelect() {
    const headers = new HttpHeaders({ 
      'Content-Type': 'application/json', 
    });
    return this.http.post<any>(this.recommendedRecipeUrl, "", { headers, withCredentials: true });
  }

  postNewRecipe(recipe: Recipe){
    const headers = new HttpHeaders({ 
      'Content-Type': 'application/json', 
    });
    return this.http.post<any>(this.newRecipeUrl, recipe, { headers, withCredentials: true });
  }

}

