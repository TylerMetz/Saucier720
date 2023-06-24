import { Component } from '@angular/core';
import { PANTRY } from 'src/app/mocks/pantry.mock';
import { lastValueFrom } from 'rxjs';


@Component({
  selector: 'app-recipe-nav-bar',
  templateUrl: './recipe-nav-bar.component.html',
  styleUrls:[ './recipe-nav-bar.component.scss'],
  providers: []
})

export class RecipeNavBarComponent {
  constructor() { }


}
