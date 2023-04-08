import { Component } from '@angular/core';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';
import { PANTRY } from 'src/app/mocks/pantry.mock';
import { lastValueFrom } from 'rxjs';
import { Ingredient } from 'src/app/core/interfaces/ingredient';

@Component({
  selector: 'app-new-pantry-item-button',
  templateUrl: './new-pantry-item-button.component.html',
  providers: [PantryService]
})

export class NewPantryItemButtonComponent {
  name: string = '';
  storeCost: number = 100;
  onSale: boolean = true;
  salePrice: number = 0;
  saleDetails: string = '';
  quantity: number = 1;

  constructor(private pantryService: PantryService) { }

  async postPantryItem() {
    const newPantryItem: Ingredient = {
      Name: this.name,
      StoreCost: this.storeCost,
      OnSale: this.onSale,
      SalePrice: this.salePrice,
      SaleDetails: this.saleDetails,
      Quantity: this.quantity,
    };
    try {
      const response = await lastValueFrom(this.pantryService.postPantryItem(newPantryItem));
      console.log(response);
    } catch (error) {
      console.error(error);
    }
  }
}
