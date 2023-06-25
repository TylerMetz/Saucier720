import { Component, OnInit } from '@angular/core';
import { PANTRY } from 'src/app/mocks/pantry.mock';
import { lastValueFrom } from 'rxjs';
import { RecipeService } from 'src/app/core/services/recipes/recipe.service';

@Component({
  selector: 'app-recipe-nav-bar',
  templateUrl: './recipe-nav-bar.component.html',
  styleUrls: ['./recipe-nav-bar.component.scss'],
  providers: [RecipeService]
})
export class RecipeNavBarComponent implements OnInit {
  showMyRecipesButton = true;
  showFavoriteRecipesButton = true;
  showRecommendedRecipesButton = false;

  constructor(private recipeService: RecipeService) {}

  ngOnInit() {
    this.loadButtonState();
  }

  toggleButtons(buttonType: string) {
    if (buttonType === 'myRecipes') {
      this.showMyRecipesButton = false;
      this.showFavoriteRecipesButton = true;
      this.showRecommendedRecipesButton = true;
    } else if (buttonType === 'favoriteRecipes') {
      this.showMyRecipesButton = true;
      this.showFavoriteRecipesButton = false;
      this.showRecommendedRecipesButton = true;
    } else if (buttonType === 'recommendedRecipes') {
      this.showMyRecipesButton = true;
      this.showFavoriteRecipesButton = true;
      this.showRecommendedRecipesButton = false;
    }

    this.saveButtonState();
  }

  async postFavoriteRecipesSelect() {
    try {
      const response = await lastValueFrom(this.recipeService.postFavoriteRecipesSelect());
      console.log(response);
      window.location.reload();
    } catch (error) {
      console.error(error);
    }
  }

  async postMyRecipesSelect() {
    try {
      const response = await lastValueFrom(this.recipeService.postMyRecipesSelect());
      console.log(response);
      window.location.reload();
    } catch (error) {
      console.error(error);
    }
  }

  async postRecommendedRecipesSelect() {
    try {
      const response = await lastValueFrom(this.recipeService.postRecommendedRecipesSelect());
      console.log(response);
      window.location.reload();
    } catch (error) {
      console.error(error);
    }
  }

  loadButtonState() {
    const buttonState = localStorage.getItem('recipeNavBarButtonState');
    if (buttonState) {
      const state = JSON.parse(buttonState);
      this.showMyRecipesButton = state.showMyRecipesButton;
      this.showFavoriteRecipesButton = state.showFavoriteRecipesButton;
      this.showRecommendedRecipesButton = state.showRecommendedRecipesButton;
    }
  }

  saveButtonState() {
    const buttonState = {
      showMyRecipesButton: this.showMyRecipesButton,
      showFavoriteRecipesButton: this.showFavoriteRecipesButton,
      showRecommendedRecipesButton: this.showRecommendedRecipesButton
    };
    localStorage.setItem('recipeNavBarButtonState', JSON.stringify(buttonState));
  }
}
