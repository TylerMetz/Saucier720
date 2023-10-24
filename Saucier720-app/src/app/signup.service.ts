import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from './core/interfaces/user';
import MealDealzRoutes from './core/interfaces/routes';
import { SignupRequest } from './core/interfaces/types';

let mockUser: any = {
  "FirstName":"Riley",
  "LastName":"Cleavenger",
  "Email":"riley.cleavenger@gmail.com",
  "UserName":"RileyButterDrip",
  "Password":"bah69FantasticFour",
};

@Injectable({
  providedIn: 'root'
})
export class SignupService {

  constructor(private http: HttpClient) {}

  public signup(request: SignupRequest): Observable<any> {
    const body = request;
    console.log('SignupRequest', body);
    return this.http.post(MealDealzRoutes.signupUrl, body, { observe: 'response', responseType: 'json', withCredentials: true})
  }
}
