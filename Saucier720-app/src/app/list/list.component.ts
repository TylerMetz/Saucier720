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
  newIngredientName: string = '';
  newIngredientQuantity: number = 0;

  constructor(private ingredientService: IngredientService) { }

  ngOnInit() {
    this.ingredients = this.ingredientService.getPantry();
  }

  adjustQuantity(ingredient: Ingredient, action: string) {
    if (action === 'increment') {
      ingredient.Quantity += 1;
    } else if (action === 'decrement' && ingredient.Quantity > 0) {
      ingredient.Quantity -= 1;
    }
  }

  deleteIngredient(ingredient: Ingredient) {
    const index = this.ingredients.indexOf(ingredient);
    if (index > -1) {
      this.ingredients.splice(index, 1);
    }
  }

  addIngredient() {
    if (this.newIngredientName && this.newIngredientQuantity > 0) {
      const newIngredient: Ingredient = {
        Name: this.newIngredientName,
        Quantity: this.newIngredientQuantity,
        StoreCost: 0, // Example value, replace with actual value if needed
        OnSale: false, // Example value, replace with actual value if needed
        SalePrice: 0, // Example value, replace with actual value if needed
        SaleDetails: '' // Example value, replace with actual value if needed
      };
      this.ingredients.push(newIngredient);

      // Clear input fields
      this.newIngredientName = '';
      this.newIngredientQuantity = 0;
    }
  }
}
