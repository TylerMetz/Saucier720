import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

let mockUser: any = {
  "FirstName":"Eddie",
  "LastName":"Menello",
  "Email":"Edward@gmail.com",
  "UserName":"Eddiefye69",
  "Password":"ILoveGraham420",
};

@Injectable({
  providedIn: 'root'
})
export class SignupService {

  private signupUrl: string = "api/Signup"

  constructor(private http: HttpClient) {}

  public signup(): Observable<any>{
    const headers = new HttpHeaders({ 'Content-Type': 'application/json' });
    const body = { mockUser };
    return this.http.post<any>(this.signupUrl, body, { headers, withCredentials: true });
  }
}
