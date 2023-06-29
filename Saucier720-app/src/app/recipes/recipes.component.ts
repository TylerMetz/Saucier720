import { Component } from '@angular/core';
// import { RecipeNavBarComponent } from 'src/app/recipes/recipe-nav-bar.component';

@Component({
  selector: 'app-recipes',
  templateUrl: './recipes.component.html',
  styleUrls: ['./recipes.component.scss']
})

export class RecipesComponent{
  public isNewRecipeEnabled: boolean = false;
}
