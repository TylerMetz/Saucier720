import { Component, OnInit } from '@angular/core';
import { Ingredient } from '../core/interfaces/ingredient';
import { IngredientService } from '../core/services/ingredient.service';
@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements OnInit {
  ingredients: Array<Ingredient> = [];

  constructor(private ingredientService: IngredientService) { }

  ngOnInit(){
    this.ingredients = this.ingredientService.getPantry();
  }
}
