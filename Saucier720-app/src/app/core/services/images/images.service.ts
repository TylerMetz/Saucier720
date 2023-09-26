import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ImagesService {
  private apiKey = 'AIzaSyDTd_pNG1PpPNrNPl-DzSqQf677hhZgzhY';
  private cx = '841ce4d9e2cf24591'; 

  constructor(private http: HttpClient) {}

  searchImage(recipeTitle: string) {
    const query = encodeURIComponent(recipeTitle);
    const apiUrl = `https://www.googleapis.com/customsearch/v1?q=${query}&cx=${this.cx}&key=${this.apiKey}&searchType=image`;

    return this.http.get(apiUrl);
  }
}
