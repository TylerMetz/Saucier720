import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

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

  private signupUrl: string = 'http://localhost:8085/api/Signup';

  constructor(private http: HttpClient) {}

  public signup(): Observable<any>{
    const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
    const body = { mockUser };
    return this.http.post<any>(this.signupUrl, body, { headers, withCredentials: true });
  }
}
