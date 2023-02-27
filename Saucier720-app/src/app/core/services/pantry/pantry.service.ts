import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Ingredient } from '../../interfaces/ingredient';
import { PANTRY } from 'src/app/mocks/ingredients.mock';

@Injectable({
  providedIn: 'root'
})
export class PantryService {
  pantryUrl = 'http://localhost:8080/'

  constructor(private http: HttpClient) { }

  getPantry() {
    return this.http.get<Array<Ingredient>>(this.pantryUrl);
  }

  getMockPantry(): Array<Ingredient> {
    return PANTRY;
  }
}
