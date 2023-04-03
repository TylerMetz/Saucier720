import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable()
export class AuthService {

  private loggedIn = false;

  private loginUrl = 'api/login';
  private logoutUrl = 'api/logout';

  constructor(private http: HttpClient) { }

  public login(username: string, password: string): Observable<any> {
    return this.http.post(this.loginUrl, { username, password });
    
    this.loggedIn = true;
  }

  public logout(): Observable<any> {
    return this.http.post(this.logoutUrl, {});
    
    this.loggedIn = false;
  }

  public isLoggedIn(): boolean {
    return this.loggedIn;
  }

}
