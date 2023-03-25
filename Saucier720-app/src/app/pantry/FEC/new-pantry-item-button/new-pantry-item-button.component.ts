import { Component } from '@angular/core';
import { HttpClient, HttpRequest } from '@angular/common/http';

@Component({
  selector: 'app-new-pantry-item-button',
  template: '<button (click)="postPantryItem()">Submit</button>',
})
export class NewPantryItemButtonComponent {
  pantryPostUrl = 'http://localhost:8082/api/NewPantryItem'

  constructor(private http: HttpClient) { }

  postPantryItem() {
    const req = new HttpRequest('POST', this.pantryPostUrl, { 
      reportProgress: true
    });

  }
}
