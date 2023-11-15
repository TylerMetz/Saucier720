import { Component, OnInit } from '@angular/core';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';
import { AuthService } from 'src/app/core/services/Auth/auth.service';
import { Pantry, Ingredient } from 'src/app/core/interfaces/ingredient';
import { UpdatePantryRequest, PostPantryRequest } from 'src/app/core/interfaces/types';
import { lastValueFrom } from 'rxjs';

@Component({
  selector: 'app-pantry-form',
  templateUrl: './pantry-form.component.html',
  styleUrls: ['./pantry-form.component.scss']
})
export class PantryFormComponent implements OnInit {

  pantry: Pantry = {
    Ingredients: [],
  };

  name: string = '';
  foodtype: string = '';
  saleDetails: string = '';
  quantity: number = 100;

  constructor(private pantryService: PantryService, private authService: AuthService) { }

  async ngOnInit() {
    await this.populatePantry();
  }

  public async populatePantry(): Promise<void> {
    console.log('username: ', this.authService.getUsername());
    this.pantryService.getPantry(this.authService.getUsername()).subscribe({
      next: (response: any) => {
        console.log('GetPantryResponse: ', response);
        this.pantry = response.Pantry;
        console.log('pantry updated: ', this.pantry);
      },
      error: (err: any) => {
        console.log(err, 'errors');
      }
    });
  }

  async updatePantry() {
    const zeroQuantityItems = this.pantry.Ingredients.filter((item: Ingredient) => item.Quantity === 0);
    this.pantry.Ingredients = this.pantry.Ingredients.filter((item: Ingredient) => item.Quantity !== 0);
    const request: UpdatePantryRequest = {
      UserName: this.authService.getUsername(),
      Pantry: this.pantry,
      ItemsToDelete: zeroQuantityItems,
    };
    const response = await lastValueFrom(this.pantryService.updatePantry(request));
    console.log('UpdatePantryResponse: ', response);
    this.populatePantry();
  }

  addTempValue(name: string, quantity: number) {
    const newIngredient: Ingredient = {
      Name: name,
      FoodType: '',
      SaleDetails: '',
      Quantity: quantity,
    };
    if (this.pantry.Ingredients === null) {
      this.pantry.Ingredients = [];
    }
    this.pantry.Ingredients.push(newIngredient);
  }

  async postPantryItem() {
    if (!this.name) {
      return;
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
        console.log('New Pantry Item Posted: ', response);
        this.name = '';
        // You may want to refresh the pantry data after posting a new item
        this.populatePantry();
      },
      error: (err: any) => {
        console.log(err, 'errors');
      }
    });
  }
}
