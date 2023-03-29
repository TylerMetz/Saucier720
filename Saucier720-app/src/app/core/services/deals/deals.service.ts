import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest } from '@angular/common/http';
import { Ingredient } from '../../interfaces/ingredient';
import { PANTRY } from 'src/app/mocks/pantry.mock';

@Injectable({
  providedIn: 'root'
})
export class DealsService {
  dealsUrl = 'http://localhost:8081/api/Deals'

  constructor(private http: HttpClient) { }

  getDeals() {
    const req = new HttpRequest('GET', this.dealsUrl, { 
      reportProgress: true
    });
    
    return this.http.request(req);
  }

  getMockPantry(): Array<Ingredient> {
    return PANTRY;
  }
}
