import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from './core/interfaces/user';

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

  private signupUrl: string = 'http://localhost:8081/api/Signup';

  constructor(private http: HttpClient) {}

  public signup(user: User): Observable<any>{
    const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
    const body = { user };
    console.log(body)
    return this.http.post<any>(this.signupUrl, body, { headers, withCredentials: true });
  }
}
