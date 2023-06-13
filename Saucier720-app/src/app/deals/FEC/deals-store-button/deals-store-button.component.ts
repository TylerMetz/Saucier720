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

  constructor(private dealsService: DealsService) { }

  async postStore() {
    const newStore: Store = {
      Name: this.name,
    };
    console.log(newStore)
    try {
      const response = await lastValueFrom(this.dealsService.postStore(newStore));
      console.log(response);
    } catch (error) {
      console.error(error);
    }
  }
}
