import { Injectable } from '@angular/core';
import { Ingredient } from '../interfaces/ingredient';
import { PANTRY } from '../../mocks/ingredients.mock';

@Injectable({
  providedIn: 'root'
})
export class IngredientService {

  getPantry(): Array<Ingredient> {
    return PANTRY;
  }

}
