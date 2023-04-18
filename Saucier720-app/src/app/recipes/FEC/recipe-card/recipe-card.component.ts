import { Component, OnInit } from '@angular/core';
import { HttpEvent, HttpEventType } from '@angular/common/http';
import { RecipeService } from 'src/app/core/services/recipes/recipe.service';
import { lastValueFrom } from 'rxjs';
import { RecipePost } from 'src/app/core/interfaces/recipe';

@Component({
  selector: 'app-recipe-card',
  templateUrl: './recipe-card.component.html',
  styleUrls: ['./recipe-card.component.scss']
})
export class RecipeCardComponent implements OnInit {

  recipes: RecipePost[] = [];
  currentRecipeIndex: number = 0;
  currentRecipe: RecipePost = this.recipes[this.currentRecipeIndex];

  constructor(
    private recipeService: RecipeService
    ) {}

    ngOnInit(): void {
      this.populateRecipes();
    }

    public async populateRecipes(): Promise<void> {
      try {
        const event: HttpEvent<any> = await lastValueFrom(this.recipeService.getRecipes());
        switch(event.type) {
          case HttpEventType.Sent:
            console.log('Request sent!');
            break;
          case HttpEventType.ResponseHeader:
            console.log('Response header received!');
            break;
          case HttpEventType.DownloadProgress:
            const kbLoaded = Math.round(event.loaded / 1024);
            console.log(`Download in progress! ${kbLoaded}Kb loaded`);
            break;
          case HttpEventType.Response:
            console.log('Done!', event.body);
            let recipeStr = JSON.stringify(event.body);
            let parsedRecipes = JSON.parse(recipeStr);
            this.recipes = parsedRecipes;
            console.log('Parse', parsedRecipes)
            console.log('Recipes',this.recipes);
            break;
        }
      } catch (error) {
        console.error(error);
      }
    }

    goToNextRecipe() {
      this.currentRecipeIndex++;
      if (this.currentRecipeIndex >= this.recipes.length) {
        this.currentRecipeIndex = 0;
      }
    }

    goToPrevRecipe() {
      this.currentRecipeIndex++;
      if (this.currentRecipeIndex >= this.recipes.length) {
        this.currentRecipeIndex = 0;
      }
    }

    checkForRecipeFollows(ingredient: string): boolean {
      const pattern = /, recipe follows,$/;
      return pattern.test(ingredient);
    }
}
