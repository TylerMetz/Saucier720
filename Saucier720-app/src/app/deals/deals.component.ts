import { Component, ViewChild } from '@angular/core';
import { DealsTableComponent } from './FEC/deals-table/deals-table.component';
import { DealsStoreButtonComponent } from './FEC/deals-store-button/deals-store-button.component';

@Component({
  selector: 'app-deals',
  templateUrl: './deals.component.html',
  styleUrls: ['./deals.component.scss']
})
export class DealsComponent {

  @ViewChild('dealsTable') private dealsTable!: DealsTableComponent;
  handleDealsRefresh(store: string){
    console.log('hi')
    this.dealsTable.populateDeals(store);
  }

  @ViewChild('storeButton') private storeButton!: DealsStoreButtonComponent;
  handleSendButton(name: string){
    this.storeButton.activeButton = name
    this.storeButton.setButton()
  }


}