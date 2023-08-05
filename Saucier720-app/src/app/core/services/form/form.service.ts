import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class FormService {
    private readonly scriptUrl = 'https://script.google.com/macros/s/AKfycbzZZgpZR_F3yXmMfcMyfuEL1K_r83elFxywKUW9UMAYdGw6apq09I1PVXbsKrNg1IaZ/exec';

    constructor(private http: HttpClient) { }
  
    sendFeedback(name: string, email: string, message: string): Observable<any> {
        const params = new HttpParams()
          .set('name', name)
          .set('email', email)
          .set('message', message);
    
        const headers = new HttpHeaders()
          .set('Content-Type', 'application/x-www-form-urlencoded');
    
        return this.http.post<any>(this.scriptUrl, params.toString(), { headers: headers });
      }
  }
 
  
