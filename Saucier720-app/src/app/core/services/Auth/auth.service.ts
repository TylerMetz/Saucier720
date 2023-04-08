import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable()
export class AuthService {

  private loggedIn: boolean = true;

  private loginUrl: string = 'api/Login';
  private logoutUrl: string = 'api/logout'; //we dont have one yet

  constructor(private http: HttpClient) { }

  public login(username: string, password: string): Observable<any> {
    const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
    const body = { username, password };
    this.loggedIn = true;
    return this.http.post<any>(this.loginUrl, body, { headers, withCredentials: true });
  }

  public logout(): Observable<any> {
    return this.http.post(this.logoutUrl, {});
    
    this.loggedIn = false;
  }

  public isLoggedIn(): boolean {
    return this.loggedIn;
  }

}
