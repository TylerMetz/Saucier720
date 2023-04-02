import { Component } from '@angular/core';
import { HttpClient, HttpEvent } from '@angular/common/http';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';
import { PANTRY } from 'src/app/mocks/pantry.mock';

@Component({
  selector: 'app-new-pantry-item-button',
  templateUrl: './new-pantry-item-button.component.html',
  providers: [PantryService]
})

export class NewPantryItemButtonComponent {
  pantryPostUrl = 'http://localhost:8082/api/NewPantryItem';

  constructor(private pantryService: PantryService) { }

  async postPantryItem() {
    try {
      const response = await this.pantryService.postPantryItem(PANTRY[0]).toPromise();
      console.log(response);
    } catch (error) {
      console.error(error);
    }
  }
}
