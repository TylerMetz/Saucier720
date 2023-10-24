import { Component, OnInit } from '@angular/core';
import { Ingredient, List } from '../core/interfaces/ingredient';
import { ListService } from '../core/services/list/list.service';
import { HttpClient } from '@angular/common/http';
import { HttpEvent, HttpEventType } from "@angular/common/http"
import { lastValueFrom } from 'rxjs';
import { AuthService } from '../core/services/Auth/auth.service';
import { GetListRequest, PostListRequest } from '../core/interfaces/types';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit {
  name: string = '';
  quantity: number = 0;

  list: List = {
    Ingredients: []
  }

  constructor(private listService: ListService, private http: HttpClient, private authService: AuthService) { }

  async ngOnInit() {
    await this.populateList();
  }

  // public async populateList(): Promise<void> {
  //   try {
  //     const event: HttpEvent<any> = await lastValueFrom(this.listService.getList());
  //     switch(event.type) {
  //       case HttpEventType.Sent:
  //         console.log('Request sent!');
  //         break;
  //       case HttpEventType.ResponseHeader:
  //         console.log('Response header received!');
  //         break;
  //       case HttpEventType.DownloadProgress:
  //         const kbLoaded = Math.round(event.loaded / 1024);
  //         console.log(`Download in progress! ${kbLoaded}Kb loaded`);
  //         break;
  //       case HttpEventType.Response:
  //         console.log('Done!', event.body);
  //         this.ingredients = event.body;
  //         break;
  //     }
  //   } catch (error) {
  //     console.error(error);
  //   }
  // }

  public async populateList(): Promise<void> {
    const request: GetListRequest = {
      UserName: this.authService.getUsername(),
    };
    console.log('request: ', request);
    this.listService.getList(request.UserName).subscribe({
      next: (response: any) => {
        console.log('GetListResponse ', response);
        this.list = response.List;
        console.log('List Updated: ', this.list);
      },
      error: (err: any) => {
        console.log(err, 'errors')
      }
    });
  }

  async postListItem() {
    if(!this.name){
      return; 
    }
    const request: PostListRequest = {
      UserName: this.authService.getUsername(),
      Ingredient: {
        Name: this.name,
        FoodType: '',
        SaleDetails: '',
        Quantity: this.quantity,
      }
    };
    this.listService.postListItem(request).subscribe({
      next: (response: any) => {
        console.log('New Shopping List Item Posted: ', response)
        this.name = '';
      },
      error: (err: any) => {
        console.log(err, 'errors')
      }
    });
  }

  // adjustQuantity(ingredient: Ingredient, action: string) {
  //   if (action === 'increment') {
  //     ingredient.Quantity += 1;
  //   } else if (action === 'decrement' && ingredient.Quantity > 0) {
  //     ingredient.Quantity -= 1;
  //   }
  // }

  // deleteIngredient(ingredient: Ingredient) {
  //   const index = this.ingredients.indexOf(ingredient);
  //   if (index > -1) {
  //     this.ingredients.splice(index, 1);
  //   }
  // }

  // // Optional parameter so that we can call it from deals and recipes pages 
  // addIngredient(ingredient?: Ingredient | string) {

  //   // Ingredient is not in list 
  //   let newIngredient: Ingredient | null = null; 
    
  //   // Assigns values from already created foodItem
  //   if (typeof ingredient === 'string') {
  //     newIngredient = {
  //       Name: ingredient,
  //       FoodType: '',
  //       SaleDetails: '',
  //       Quantity: 1,
  //     };
  //   } else if(ingredient){
  //     newIngredient = ingredient;
  //     newIngredient.Quantity = 1;
  //   } else {
  //     // Assigns values from list page
  //     if (this.newIngredientName && this.newIngredientQuantity > 0) {
  //       newIngredient = {
  //         Name: this.newIngredientName,
  //         Quantity: this.newIngredientQuantity,
  //         FoodType: '',
  //         SaleDetails: '',
  //       };

  //     // Clear input fields
  //     this.newIngredientName = '';
  //     this.newIngredientQuantity = 0;      
  //     }
  //   }

  //   // Checks if not null 
  //   if(newIngredient){
  //           // post new list item to backend
  //           this.postList(newIngredient)
  //   }
  // }

  // async validateIngredient(ingredient: Ingredient): Promise<boolean> {
  //   //console.log("Checking " + ingredient.Name);
  //   const response = await this.listService.checkIfExists(ingredient)
  //   if (response){
  //     //console.log(ingredient.Name + " was found in list!")
  //   } else {
  //     //console.log(ingredient.Name + " was not found in list!")
  //   }
  //   return response;
  // }

  // async postList(ingredient: Ingredient) {
  //   try {
  //     const response = await lastValueFrom(this.listService.postListItem(ingredient));
  //     console.log(response);
  //     this.populateList();
  //   } catch (error) {
  //     console.error(error);
  //   }
  // }
}
