import { Component } from '@angular/core';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';
import { PANTRY } from 'src/app/mocks/pantry.mock';
import { lastValueFrom } from 'rxjs';

@Component({
  selector: 'app-new-pantry-item-button',
  templateUrl: './new-pantry-item-button.component.html',
  providers: [PantryService]
})

export class NewPantryItemButtonComponent {

  constructor(private pantryService: PantryService) { }

  async postPantryItem() {
    try {
      const response = await lastValueFrom(this.pantryService.postPantryItem(PANTRY[0]));
      console.log(response);
    } catch (error) {
      console.error(error);
    }
  }
}
