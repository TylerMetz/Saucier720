import { Component, OnInit } from '@angular/core';
import { Ingredient } from '../core/interfaces/ingredient';
import { IngredientService } from '../core/services/ingredient.service';
@Component({
  selector: 'app-recipes',
  templateUrl: './recipes.component.html',
  styleUrls: ['./recipes.component.scss']
})
export class RecipesComponent implements OnInit{
  ingredients: Array<Ingredient> = [];

  constructor(private ingredientService: IngredientService) { }

  ngOnInit(){
    this.ingredients = this.ingredientService.getPantry();
  }
}
