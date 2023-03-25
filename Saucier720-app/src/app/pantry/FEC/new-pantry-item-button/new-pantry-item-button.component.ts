import { Component } from '@angular/core';
import { HttpClient, HttpEvent } from '@angular/common/http';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';


@Component({
  selector: 'app-new-pantry-item-button',
  template: '<button (click)="postPantryItem()">Post</button>',
  providers: [PantryService]
})

export class NewPantryItemButtonComponent {
  pantryPostUrl = 'http://localhost:8082/api/NewPantryItem';

  constructor(private pantryService: PantryService) { }

postPantryItem() {
  const itemData = {
    name: 'New Item',
    quantity: 1,
    category: 'Other'
  };
  this.pantryService.postPantryItem(itemData)
    .subscribe(
      response => {
        console.log(response);
      },
      error => {
        console.error(error);
      }
    );
}

}
