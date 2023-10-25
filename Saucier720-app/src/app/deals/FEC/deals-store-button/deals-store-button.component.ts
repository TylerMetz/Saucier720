import { Component, EventEmitter, Output } from '@angular/core';
import { DealsService } from 'src/app/core/services/deals/deals.service';
import { PANTRY } from 'src/app/mocks/pantry.mock';
import { lastValueFrom } from 'rxjs';
import { Store } from 'src/app/core/interfaces/store';
import { OutletContext } from '@angular/router';


@Component({
  selector: 'app-deals-store-button',
  templateUrl: './deals-store-button.component.html',
  styleUrls:[ './deals-store-button.component.scss'],
  providers: [DealsService]
})

export class DealsStoreButtonComponent {
  lastClicked: string = '';

  @Output() storeClickEvent: EventEmitter<string> = new EventEmitter<string>();

  constructor(private dealsService: DealsService) {}

  sendStore(store: string){
    if (store !== this.lastClicked) { 
      this.lastClicked = store;
      this.storeClickEvent.emit(store);
    }
  }

  isButtonDisabled(store: string): boolean {
    return store === this.lastClicked;
  }
}
