import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest } from '@angular/common/http';
import { Ingredient } from '../../interfaces/ingredient';
import { PANTRY } from 'src/app/mocks/pantry.mock';

@Injectable({
  providedIn: 'root'
})
export class PantryService {
  pantryUrl = 'http://localhost:8080/api/Pantry'

  constructor(private http: HttpClient) { }

  getPantry() {
    const req = new HttpRequest('GET', this.pantryUrl, { 
      reportProgress: true
    });
    
    return this.http.request(req);
  }

  getMockPantry(): Array<Ingredient> {
    return PANTRY;
  }
}
