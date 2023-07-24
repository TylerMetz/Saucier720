import { Component } from '@angular/core';
import { DealsService } from 'src/app/core/services/deals/deals.service';
import { PANTRY } from 'src/app/mocks/pantry.mock';
import { lastValueFrom } from 'rxjs';
import { Store } from 'src/app/core/interfaces/store';


@Component({
  selector: 'app-deals-store-button',
  templateUrl: './deals-store-button.component.html',
  styleUrls:[ './deals-store-button.component.scss'],
  providers: [DealsService]
})

export class DealsStoreButtonComponent {
  name: string = '';

  constructor(private dealsService: DealsService) {
    const buttonStateJson = localStorage.getItem('buttonState');
    if (buttonStateJson) {
      const buttonState = JSON.parse(buttonStateJson);
      if (buttonState && buttonState.storeName) {
        this.name = buttonState.storeName;
      }
    }
  }

  async postStore(storeName: string) {
    const newStore: Store = {
      Name: storeName,
    };
    console.log(newStore)
    try {
      const response = await lastValueFrom(this.dealsService.postStore(newStore));
      console.log(response);
      this.saveButtonState(storeName);

      window.location.reload();
    } catch (error) {
      console.error(error);
    }
  }

  private saveButtonState(storeName: string): void {
    localStorage.setItem('buttonState', JSON.stringify({ storeName }));
  }

  ngAfterViewInit() {
    const buttons = document.querySelectorAll('button');
    buttons.forEach((button: HTMLElement) => {
      if (button.innerText === this.name) {
        button.classList.add('clicked');
      }
    });
  }
}
