import { HttpEvent, HttpEventType } from "@angular/common/http"
import { Component, OnInit } from '@angular/core';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';

@Component({
  selector: 'app-pantry-table',
  providers: [PantryService],
  template: `
    <table class="table table-striped table-bordered">
      <thead>
        <tr>
          <th>Ingredients</th>
          <th>Cost</th>
          <th>On Sale</th>
          <th>Sale Price</th>
          <th>Sale Info</th>
          <th>Quantity</th>
        </tr>
      </thead>
      <tbody>
        <tr *ngFor="let ingredient of pantry">
          <td>{{ ingredient.Name }}</td>
          <td>{{ ingredient.StoreCost }}</td>
          <td>{{ ingredient.OnSale }}</td>
          <td>{{ ingredient.SalePrice }}</td>
          <td>{{ ingredient.SaleDetails }}</td>
          <td>{{ ingredient.Quantity }}</td>
        </tr>
      </tbody>
    </table>
  `
})
export class PantryTableComponent implements OnInit {

  pantry: any;

  constructor(private pantryService: PantryService) { }

  ngOnInit(){
    this.populatePantry();
  }
  populatePantry(): void {
    this.pantryService.getPantry()
      .subscribe((event: HttpEvent<any>) => {
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
        }
      }); //doesnt call until subscribed
  }

}
