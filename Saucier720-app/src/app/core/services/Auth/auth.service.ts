import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders } from '@angular/common/http';
import { Observable, tap } from 'rxjs';
import { CookieService } from 'ngx-cookie-service';

@Injectable()
export class AuthService {

  private loggedIn: boolean = true;

  private loginUrl: string = 'http://localhost:8084/api/Login';
  private logoutUrl: string = 'api/Logout'; //we dont have one yet

  constructor(private http: HttpClient, private cookieService: CookieService) { }

  public login(username: string, password: string): Observable<any> {
    const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
    const body = { username, password };
    return this.http.post<any>(this.loginUrl, body, { headers, withCredentials: true, observe: 'response' });
  }
  

  public logout(): Observable<any> {
    return this.http.post(this.logoutUrl, {});
    
    this.loggedIn = false;
  }

  public isLoggedIn(): boolean {
    return this.loggedIn;
  }

}
