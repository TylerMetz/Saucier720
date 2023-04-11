import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders } from '@angular/common/http';
import { Observable, tap } from 'rxjs';
import { CookieService } from 'ngx-cookie-service';

@Injectable()
export class AuthService {

  private loggedIn: boolean = false;


  private loginUrl: string = 'http://localhost:8084/api/Login';
  private logoutUrl: string = 'api/Logout'; // we don't have one yet

  constructor(private http: HttpClient, private cookieService: CookieService) { }

  public login(username: string, password: string): Observable<any> {
    this.loggedIn = true;
    const body = { username, password };
    return this.http.post(this.loginUrl, body, { observe: 'response', responseType: 'text', withCredentials: true });
  }
  

  public logout(): Observable<any> {
    return this.http.post(this.logoutUrl, {});
    this.loggedIn = false;
  }

  public isLoggedIn(): boolean {
    return this.loggedIn;
  }
}
