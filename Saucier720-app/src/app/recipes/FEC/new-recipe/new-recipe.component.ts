import { Component, EventEmitter, Output, Renderer2, ViewChild, ElementRef } from '@angular/core';
import { RecipeService } from 'src/app/core/services/recipes/recipe.service';
import { Recipe } from 'src/app/core/interfaces/recipe';
import { lastValueFrom } from 'rxjs';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-new-recipe',
  templateUrl: './new-recipe.component.html',
  styleUrls: ['./new-recipe.component.scss'],
  providers: [RecipeService]
})
export class NewRecipeComponent {
  @ViewChild('ingredientTextboxes', { static: true }) ingredientTextboxesRef!: ElementRef;
  @ViewChild('instructionTextboxes', { static: true }) instructionTextboxesRef!: ElementRef;
  @ViewChild('titleTextbox', { static: true }) titleTextbox!: ElementRef;

  @Output() addRecipeToCookbook: EventEmitter<any> = new EventEmitter<any>();

  constructor(private renderer: Renderer2, private recipeService: RecipeService, private cookieService: CookieService) {}

  // No Refresh
  @Output() refreshRecipeCard: EventEmitter<void> = new EventEmitter<void>(); 

  constructor(private renderer: Renderer2, private recipeService: RecipeService) {}

  addIngredientTextbox() {
    const ingredientTextboxes = this.ingredientTextboxesRef.nativeElement;
    const newIngredientTextbox = this.renderer.createElement('input');
    this.renderer.setAttribute(newIngredientTextbox, 'type', 'text');
    this.renderer.addClass(newIngredientTextbox, 'ingredient-textbox');
    this.renderer.setAttribute(newIngredientTextbox, 'placeholder', `Ingredient ${ingredientTextboxes.children.length + 1}`);
    this.renderer.appendChild(ingredientTextboxes, newIngredientTextbox);
  }

  removeIngredientTextbox() {
    const ingredientTextboxes = this.ingredientTextboxesRef.nativeElement;
    const lastIngredientTextbox = ingredientTextboxes.lastElementChild;
    if (lastIngredientTextbox) {
      this.renderer.removeChild(ingredientTextboxes, lastIngredientTextbox);
    }
  }

  addInstructionTextbox() {
    const instructionTextboxes = this.instructionTextboxesRef.nativeElement;
    const newInstructionTextbox = this.renderer.createElement('input');
    this.renderer.setAttribute(newInstructionTextbox, 'type', 'text');
    this.renderer.addClass(newInstructionTextbox, 'instruction-textbox');
    this.renderer.setAttribute(newInstructionTextbox, 'placeholder', `Step ${instructionTextboxes.children.length + 1}`);
    this.renderer.appendChild(instructionTextboxes, newInstructionTextbox);
  }

  removeInstructionTextbox() {
    const instructionTextboxes = this.instructionTextboxesRef.nativeElement;
    const lastInstructionTextbox = instructionTextboxes.lastElementChild;
    if (lastInstructionTextbox) {
      this.renderer.removeChild(instructionTextboxes, lastInstructionTextbox);
    }
  }

  async addToCookbook() {
    if (this.ingredientTextboxesRef && this.instructionTextboxesRef) {
      const ingredientTextboxes = Array.from(this.ingredientTextboxesRef.nativeElement.querySelectorAll('input'));
      const instructionTextboxes = Array.from(this.instructionTextboxesRef.nativeElement.querySelectorAll('input'));

      const ingredients = ingredientTextboxes.map((textbox: unknown) => (textbox as HTMLInputElement).value.trim()).filter(Boolean);
      const instructions = instructionTextboxes.map((textbox: unknown) => (textbox as HTMLInputElement).value.trim()).filter(Boolean);

      const recipe: Recipe = {
        instructions: instructions.join('. '),
        ingredients,
        title: this.titleTextbox.nativeElement.value.trim(),
        pictureLink: null, // Provide the appropriate picture link here
        recipeID: 'test', // Provide the appropriate recipe ID here
        userFavorite: false,
        recipeAuthor: this.cookieService.get("sessionID").replace(/\d+/g, '') // pass in the current user's username, but doesn't matter because we store in backend anyways, just adding this incase we change it in backend
      };

      try {
        const response = await lastValueFrom(this.recipeService.postNewRecipe(recipe));
        console.log(response);
        this.refreshRecipeCard.emit()
      } catch (error) {
        console.error(error);
      }

    }
  }
}
