import { HttpEvent, HttpEventType } from "@angular/common/http"
import { Component, OnInit, Input } from '@angular/core';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';
import { lastValueFrom } from 'rxjs';
import { Ingredient, Pantry } from "src/app/core/interfaces/ingredient";
import { GetPantryRequest, UpdatePantryRequest } from "src/app/core/interfaces/types";
import { AuthService } from "src/app/core/services/Auth/auth.service";

@Component({
  selector: 'app-pantry-table',
  providers: [PantryService],
  templateUrl: './pantry-table.component.html',
  styleUrls: ['./pantry-table.component.scss']
})
export class PantryTableComponent implements OnInit {

  pantry: Pantry = {
    Ingredients: [],
  };

  constructor(private pantryService: PantryService, private authService: AuthService) { }

  async ngOnInit(){
    await this.populatePantry();
  }

  public async populatePantry(): Promise<void> {
    console.log('username: ', this.authService.getUsername())
    this.pantryService.getPantry(this.authService.getUsername()).subscribe({
      next: (response: any) => {
        console.log('GetPantryResponse: ', response)
        this.pantry = response.Pantry;
        console.log('pantry updated: ', this.pantry)
      },
      error: (err: any) => {
        console.log(err, 'errors')
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
    console.log(response);
  } 

  addTempValue(name: string, quantity: number){
    const newIngredient: Ingredient = {
      Name: name, 
      FoodType: '',
      SaleDetails: '',
      Quantity: quantity,
    };
    if(this.pantry.Ingredients === null){
      this.pantry.Ingredients = [];
    }
    this.pantry.Ingredients.push(newIngredient);
  }
}
