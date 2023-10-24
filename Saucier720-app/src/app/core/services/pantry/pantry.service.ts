import { Injectable } from '@angular/core';
import { HttpClient, HttpParams, HttpHeaders } from '@angular/common/http';
import { Ingredient, Pantry } from '../../interfaces/ingredient';
import { PANTRY } from 'src/app/mocks/pantry.mock';
import { CookieService } from 'ngx-cookie-service';
import MealDealzRoutes from '../../interfaces/routes';
import { GetPantryRequest, PostPantryRequest, UpdatePantryRequest } from '../../interfaces/types';
import { AuthService } from '../Auth/auth.service';
import { catchError, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class PantryService {
  private pantryPostUrl = 'http://localhost:8082/api/NewPantryItem';
  private pantryUpdateUrl = 'http://localhost:8082/api/UpdatePantry';

  constructor(private http: HttpClient, 
    private cookieService: CookieService,
    private authService: AuthService) { }

  getPantry(username: string): Observable<Pantry> {
    console.log('username: ', username);
    const options = username ?
    { params: new HttpParams().set('username', username) } : {};
    console.log('pantry request username: ', username);
    console.log('pantry request options: ', options);
    return this.http.get<Pantry>(MealDealzRoutes.getPantryUrl, options);
  }

  updatePantry(request: UpdatePantryRequest) {
    console.log('UpdatePantryRequest', request)
    return this.http.put<any>(MealDealzRoutes.updatePantryUrl, request, { observe: 'response', responseType: 'json', withCredentials: true });
  }

  postPantryItem(request: PostPantryRequest) {
    console.log('request', request)
    return this.http.post<any>(MealDealzRoutes.postPantryUrl, request, { observe: 'response', responseType: 'json', withCredentials: true });
  }


  getMockPantry(): Array<Ingredient> {
    return PANTRY;
  }
}
