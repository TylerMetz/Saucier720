import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders } from '@angular/common/http';
import { Recipe } from 'src/app/core/interfaces/recipe';

@Injectable({
  providedIn: 'root'
})
export class UserDashboardService {
  private userDashboardDataUrl = 'http://localhost:8080/api/UserDashboard';

  constructor(private http: HttpClient) { }

  getUserDashboardData() {
    const req = new HttpRequest('GET', this.userDashboardDataUrl, { 
      reportProgress: true
    });
    
    return this.http.request(req);
  }

}

