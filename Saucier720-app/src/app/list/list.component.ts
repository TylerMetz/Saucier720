import { Component, OnInit } from '@angular/core';
import { Ingredient, List } from '../core/interfaces/ingredient';
import { ListService } from '../core/services/list/list.service';
import { HttpClient } from '@angular/common/http';
import { HttpEvent, HttpEventType } from "@angular/common/http"
import { lastValueFrom } from 'rxjs';
import { AuthService } from '../core/services/Auth/auth.service';
import { GetListRequest, PostListRequest, UpdateListRequest } from '../core/interfaces/types';

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

  async updateList(){
    const zeroQuantityItems = this.list.Ingredients.filter((item: Ingredient) => item.Quantity === 0); 
    this.list.Ingredients = this.list.Ingredients.filter((item: Ingredient) => item.Quantity !== 0);
    const request: UpdateListRequest = { 
      UserName: this.authService.getUsername(),
      List: this.list,
      ItemsToDelete: zeroQuantityItems,
    };
    const response = await lastValueFrom(this.listService.updateList(request));
    console.log('UpdateListResponse: ', response);
    this.populateList();
  }

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
  //   const response = await this.listService.checkIfExists(ingredient, this.authService.getUsername())
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
