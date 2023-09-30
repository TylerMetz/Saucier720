import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders } from '@angular/common/http';
import { Ingredient } from '../../interfaces/ingredient';
import { Store } from '../../interfaces/store';
import { DEALS } from 'src/app/mocks/deals.mock';
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class DealsService {
  dealsUrl = 'http://localhost:8081/api/Deals'
  private storePostUrl = 'http://localhost:8082/api/DealsStore';

  constructor(private http: HttpClient, private cookieService: CookieService) { }

  getDeals() {
    const req = new HttpRequest('GET', this.dealsUrl, { 
      reportProgress: true
    });
    
    return this.http.request(req);
  }

  postStore(store: Store) {
    const headers = new HttpHeaders({ 
      'Content-Type': 'application/json', 
    });
  
    const body = { store };
    console.log(body)
    return this.http.post<any>(this.storePostUrl, body, { headers, withCredentials: true });
  }

  getMockDeals(): Array<Ingredient> {
    return DEALS;
  }
}
