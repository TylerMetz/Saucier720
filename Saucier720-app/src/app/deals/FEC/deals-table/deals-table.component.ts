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

@Component({
  selector: 'app-deals-table',
  providers: [DealsService,ListComponent],
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

  constructor(private dealsService: DealsService, private listComponent: ListComponent) { }

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

  // Add to shopping list 
  addToList(ingredient: Ingredient, event: Event) {
    // Change button state
    const addBtn = event.target as HTMLElement;
    this.toggleInList(addBtn)
    // Nav to actual list function 
    //this.listComponent.addIngredient(ingredient);
  }

  toggleInList(element: HTMLElement){
    element.classList.remove("not-in-list")
    element.classList.add("in-list");
    element.title = "Already in list!"
  }

}
