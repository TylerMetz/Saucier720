import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Pantry } from '../../interfaces/pantry';

@Injectable({
  providedIn: 'root'
})
export class PantryService {
  pantryUrl = 'api/Pantry/getPantry'

  constructor(private http: HttpClient) { }

  getPantry() {
    return this.http.get<Pantry>(this.pantryUrl);
  }
}
