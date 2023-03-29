import { TestBed } from '@angular/core/testing';

import { DealsService } from '../../core/services/deals/deals.service';

describe('DealsService', () => {
  let service: DealsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(DealsService);
  });
});
