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

  // Output 
  //@Output() refreshDealsTable: EventEmitter<void> = new EventEmitter<void>();

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
  // async postStore(storeName: string) {
  //   const newStore: Store = {
  //     Name: storeName,
  //   };
  //   console.log(newStore)
  //   try {
  //     const response = await lastValueFrom(this.dealsService.postStore(newStore));
  //     console.log(response);

  //     this.refreshDealsTable.emit()
  //     const buttons = document.querySelectorAll('button');
  //     buttons.forEach((button: HTMLElement) => {
  //       if (button.innerText === storeName) {
  //         button.classList.add('clicked');
  //       } else {
  //         button.classList.remove('clicked');
  //       }
  //    });
      
  //   } catch (error) {
  //     console.error(error);
  //   }
  // }

  // setButton() {
  //   const buttons = document.querySelectorAll('button');
  //   buttons.forEach((button: HTMLElement) => {
  //     if (button.innerText === this.activeButton) {
  //       button.classList.add('clicked');
  //     } else {
  //       button.classList.remove('clicked');
  //     }
  //  });
  // }

}
