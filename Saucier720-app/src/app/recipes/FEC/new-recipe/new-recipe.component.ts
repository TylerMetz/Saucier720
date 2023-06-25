import { Component, EventEmitter, Output, Renderer2, ViewChild, ElementRef } from '@angular/core';
import { RecipeService } from 'src/app/core/services/recipes/recipe.service';

@Component({
  selector: 'app-new-recipe',
  templateUrl: './new-recipe.component.html',
  styleUrls: ['./new-recipe.component.scss'],
  providers: [RecipeService]
})
export class NewRecipeComponent {
  @ViewChild('ingredientTextboxes', { static: true }) ingredientTextboxesRef!: ElementRef;
  @ViewChild('instructionTextboxes', { static: true }) instructionTextboxesRef!: ElementRef;
  
  @Output() addRecipeToCookbook: EventEmitter<any> = new EventEmitter<any>();

  constructor(private renderer: Renderer2, private recipeService: RecipeService) {}

  addIngredientTextbox() {
    const ingredientTextboxes = this.ingredientTextboxesRef.nativeElement;
    const newIngredientTextbox = this.renderer.createElement('input');
    this.renderer.setAttribute(newIngredientTextbox, 'type', 'text');
    this.renderer.addClass(newIngredientTextbox, 'ingredient-textbox');
    this.renderer.setAttribute(newIngredientTextbox, 'placeholder', `Ingredient ${ingredientTextboxes.children.length + 1}`);
    this.renderer.appendChild(ingredientTextboxes, newIngredientTextbox);
  }

  addInstructionTextbox() {
    const instructionTextboxes = this.instructionTextboxesRef.nativeElement;
    const newInstructionTextbox = this.renderer.createElement('input');
    this.renderer.setAttribute(newInstructionTextbox, 'type', 'text');
    this.renderer.addClass(newInstructionTextbox, 'instruction-textbox');
    this.renderer.setAttribute(newInstructionTextbox, 'placeholder', `Step ${instructionTextboxes.children.length + 1}`);
    const newListItem = this.renderer.createElement('li');
    this.renderer.appendChild(newListItem, newInstructionTextbox);
    this.renderer.appendChild(instructionTextboxes, newListItem);
  }

  addToCookbook() {
    // Your logic to save the recipe to the cookbook goes here
    // Example: this.recipeService.saveRecipe(recipe);
    // Emit an event to notify the parent component about the addition to the cookbook
    this.addRecipeToCookbook.emit();
    window.location.reload()
  }
}
