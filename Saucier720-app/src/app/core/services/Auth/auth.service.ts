import { Injectable, OnInit } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { CookieService } from 'ngx-cookie-service';

@Injectable()
export class AuthService {

  private loggedIn: boolean = false;
  private validCookie: boolean = false;


  private loginUrl: string = 'http://localhost:8081/api/Login';
  private logoutUrl: string = 'http://localhost:8081/api/Logout';

  constructor(private http: HttpClient, private cookieService: CookieService) {
    this.validCookie = this.cookieService.check('sessionID');
    console.log('cookie status: ', this.validCookie);
    if (this.validCookie) {
      this.loggedIn = true;
    }
   }
   

  public login(username: string, password: string): Observable<any> {
    this.loggedIn = true;
    const body = { username, password };
    return this.http.post(this.loginUrl, body, { observe: 'response', responseType: 'json', withCredentials: true });
  }
  

  public logout(): Observable<any> {
    this.loggedIn = false;
    this.cookieService.delete('sessionID');
    console.log("post req sending");
    return this.http.post(this.logoutUrl, { });
  }

  public isLoggedIn(): boolean {
    return this.loggedIn;
  }
}
