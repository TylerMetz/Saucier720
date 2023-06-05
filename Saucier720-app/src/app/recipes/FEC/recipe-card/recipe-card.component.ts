import { Component, OnInit, Input, ChangeDetectionStrategy, ChangeDetectorRef } from '@angular/core';
import { HttpEvent, HttpEventType } from '@angular/common/http';
import { RecipeService } from 'src/app/core/services/recipes/recipe.service';
import { lastValueFrom } from 'rxjs';
import { RecipePost } from 'src/app/core/interfaces/recipe';

@Component({
  selector: 'app-recipe-card',
  templateUrl: './recipe-card.component.html',
  styleUrls: ['./recipe-card.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class RecipeCardComponent implements OnInit {

  recipes: RecipePost[] = [];
  currentRecipeIndex: number = 0;
  currentRecipe!: RecipePost;
  @Input() currentIngredients: string[] = [];
  nextRecipeFollows: any;
  printedSubRecipeLines: string[] = [];

  constructor(
    private recipeService: RecipeService,
    private cdRef: ChangeDetectorRef
    ) {}

    ngOnInit(): void {
      this.populateRecipes();
      this.updatePrintedSubRecipeLines();
    }

    ngDoCheck(): void {
      this.cdRef.detectChanges();
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
            this.currentRecipe = this.recipes[this.currentRecipeIndex];
            this.currentIngredients = this.removeQuotesAndBrackets(this.currentRecipe.R.ingredients);
            console.log(this.currentIngredients)
            break;
        }
      } catch (error) {
        console.error(error);
      }
    }

    private updatePrintedSubRecipeLines() {
      this.currentIngredients.forEach((ingredient, i) => {
        if (this.checkForRecipeFollows(ingredient)) {
          const subIngredients = this.currentIngredients.slice(i + 1);
          this.printedSubRecipeLines.push(ingredient);
          subIngredients.forEach(subIngredient => {
            if (!this.checkForRecipeFollows(subIngredient)) {
              this.printedSubRecipeLines.push(subIngredient);
            }
          });
        }
      });
    }

    goToNextRecipe() {
      this.currentRecipeIndex++;
      if (this.currentRecipeIndex >= this.recipes.length) {
        this.currentRecipeIndex = 0;
      }
      this.currentRecipe = this.recipes[this.currentRecipeIndex];
      this.currentIngredients = this.removeQuotesAndBrackets(this.currentRecipe.R.ingredients);
      console.log(this.currentIngredients)
      console.log(this.currentRecipe.R.title)
    }

    goToPrevRecipe() {
      this.currentRecipeIndex--;
      if (this.currentRecipeIndex < 0) {
        this.currentRecipeIndex = this.recipes.length - 1;
      }
      this.currentRecipe = this.recipes[this.currentRecipeIndex];
      this.currentIngredients = this.removeQuotesAndBrackets(this.currentRecipe.R.ingredients);
      console.log(this.currentIngredients)
      console.log(this.currentRecipe.R.title)
    }

    checkForRecipeFollows(ingredient: string): boolean {
      const pattern = /\brecipe\s+follows\b/i;
      console.log('recipe follows',pattern.test(ingredient))
      return pattern.test(ingredient);
    }
    

    getSubIngredients(): string[][] {
  const subIngredients: string[][] = [];
  let currentSubIngredients: string[] = [];

  for (let i = 0; i < this.currentIngredients.length; i++) {
    if (this.checkForRecipeFollows(this.currentIngredients[i])) {
      // If we've encountered a "recipe follows" line, add the current sub-ingredients
      // to the sub-ingredients array and start a new one.
      if (currentSubIngredients.length > 0) {
        subIngredients.push(currentSubIngredients);
        currentSubIngredients = [];
      }
    } else {
      // If the current ingredient is not a "recipe follows" line, add it to the current
      // sub-ingredients array.
      currentSubIngredients.push(this.currentIngredients[i]);
    }
  }

  // Add the final sub-ingredients array to the sub-ingredients array.
  subIngredients.push(currentSubIngredients);

  return subIngredients;
}


    removeQuotesAndBrackets(arr: string[]): string[] {
      const regex = /["\[\]]/g; // Matches any occurrence of ", [, or ] globally
      return arr.map(str => str.replace(regex, '')); // Replace all matches in each string in the array
    }
    
}
