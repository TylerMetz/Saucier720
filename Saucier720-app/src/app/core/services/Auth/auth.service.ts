import { Injectable, OnInit } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders } from '@angular/common/http';
import { Observable, first } from 'rxjs';
import { CookieService } from 'ngx-cookie-service';
import MealDealzRoutes from '../../interfaces/routes';
import { LoginRequest, LogoutRequest } from '../../interfaces/types';

@Injectable()
export class AuthService {

  public loggedIn: boolean = false; 
  private validCookie: boolean = false;


  constructor(private http: HttpClient, private cookieService: CookieService) {
    this.validCookie = this.cookieService.check('Cookie');
    console.log('cookie status: ', this.validCookie);
    if (this.validCookie) {
      this.loggedIn = true;
    }
   }

   public getUsername(): string {
      console.log('cookie value: ', this.cookieService.get('Cookie'))
      const firstDash = this.cookieService.get('Cookie').indexOf('-');
      return this.cookieService.get('Cookie').slice(1, firstDash); // If there's no dash, return the whole value
   }
   

  public login(request: LoginRequest): Observable<any> {
    const body = request;
    console.log('LoginRequest', body)
    return this.http.post(MealDealzRoutes.loginUrl, body, { observe: 'response', responseType: 'json', withCredentials: true });
  }

  public logout(request: LogoutRequest): Observable<any> {
    this.loggedIn = false;
    this.cookieService.delete('Cookie');
    console.log('cookie deleted');
    console.log("post req sending");
    const body = request;
    console.log('LogoutRequest', body)
    return this.http.post(MealDealzRoutes.logoutUrl, body, { observe: 'response', responseType: 'json', withCredentials: true });
  }

  public isLoggedIn(): boolean {
    return this.loggedIn;
  }
}
