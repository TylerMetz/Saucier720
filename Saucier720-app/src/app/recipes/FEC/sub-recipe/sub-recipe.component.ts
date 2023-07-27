import { Component, Input } from '@angular/core';
import { ListComponent } from 'src/app/list/list.component';

@Component({
  selector: 'app-sub-recipe',
  templateUrl: './sub-recipe.component.html',
  styleUrls: ['./sub-recipe.component.scss'],
  providers: [ListComponent],
})
export class SubRecipeComponent {
  @Input() recipe: string = '';
  @Input() ingredients: string[] = [];
  @Input() headerText: string = '';

  constructor (private listComponent: ListComponent) {}

  filteredIngredients(): string[] {
    const index = this.ingredients.findIndex(ingredient => ingredient === 'recipe follows');
    return this.ingredients.slice(0, index);
  }

  addToList(ingredient: string) {
    this.listComponent.addIngredient(ingredient);
  }
}


