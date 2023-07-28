import { Component, OnInit, Input, ChangeDetectionStrategy, ChangeDetectorRef } from '@angular/core';
import { HttpEvent, HttpEventType } from '@angular/common/http';
import { RecipeService } from 'src/app/core/services/recipes/recipe.service';
import { lastValueFrom } from 'rxjs';
import { RecipePost } from 'src/app/core/interfaces/recipe';
import { CookieService } from 'ngx-cookie-service';
import { ListComponent } from 'src/app/list/list.component';
import { Ingredient } from 'src/app/core/interfaces/ingredient';

@Component({
  selector: 'app-recipe-card',
  templateUrl: './recipe-card.component.html',
  styleUrls: ['./recipe-card.component.scss'],
  providers: [ListComponent],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class RecipeCardComponent implements OnInit {

  recipes: RecipePost[] = [];
  currentRecipeIndex: number = 0;
  currentRecipe!: RecipePost;
  @Input() currentIngredients: string[] = [];
  nextRecipeFollows: any;
  printedSubRecipeLines: string[] = [];
  hasError: boolean = false;

  constructor(
    private recipeService: RecipeService,
    private cdRef: ChangeDetectorRef,
    private cookieService: CookieService,
    private listComponent: ListComponent,
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

          if (!this.recipes) {
            this.hasError = true; // Set the error flag
          } else {
            this.hasError = false; // Clear the error flag
            this.currentRecipe = this.recipes[this.currentRecipeIndex];
            this.currentIngredients = this.removeQuotesAndBrackets(this.currentRecipe.R.ingredients);
            console.log(this.currentIngredients);
            // Validate ingredient 
            for (var i = 0; i < this.currentIngredients.length; i++){
               this.checkInList(this.currentIngredients[i], "#row" + i)
            }
          }
          break;
      }
    } catch (error) {
      console.error(error);
    }
  }

  public getAuthorCreditFromRecipeID(recipeID: string): string {
    // used to get recipe author from recipeID
    const author = recipeID.replace(/\d+/g, '');
    if (author === 'json') {
      return 'MealDealz Classic Recipe';
    } else {
      return 'Created by ' + author;
    }
  }

  public isCurrentUserRecipe(recipeID: string): boolean {
    // used to check if current recipe is made my the current user
    if (recipeID.replace(/\d+/g, '') === this.cookieService.get("sessionID").replace(/\d+/g, '')){
      return true
    } else{
      return false;
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
    console.log('hi2', this.currentIngredients)
    console.log(this.currentRecipe.R.title)
  }

  goToPrevRecipe() {
    this.currentRecipeIndex--;
    if (this.currentRecipeIndex < 0) {
      this.currentRecipeIndex = this.recipes.length - 1;
    }
    this.currentRecipe = this.recipes[this.currentRecipeIndex];
    this.currentIngredients = this.removeQuotesAndBrackets(this.currentRecipe.R.ingredients);
    console.log('hi3', this.currentIngredients)
    console.log(this.currentRecipe.R.title)
  }

  checkForRecipeFollows(ingredient: string): boolean {
    const pattern = /\brecipe\s+follows\b/i;
    //console.log('recipe follows',pattern.test(ingredient))
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

  // for favorite button
  async toggleFavorite() {

    // switch favorite val
    this.recipes[this.currentRecipeIndex].R.userFavorite = !this.recipes[this.currentRecipeIndex].R.userFavorite;

    if(this.recipes[this.currentRecipeIndex].R.userFavorite) {
      try {
        const response = await lastValueFrom(this.recipeService.postFavoriteRecipe(this.currentRecipe.R.recipeID));
        console.log(response);
      } catch (error) {
        console.error(error);
      }
    }
    else if (!this.recipes[this.currentRecipeIndex].R.userFavorite){
      try {
        const response = await lastValueFrom(this.recipeService.postRemoveFavoriteRecipe(this.currentRecipe.R.recipeID));
        console.log(response);
      } catch (error) {
        console.error(error);
      }
    }
  }

  holdTimer: any;
  showHoldToConfirm: boolean = false;
  deleteIconOpacity: number = 0.8 // Add a variable to store the current opacity of the delete icon
  
  startHoldTimer() {
    this.deleteIconOpacity = 0.1;
    this.holdTimer = setInterval(() => { // Use setInterval instead of setTimeout to update the opacity continuously
      this.showHoldToConfirm = true;
      const holdDuration = 3000; // Set the hold duration in milliseconds (3 seconds in this example)
      const opacityStep = 0.9 / (holdDuration / 100); // Calculate the step to reach opacity 1 in 3 seconds
      this.deleteIconOpacity += opacityStep;
      if (this.deleteIconOpacity >= 1) {
        this.deleteIconOpacity = 1; // Ensure the opacity does not exceed 1
        clearInterval(this.holdTimer);
        this.deleteUserRecipe(); // Call the deleteUserRecipe() method after the hold duration
      }
    }, 100); // Run the interval every 100ms for smoother transition
  }
  
  clearHoldTimer() {
    clearInterval(this.holdTimer); // Use clearInterval to stop the interval from updating the opacity
    this.showHoldToConfirm = false;
    this.deleteIconOpacity = 0.8; // Reset the opacity to 1
  }
  
  endHoldTimer() {
    clearInterval(this.holdTimer); // Use clearInterval to stop the interval from updating the opacity
    this.showHoldToConfirm = false;
    this.deleteIconOpacity = 0.8; // Reset the opacity to 1
  }

  // for delete button
  async deleteUserRecipe() {
    try {
      // make post req
      const response = await lastValueFrom(this.recipeService.postDeleteUserRecipe(this.currentRecipe.R.recipeID));
      console.log(response);

      // delete this recipe card and move to the next
      this.recipes.splice(this.currentRecipeIndex, 1);

      // need to reload if there are no recipes left
      if(this.currentRecipeIndex === 0){
        window.location.reload();
      }
      else{
        this.goToNextRecipe();
      }
      
      
    } catch (error) {
      console.error(error);
    }
  }

  addToList(ingredient: string) {
    this.listComponent.addIngredient(ingredient);
  }

  async checkInList(ingredient: string, rowId: string) {
    // Create a temporary variable to easily fill into the check 
    let tempIngredient: Ingredient | null = null;
    if(ingredient){
      tempIngredient = {
        Name: ingredient, // Necessary for check
        Quantity: 1, // Necessary for check
        StoreCost: 0, // Filler
        OnSale: false, // Filler
        SalePrice: 0, // Filler
        SaleDetails: '' // Filler
      }
      const isValid = await this.listComponent.validateIngredient(tempIngredient)
      if (isValid){
        const element = document.querySelector(rowId) as HTMLElement
        if(element){
          this.toggleInList(element)
        }
      }
      
    }
  }

  toggleInList(element: HTMLElement){
    element.classList.remove("not-in-list")
    element.classList.add("in-list");
    element.title = "Already in list!"
  }



}
