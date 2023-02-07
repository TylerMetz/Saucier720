import { Component, OnInit } from '@angular/core';
import { Ingredient } from '../core/interfaces/ingredient';
import { IngredientService } from '../core/services/ingredient.service';
@Component({
  selector: 'app-deals',
  templateUrl: './deals.component.html',
  styleUrls: ['./deals.component.scss']
})
export class DealsComponent implements OnInit{
  ingredients: Array<Ingredient> = [];

  constructor(private ingredientService: IngredientService) { }

  ngOnInit(){
    this.ingredients = this.ingredientService.getPantry();
  }
}
