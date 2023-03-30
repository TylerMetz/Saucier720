import { DealsService } from 'src/app/core/services/deals/deals.service';
import { TestBed, inject } from '@angular/core/testing';
import { HttpEvent, HttpEventType } from '@angular/common/http';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { DealsComponent } from '../../deals/deals.component';
import { DEALS } from 'src/app/mocks/deals.mock';

describe('DealsComponent', () => {

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [ HttpClientTestingModule ],
      providers: [ DealsService ]
    });
    cy.mount(DealsComponent);
  })
  it('displays table', () =>{
    cy.get('app-deals-table');
  });

  it('should get pantry',
  inject(
    [HttpTestingController, DealsService],
    (httpMock: HttpTestingController, dealsService: DealsService) => {
      dealsService.getDeals().subscribe((event: HttpEvent<any>) => {
        switch (event.type) {
          case HttpEventType.Response:
            expect(event.body).equal(PANTRY);
        }
      });

      const mockReq = httpMock.expectOne(dealsService.dealsUrl);

      expect(mockReq.cancelled).to.equal(false);
      expect(mockReq.request.responseType).to.equal('json');
      mockReq.flush(DEALS);

      httpMock.verify();

    }
  ));
});
