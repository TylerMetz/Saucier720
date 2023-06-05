import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-sub-recipe',
  templateUrl: './sub-recipe.component.html',
  styleUrls: ['./sub-recipe.component.scss']
})
export class SubRecipeComponent {
  @Input() recipe: string = '';
  @Input() ingredients: string[] = [];
  @Input() headerText: string = '';

  filteredIngredients(): string[] {
    const index = this.ingredients.findIndex(ingredient => ingredient === 'recipe follows');
    return this.ingredients.slice(0, index);
  }
}
