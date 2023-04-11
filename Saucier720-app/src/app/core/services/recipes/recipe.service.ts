import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders } from '@angular/common/http';


@Injectable({
  providedIn: 'root'
})
export class RecipeService {
  private recipeUrl = 'http://localhost:8082/api/Recipes'

  constructor(private http: HttpClient) { }

  getRecipes() {
    const req = new HttpRequest('GET', this.recipeUrl, {
    reportProgress: true
  });

  return this.http.request(req);
  }
}

