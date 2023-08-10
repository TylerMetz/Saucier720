import { HttpEvent, HttpEventType } from "@angular/common/http"
import { Component, OnInit, Input } from '@angular/core';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';
import { lastValueFrom } from 'rxjs';
import { Ingredient } from "src/app/core/interfaces/ingredient";

@Component({
  selector: 'app-pantry-table',
  providers: [PantryService],
  templateUrl: './pantry-table.component.html',
  styleUrls: ['./pantry-table.component.scss']
})
export class PantryTableComponent implements OnInit {

  pantry: Ingredient[] = [];

  constructor(private pantryService: PantryService) { }

  async ngOnInit(){
    await this.populatePantry();
  }

  public async populatePantry(): Promise<void> {
    try {
      const event: HttpEvent<any> = await lastValueFrom(this.pantryService.getPantry());
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
          this.pantry = event.body;
          break;
      }
    } catch (error) {
      console.error(error);
    }
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
      StoreCost: 0, 
      OnSale: false,
      SalePrice: 0, 
      SaleDetails: '', 
      Quantity: quantity
    };
    if(this.pantry === null){
      this.pantry = []
    }
    this.pantry.push(newIngredient)
    console.log('hi')
  }

  

}
