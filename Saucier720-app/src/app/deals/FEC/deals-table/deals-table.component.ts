import { HttpClient, HttpEvent, HttpEventType} from "@angular/common/http"
import { Component, OnInit, EventEmitter, Output, Input } from '@angular/core';
import { DealsService } from 'src/app/core/services/deals/deals.service';
import { count, lastValueFrom } from "rxjs";
import { ListComponent } from "src/app/list/list.component";
import { Ingredient } from "src/app/core/interfaces/ingredient";
import { delay } from "cypress/types/bluebird";
import { forEach } from "cypress/types/lodash";
import { createUrlTreeFromSnapshot } from "@angular/router";
import { Deals } from "src/app/core/interfaces/ingredient";
import { ListService } from "src/app/core/services/list/list.service";
import { AuthService } from "src/app/core/services/Auth/auth.service";
import { PostListRequest } from "src/app/core/interfaces/types";

@Component({
  selector: 'app-deals-table',
  providers: [DealsService,ListComponent,ListService],
  templateUrl: 'deals-table.component.html',
  styleUrls: ['deals-table.component.scss']
})
export class DealsTableComponent implements OnInit {

  currentStore: string = '';
  deals: Deals = {
    Ingredients: []
  }; 

  @Output() sendButtonData: EventEmitter<string> = new EventEmitter<string>();
  @Input() selectedStore!: string;

  constructor(private dealsService: DealsService, private listComponent: ListComponent, private listService: ListService, private authService: AuthService) { }

  ngOnInit() {
    if(this.selectedStore){
      this.populateDeals(this.selectedStore)
    }
    //await this.populateDeals();
    //var count = 0;
    // for (const deal of this.pantry){
    //   const isValid = await this.listComponent.validateIngredient(deal);
    //   if(isValid){
    //     const selector = `#row` + count;
    //     const element = document.querySelector(selector) as HTMLElement
    //     if(element){
    //       this.toggleInList(element)
    //     }
    //   }
    //   ++count;
    // }
  }

  public async populateDeals(store: string): Promise<void> {
    console.log('store: ', store);
    this.dealsService.getDeals(store).subscribe({
      next: (response: any) => {
        console.log('GetDealsbyStoreResponse: ', response);
        this.deals = response.Deals
        this.listService.getList(this.authService.getUsername()).subscribe({
          next: (listResponse: any) => {
            var rowCount = 0; 
            for (const deal of this.deals.Ingredients) {
              // Iterate through the list of ingredients to check for name matches
              for (const ingredient of listResponse.List.Ingredients) {
                if (deal.Name === ingredient.Name) {
                  console.log(`'${deal.Name}' was found in your list!`);
                  const selector = `#row` + rowCount;
                  const element = document.querySelector(selector) as HTMLElement
                  if(element){
                    this.toggleInList(element)
                  }
                }
              }
            }
          },
          error: (err: any) => {
            console.log(err, 'errors')
          }
        })
        console.log(this.listService.getList(this.authService.getUsername()))
      },
      error: (err: any) => {
        console.log(err, 'errors')
      }
    });
  }

  // async populateDeals() {
  //    try {
  //     const event: HttpEvent<any> = await lastValueFrom(this.dealsService.getDeals());
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
  //         this.pantry = event.body;
  //         this.currentStore = this.pantry[0].Name
  //         //console.log(this.currentStore)
  //         this.sendButtonData.emit(this.currentStore)
  //         this.pantry.shift()
  //         break;
  //     }
  //   } catch (error) {
  //     console.error(error);
  //   }
  // }

  // // Add to shopping list 
  // addToList(ingredient: Ingredient, event: Event) {
  //   // Change button state
  //   const addBtn = event.target as HTMLElement;
  //   this.toggleInList(addBtn)
  //   // Nav to actual list function 
  //   //this.listComponent.addIngredient(ingredient);
  // }

  async postListItem(ingredient: Ingredient){
    const request: PostListRequest = {
      UserName: this.authService.getUsername(),
      Ingredient: {
        Name: ingredient.Name,
        FoodType: ingredient.FoodType,
        SaleDetails: ingredient.SaleDetails,
        Quantity: 1,
      }
    };
    this.listService.postListItem(request).subscribe({
      next: (response: any) => {
        console.log('New Shopping List Item Posted: ', response)
      },
      error: (err: any) => {
        console.log(err, 'errors')
      }
    });
  }

  toggleInList(element: HTMLElement){
    element.classList.remove("not-in-list")
    element.classList.add("in-list");
    element.title = "Already in list!"
  }

}
