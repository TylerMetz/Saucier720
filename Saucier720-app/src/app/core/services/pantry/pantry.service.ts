import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders } from '@angular/common/http';
import { Ingredient } from '../../interfaces/ingredient';
import { PANTRY } from 'src/app/mocks/pantry.mock';

@Injectable({
  providedIn: 'root'
})
export class PantryService {
  private pantryUrl = 'http://localhost:8080/api/Pantry';
  private pantryPostUrl = 'http://localhost:8083/api/NewPantryItem';
  private cookiesPostUrl = 'http://localhost:8083/api/cookies'
  private pantryUpdateUrl = 'http://localhost:8086/UpdatePantry'

  constructor(private http: HttpClient) { }

  getPantry() {
    const req = new HttpRequest('GET', this.pantryUrl, { 
      reportProgress: true
    });
    
    return this.http.request(req);
  }

  updatePantry(pantry: Ingredient[]) {
    const headers = new HttpHeaders({ 
      'Content-Type': 'application/json', 
      'Cookie': document.cookie // Set the cookie value in the header
    });
  
    const body = { pantry };
    console.log(body)
    return this.http.post<any>(this.pantryPostUrl, body, { headers, withCredentials: true });
  }

  postPantryItem(ingredient: Ingredient) {
    const headers = new HttpHeaders({ 
      'Content-Type': 'application/json', 
      'Cookie': document.cookie // Set the cookie value in the header
    });
  
    const body = { ingredient };
    console.log(body)
    return this.http.post<any>(this.pantryUpdateUrl, body, { headers, withCredentials: true });
  }


  getMockPantry(): Array<Ingredient> {
    return PANTRY;
  }
}
