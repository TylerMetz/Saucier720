import { TestBed } from '@angular/core/testing';

import { PantryService } from '../../core/services/pantry/pantry.service';

describe('PantryService', () => {
  let service: PantryService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PantryService);
  });
});
