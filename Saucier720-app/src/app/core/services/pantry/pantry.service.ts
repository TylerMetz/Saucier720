import { Injectable } from '@angular/core';
import { HttpClient, HttpParams, HttpHeaders } from '@angular/common/http';
import { Ingredient, Pantry } from '../../interfaces/ingredient';
import { PANTRY } from 'src/app/mocks/pantry.mock';
import { CookieService } from 'ngx-cookie-service';
import MealDealzRoutes from '../../interfaces/routes';
import { GetPantryRequest } from '../../interfaces/types';
import { AuthService } from '../Auth/auth.service';
import { catchError, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class PantryService {
  private pantryUrl = 'http://localhost:8080/api/Pantry';
  private pantryPostUrl = 'http://localhost:8082/api/NewPantryItem';
  private pantryUpdateUrl = 'http://localhost:8082/api/UpdatePantry'

  constructor(private http: HttpClient, 
    private cookieService: CookieService,
    private authService: AuthService) { }

  getPantry(username: string): Observable<Pantry> {
    console.log('username: ', username);
    const options = username ?
    { params: new HttpParams().set('name', username) } : {};
    console.log(options);
    return this.http.get<Pantry>(MealDealzRoutes.getPantryUrl, options);
  }

  updatePantry(pantry: Ingredient[]) {
    const headers = new HttpHeaders({ 
      'Content-Type': 'application/json', 
    });
  
    const body = { pantry };
    console.log(body)
    return this.http.post<any>(this.pantryUpdateUrl, body, { headers, withCredentials: true });
  }

  postPantryItem(ingredient: Ingredient) {
    const headers = new HttpHeaders({ 
      'Content-Type': 'application/json', 
      'Cookie': this.cookieService.get('sessionID') // Set the cookie value in the header
    });
  
    const body = { ingredient };
    console.log(body)
    return this.http.post<any>(this.pantryPostUrl, body, { headers, withCredentials: true });
  }


  getMockPantry(): Array<Ingredient> {
    return PANTRY;
  }
}
