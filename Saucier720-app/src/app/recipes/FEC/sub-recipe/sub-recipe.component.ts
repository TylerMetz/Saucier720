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

  sanitizeHtmlId(inputStr: string): string {
    // Remove all characters except letters (a-z or A-Z), digits (0-9), hyphens (-), and underscores (_)
    const sanitizedStr = inputStr.replace(/[^a-zA-Z0-9-_]/g, '');
    return "s-"+ sanitizedStr;
  }

  checkIngredient(rowId: string){
    //console.log(rowId)
    const element = document.querySelector(rowId) as HTMLElement
    if(element){
      //console.log("valid id: " + rowId)
      this.toggleInList(element)
    }
  }

  toggleInList(element: HTMLElement){
    element.classList.remove("not-in-list")
    element.classList.add("in-list");
    element.title = "Already in list!"
  }
}


