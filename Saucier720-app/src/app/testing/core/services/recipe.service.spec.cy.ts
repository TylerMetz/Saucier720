import { TestBed } from '@angular/core/testing';

import { RecipeService } from '../../../core/services/recipes/recipe.service';

describe('RecipeService', () => {
  let service: RecipeService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RecipeService);
  });
});
