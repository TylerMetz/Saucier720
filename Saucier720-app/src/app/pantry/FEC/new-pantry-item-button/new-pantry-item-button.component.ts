import { Component } from '@angular/core';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';
import { PANTRY } from 'src/app/mocks/pantry.mock';
import { lastValueFrom } from 'rxjs';
import { Ingredient } from 'src/app/core/interfaces/ingredient';

@Component({
  selector: 'app-new-pantry-item-button',
  templateUrl: './new-pantry-item-button.component.html',
  styleUrls:[ './new-pantry-item-button.component.scss'],
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
    if(!this.name){
      return; // Won't post if empty 
    }
    const newPantryItem: Ingredient = {
      Name: this.name,
      StoreCost: this.storeCost,
      OnSale: this.onSale,
      SalePrice: this.salePrice,
      SaleDetails: this.saleDetails,
      Quantity: this.quantity,
    };
    console.log(newPantryItem)
    try {
      const response = await lastValueFrom(this.pantryService.postPantryItem(newPantryItem));
      console.log(response);
      window.location.reload();
    } catch (error) {
      console.error(error);
    }
  }
}
