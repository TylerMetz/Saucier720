import { TestBed, inject } from '@angular/core/testing';
import { HttpEvent, HttpEventType } from '@angular/common/http';
import {HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { DealsService } from '../../core/services/deals/deals.service';
import { DEALS } from 'src/app/mocks/deals.mock';

describe('DealsService', () => {
  
  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [ HttpClientTestingModule ],
      providers: [ DealsService ]
    });
  });

  it('Receives deals when loading page',
  inject(
    [HttpTestingController, DealsService],
    (httpMock: HttpTestingController, dealsService: DealsService) => {
      dealsService.getDeals().subscribe((event: HttpEvent<any>) => {
        switch (event.type) {
          case HttpEventType.Response:
            expect(event.body).equal(DEALS);
        }
      });

      const mockReq = httpMock.expectOne(dealsService.dealsUrl);
      expect(mockReq.cancelled).to.equal(false);
      expect(mockReq.request.responseType).to.equal('json');
      mockReq.flush(DEALS);

      httpMock.verify();

    }
  ))
});
