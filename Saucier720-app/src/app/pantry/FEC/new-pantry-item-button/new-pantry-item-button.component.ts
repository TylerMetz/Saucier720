import { Component, Output, EventEmitter } from '@angular/core';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';
import { PANTRY } from 'src/app/mocks/pantry.mock';
import { lastValueFrom } from 'rxjs';
import { Ingredient } from 'src/app/core/interfaces/ingredient';
import { PantryTableComponent } from '../pantry-table/pantry-table.component';
import { PostPantryRequest } from 'src/app/core/interfaces/types';
import { AuthService } from 'src/app/core/services/Auth/auth.service';

@Component({
  selector: 'app-new-pantry-item-button',
  templateUrl: './new-pantry-item-button.component.html',
  styleUrls:[ './new-pantry-item-button.component.scss'],
  providers: [PantryService, PantryTableComponent]
})

export class NewPantryItemButtonComponent {
  name: string = '';
  foodtype: string = '';
  saleDetails: string = '';
  quantity: number = 100;

  @Output() newItemAdded: EventEmitter<{ name: string, quantity: number }> = new EventEmitter<{ name: string, quantity: number }>();

  constructor(private pantryService: PantryService, private authService: AuthService) { }

  async postPantryItem() {
    if(!this.name){
      return
    }
    const request: PostPantryRequest = { 
      UserName: this.authService.getUsername(),
      Ingredient: {
        Name: this.name,
        FoodType: this.foodtype,
        SaleDetails: this.saleDetails,
        Quantity: this.quantity,
      }
    };
    this.pantryService.postPantryItem(request).subscribe({
      next: (response: any) => {
        console.log('New Pantry Item Posted: ', response)
        this.name = '';
      },
      error: (err: any) => {
        console.log(err, 'errors')
      }
    });
  }
}
