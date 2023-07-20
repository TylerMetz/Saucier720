import { Component, EventEmitter, Output, Renderer2, ViewChild, ElementRef, OnInit } from '@angular/core';
import { RecipeService } from 'src/app/core/services/recipes/recipe.service';
import { Recipe } from 'src/app/core/interfaces/recipe';
import { lastValueFrom } from 'rxjs';

@Component({
  selector: 'app-filter-menu',
  templateUrl: './filter-menu.component.html',
  styleUrls: ['./filter-menu.component.scss'],
  providers: [RecipeService]
})
export class FilterRecipeComponent implements OnInit{

  // set all checkboxes to false by default
  public myRecipesValue: boolean = false;
  public userRecipesValue: boolean = false;
  public mdRecipesValue: boolean = true; // true because recommended recipes start with mealdealz classic recipes

  constructor(private recipeService: RecipeService) {}

  ngOnInit() {
    // Load the checkbox values from storage when the component initializes
    this.loadCheckboxValues();
  }

  // function to load the checkbox values from local storage
  loadCheckboxValues() {
    const storedMyRecipesValue = localStorage.getItem('myRecipesValue');
    const storedUserRecipesValue = localStorage.getItem('userRecipesValue');
    const storedMdRecipesValue = localStorage.getItem('mdRecipesValue');

    if (storedMyRecipesValue !== null) {
      this.myRecipesValue = JSON.parse(storedMyRecipesValue);
    }
    if (storedUserRecipesValue !== null) {
      this.userRecipesValue = JSON.parse(storedUserRecipesValue);
    }
    if (storedMdRecipesValue !== null) {
      this.mdRecipesValue = JSON.parse(storedMdRecipesValue);
    }
  }


  // function to post filter checkbox values to backend
  async postFilterValues(){
    
    // save checkbox values to local storage in the case of a page refresh
    localStorage.setItem('myRecipesValue', JSON.stringify(this.myRecipesValue));
    localStorage.setItem('userRecipesValue', JSON.stringify(this.userRecipesValue));
    localStorage.setItem('mdRecipesValue', JSON.stringify(this.mdRecipesValue));

    try {
      // create object to hold filter checkbox values
      const filterValues = {
        myRecipesCheckbox: this.myRecipesValue,
        userRecipesCheckbox: this.userRecipesValue,
        mdRecipescheckbox: this.mdRecipesValue
      }
      const response = await lastValueFrom(this.recipeService.postFilterValues(filterValues));
      console.log(response);
      window.location.reload(); // refresh page to show updated data
    } catch (error) {
      console.error(error);
    }
  }
}