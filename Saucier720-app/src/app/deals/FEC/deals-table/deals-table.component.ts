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

  async postListItem(ingredient: Ingredient, index: number){
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
        const selector = `#row` + index;
        const element = document.querySelector(selector) as HTMLElement
        if(element){
          this.toggleInList(element)
        }
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
