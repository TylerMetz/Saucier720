import { TestBed } from '@angular/core/testing';

import { IngredientService } from '../../../core/services/ingredient.service';

describe('IngredientService', () => {
  let service: IngredientService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(IngredientService);
  });
});
