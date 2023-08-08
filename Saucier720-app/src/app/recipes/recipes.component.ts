import { Component, ViewChild } from '@angular/core';
import { RecipeCardComponent } from './FEC/recipe-card/recipe-card.component';
// import { RecipeNavBarComponent } from 'src/app/recipes/recipe-nav-bar.component';

@Component({
  selector: 'app-recipes',
  templateUrl: './recipes.component.html',
  styleUrls: ['./recipes.component.scss']
})

export class RecipesComponent{
  public isNewRecipeEnabled: boolean = false;
  public isFilterMenuEnabled: boolean = false;

  @ViewChild('recipeCard') private recipeCard!: RecipeCardComponent;

  handleRecipeRefresh(){
    this.recipeCard.populateRecipes();
    this.recipeCard.currentRecipeIndex = 0;
  }
}
