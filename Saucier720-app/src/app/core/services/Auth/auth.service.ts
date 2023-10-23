import { Injectable, OnInit } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { CookieService } from 'ngx-cookie-service';
import MealDealzRoutes from '../../interfaces/routes';
import { LoginRequest } from '../../interfaces/types';

@Injectable()
export class AuthService {

  public loggedIn: boolean = false;
  private validCookie: boolean = false;


  constructor(private http: HttpClient, private cookieService: CookieService) {
    this.validCookie = this.cookieService.check('sessionID');
    console.log('cookie status: ', this.validCookie);
    if (this.validCookie) {
      this.loggedIn = true;
    }
   }
   

  public login(request: LoginRequest): Observable<any> {
    const body = request;
    console.log(body)
    return this.http.post(MealDealzRoutes.loginUrl, body, { observe: 'response', responseType: 'json', withCredentials: true });
  }

  public logout(): Observable<any> {
    this.loggedIn = false;
    this.cookieService.delete('sessionID');
    console.log("post req sending");
    return this.http.post(MealDealzRoutes.logoutUrl, { });
  }

  public isLoggedIn(): boolean {
    return this.loggedIn;
  }
}
