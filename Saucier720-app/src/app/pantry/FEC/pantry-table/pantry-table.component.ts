import { HttpEvent, HttpEventType } from "@angular/common/http"
import { Component, OnInit, Input } from '@angular/core';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';
import { lastValueFrom } from 'rxjs';
import { Ingredient } from "src/app/core/interfaces/ingredient";
import { GetPantryRequest } from "src/app/core/interfaces/types";
import { AuthService } from "src/app/core/services/Auth/auth.service";

@Component({
  selector: 'app-pantry-table',
  providers: [PantryService],
  templateUrl: './pantry-table.component.html',
  styleUrls: ['./pantry-table.component.scss']
})
export class PantryTableComponent implements OnInit {

  pantry: Ingredient[] = [];

  constructor(private pantryService: PantryService, private authService: AuthService) { }

  async ngOnInit(){
    await this.populatePantry();
  }

  public async populatePantry(): Promise<void> {
    const request: GetPantryRequest = {
      UserName: this.authService.getUsername(),
    };
    console.log('request: ', request)
    this.pantryService.getPantry(request.UserName).subscribe({
      next: (response: any) => {
        console.log('GetPantryResponse: ', response)
        this.pantry = response;
        console.log('pantry updated: ', this.pantry)
      },
      error: (err: any) => {
        console.log(err, 'errors')
      }
    });
  }

  async updatePantry() {
    try {
      // Check and remove items with quantity 0
      this.pantry = this.pantry.filter((item: Ingredient) => item.Quantity !== 0);
  
      // Call pantryService to update pantry
      const response = await lastValueFrom(this.pantryService.updatePantry(this.pantry));
      console.log(response);
    } catch (error) {
      console.error(error);
    }
  } 

  addTempValue(name: string, quantity: number){
    const newIngredient: Ingredient = {
      Name: name, 
      FoodType: '',
      SaleDetails: '',
      Quantity: quantity,
    };
    if(this.pantry === null){
      this.pantry = []
    }
    this.pantry.push(newIngredient)
  }

  

}
