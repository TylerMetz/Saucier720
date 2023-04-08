import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders } from '@angular/common/http';
import { Ingredient } from '../../interfaces/ingredient';
import { PANTRY } from 'src/app/mocks/pantry.mock';

@Injectable({
  providedIn: 'root'
})
export class PantryService {
  pantryUrl = 'http://localhost:8080/api/Pantry';
  pantryPostUrl = 'http://localhost:8083/api/NewPantryItem';

  constructor(private http: HttpClient) { }

  getPantry() {
    const req = new HttpRequest('GET', this.pantryUrl, { 
      reportProgress: true
    });
    
    return this.http.request(req);
  }

  postPantryItem(itemData: Ingredient) {
    const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
    const body = { itemData };
    console.log(body)
    return this.http.post<any>(this.pantryPostUrl, body, { headers, withCredentials: true });
  }

  getMockPantry(): Array<Ingredient> {
    return PANTRY;
  }
}
