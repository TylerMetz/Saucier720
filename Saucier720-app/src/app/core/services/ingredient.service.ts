import { Injectable } from '@angular/core';
import { Ingredient } from '../interfaces/ingredient';
import { FOODITEM } from 'src/app/mocks/ingredients.mock';

@Injectable({
  providedIn: 'root'
})
export class IngredientService {

  getPantry(): Array<Ingredient> {
    return FOODITEM;
  }

}
