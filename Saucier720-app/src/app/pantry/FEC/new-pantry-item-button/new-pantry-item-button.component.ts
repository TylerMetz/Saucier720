import { Component, Output, EventEmitter } from '@angular/core';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';
import { PANTRY } from 'src/app/mocks/pantry.mock';
import { lastValueFrom } from 'rxjs';
import { Ingredient } from 'src/app/core/interfaces/ingredient';
import { PantryTableComponent } from '../pantry-table/pantry-table.component';

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

  constructor(private pantryService: PantryService) { }

  async postPantryItem() {
    if(!this.name){
      return; // Won't post if empty 
    }
    const newPantryItem: Ingredient = {
      Name: this.name,
      FoodType: this.foodtype,
      SaleDetails: this.saleDetails,
      Quantity: this.quantity,
    };
    console.log(newPantryItem)
    try {
      const response = await lastValueFrom(this.pantryService.postPantryItem(newPantryItem));
      console.log(response);
      const newItem = {
        name: newPantryItem.Name,
        quantity: newPantryItem.Quantity,
      };
      this.newItemAdded.emit(newItem);
      this.name = '';
      this.quantity = 1;
    } catch (error) {
      console.error(error);
    }
  }
}
