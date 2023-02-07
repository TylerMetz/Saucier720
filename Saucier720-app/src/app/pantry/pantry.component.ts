import { Component, OnInit } from '@angular/core';
import { Ingredient } from '../core/interfaces/ingredient';
import { IngredientService } from '../core/services/ingredient.service';

@Component({
  selector: 'app-pantry',
  templateUrl: './pantry.component.html',
  styleUrls: ['./pantry.component.scss']
})
export class PantryComponent implements OnInit {
  ingredients: Array<Ingredient> = [];

  constructor(private ingredientService: IngredientService) { }

  ngOnInit(){
    this.ingredients = this.ingredientService.getPantry();
  }



}
