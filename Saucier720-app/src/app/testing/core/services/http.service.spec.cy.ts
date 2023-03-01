import { TestBed } from '@angular/core/testing';

import { HttpService } from '../../../core/services/http.service';

describe('HttpService', () => {
  let service: HttpService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(HttpService);
  });
});
