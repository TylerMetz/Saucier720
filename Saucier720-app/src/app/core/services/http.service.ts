import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  constructor(private http: HttpClient) { }

  // GET method
  get(url: string): Observable<any> {
    return this.http.get(url);
  }

  // POST method
  post(url: string, data: any): Observable<any> {
    return this.http.post(url, data);
  }

  // PUT method
  put(url: string, data: any): Observable<any> {
    return this.http.put(url, data);
  }

  // DELETE method
  delete(url: string): Observable<any> {
    return this.http.delete(url);
  }
}

