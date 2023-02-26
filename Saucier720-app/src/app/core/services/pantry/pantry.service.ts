import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Ingredient } from '../../interfaces/ingredient';

@Injectable({
  providedIn: 'root'
})
export class PantryService {
  pantryUrl = 'Pantry/getPantry'

  constructor(private http: HttpClient) { }

  getPantry() {
    return this.http.get<Ingredient[]>(this.pantryUrl);
  }
}
