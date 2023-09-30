import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders, HttpEventType, HttpEvent, HttpResponse } from '@angular/common/http';
import { Ingredient } from '../../interfaces/ingredient';
import { CookieService } from 'ngx-cookie-service';
import { from, lastValueFrom } from 'rxjs';

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

  // Function checks against current list to see if an item is already in the user list 
  async checkIfExists(ingredient: Ingredient): Promise<boolean> {
    try {
      const response = await lastValueFrom(this.getList());
      //console.log('API Response:', response);
      if (response instanceof HttpResponse) {
        const responseBody: any = response.body;
        //console.log('API Response:', responseBody);
        if (Array.isArray(responseBody)){
          const currentList: Ingredient[] = responseBody;
          return currentList.some(foodItem => foodItem.Name === ingredient.Name);
        }
      }

      return false; 
      } catch (error) {
        console.error(error);
        return false; 
      }
    }
  }
