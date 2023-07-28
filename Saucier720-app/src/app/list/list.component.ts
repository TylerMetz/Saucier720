import { Component, OnInit } from '@angular/core';
import { Ingredient } from '../core/interfaces/ingredient';
import { ListService } from '../core/services/list/list.service';
import { HttpClient } from '@angular/common/http';
import { HttpEvent, HttpEventType } from "@angular/common/http"
import { lastValueFrom } from 'rxjs';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit {
  ingredients: Array<Ingredient> = [];
  newIngredientName: string = '';
  newIngredientQuantity: number = 0;

  constructor(private listService: ListService, private http: HttpClient) { }

  async ngOnInit() {
    await this.populateList();
  }

  public async populateList(): Promise<void> {
    try {
      const event: HttpEvent<any> = await lastValueFrom(this.listService.getList());
      switch(event.type) {
        case HttpEventType.Sent:
          console.log('Request sent!');
          break;
        case HttpEventType.ResponseHeader:
          console.log('Response header received!');
          break;
        case HttpEventType.DownloadProgress:
          const kbLoaded = Math.round(event.loaded / 1024);
          console.log(`Download in progress! ${kbLoaded}Kb loaded`);
          break;
        case HttpEventType.Response:
          console.log('Done!', event.body);
          this.ingredients = event.body;
          break;
      }
    } catch (error) {
      console.error(error);
    }
  }

  adjustQuantity(ingredient: Ingredient, action: string) {
    if (action === 'increment') {
      ingredient.Quantity += 1;
    } else if (action === 'decrement' && ingredient.Quantity > 0) {
      ingredient.Quantity -= 1;
    }
  }

  deleteIngredient(ingredient: Ingredient) {
    const index = this.ingredients.indexOf(ingredient);
    if (index > -1) {
      this.ingredients.splice(index, 1);
    }
  }

  // Optional parameter so that we can call it from deals and recipes pages 
  addIngredient(ingredient?: Ingredient | string) {

    // Check if ingredient exists in userList already
    // create function validateIngredient 



    // Ingredient is not in list 
    let newIngredient: Ingredient | null = null; 
    
    // Assigns values from already created foodItem
    if (typeof ingredient === 'string') {
      newIngredient = {
        Name: ingredient,
        Quantity: 1,
        StoreCost: 0, // Example value, replace with actual value if needed
        OnSale: false, // Example value, replace with actual value if needed
        SalePrice: 0, // Example value, replace with actual value if needed
        SaleDetails: '' // Example value, replace with actual value if needed
      };
    } else if(ingredient){
      newIngredient = ingredient;
      newIngredient.Quantity = 1;
    } else {
      // Assigns values from list page
      if (this.newIngredientName && this.newIngredientQuantity > 0) {
        newIngredient = {
          Name: this.newIngredientName,
          Quantity: this.newIngredientQuantity,
          StoreCost: 0, // Example value, replace with actual value if needed
          OnSale: false, // Example value, replace with actual value if needed
          SalePrice: 0, // Example value, replace with actual value if needed
          SaleDetails: '' // Example value, replace with actual value if needed
        };

      // Clear input fields
      this.newIngredientName = '';
      this.newIngredientQuantity = 0;      
      }
    }

    // Checks if not null 
    if(newIngredient){
            // post new list item to backend
            this.postList(newIngredient)
    }
  }

  validateIngredient(ingredient: Ingredient){
    console.log("Checking " + ingredient.Name);
  }

  async postList(ingredient: Ingredient) {
    try {
      const response = await lastValueFrom(this.listService.postListItem(ingredient));
      console.log(response);
     //if(window.location.href == "http://localhost:4200/List"){
     //   window.location.reload();
     //}
     window.location.reload();
    } catch (error) {
      console.error(error);
    }
  }
}
