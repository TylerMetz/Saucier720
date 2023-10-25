import { Injectable } from '@angular/core';
import { HttpClient, HttpRequest, HttpHeaders, HttpEventType, HttpEvent, HttpResponse, HttpParams } from '@angular/common/http';
import { Ingredient, List } from '../../interfaces/ingredient';
import { CookieService } from 'ngx-cookie-service';
import { from, lastValueFrom, Observable } from 'rxjs';
import MealDealzRoutes from '../../interfaces/routes';
import { PostListRequest, UpdateListRequest } from '../../interfaces/types';
import { AuthService } from '../Auth/auth.service';

@Injectable({
  providedIn: 'root'
})
export class ListService {
  private listUrl = 'http://localhost:8080/api/List';
  private postListUrl = 'http://localhost:8082/api/NewItem'

  constructor(private http: HttpClient, private cookieService: CookieService, private authService: AuthService) { }

  getList(username: string): Observable<List> {
    console.log('username: ', username);
    const options = username ?
    { params: new HttpParams().set('username', username)} : {};
    console.log(options);
    return this.http.get<List>(MealDealzRoutes.getListUrl, options);
  }

  postListItem(request: PostListRequest) {
    console.log('request', request)
    return this.http.post<any>(MealDealzRoutes.postListUrl, request, { observe: 'response', responseType: 'json', withCredentials: true });
  }

  updateList(request: UpdateListRequest) {
    console.log('UpdateListRequest', request)
    return this.http.put<any>(MealDealzRoutes.updateListUrl, request, { observe: 'response', responseType: "json", withCredentials: true});
  }


  }
