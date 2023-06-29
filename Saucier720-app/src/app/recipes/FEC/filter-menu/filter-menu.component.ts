import { Component, EventEmitter, Output, Renderer2, ViewChild, ElementRef } from '@angular/core';
import { RecipeService } from 'src/app/core/services/recipes/recipe.service';
import { Recipe } from 'src/app/core/interfaces/recipe';
import { lastValueFrom } from 'rxjs';

@Component({
  selector: 'app-filter-menu',
  templateUrl: './filter-menu.component.html',
  styleUrls: ['./filter-menu.component.scss'],
  providers: [RecipeService]
})
export class FilterRecipeComponent {


  constructor(private renderer: Renderer2, private recipeService: RecipeService) {}

  
}