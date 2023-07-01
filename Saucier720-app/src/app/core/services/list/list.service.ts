import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders } from '@angular/common/http';
import { Ingredient } from '../../interfaces/ingredient';
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class ListService {
  private listUrl = 'http://localhost:8080/api/List';
  private postListUrl = 'http://localhost:8082/api/NewItem'

  constructor(private http: HttpClient, private cookieService: CookieService) { }

  getList() {
    const req = new HttpRequest('GET', this.listUrl, {
    reportProgress: true
  });

  return this.http.request(req);
  }

  postListItem(ingredient: Ingredient) {
    const headers = new HttpHeaders({ 
      'Content-Type': 'application/json', 
    });
  
    const body = { ingredient };
    console.log(body)
    return this.http.post<any>(this.postListUrl, body, { headers, withCredentials: true });
  }
}

