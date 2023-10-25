import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders, HttpParams } from '@angular/common/http';
import { Deals, Ingredient } from '../../interfaces/ingredient';
import { Store } from '../../interfaces/store';
import { DEALS } from 'src/app/mocks/deals.mock';
import { CookieService } from 'ngx-cookie-service';
import { Observable } from 'rxjs';
import MealDealzRoutes from '../../interfaces/routes';

@Injectable({
  providedIn: 'root'
})
export class DealsService {
  dealsUrl = 'http://localhost:8081/api/Deals'
  private storePostUrl = 'http://localhost:8082/api/DealsStore';

  constructor(private http: HttpClient, private cookieService: CookieService) { }

  getDeals(store: string): Observable<Deals> {
    console.log('store: ', store);
    const options = store ?
    { params: new HttpParams().set('store', store)} : {};
    console.log('deals request store: ', store);
    console.log('deals request options: ', options);
    return this.http.get<Deals>(MealDealzRoutes.getDealsbyStoreUrl, options);
  }
}
