import { PantryService } from 'src/app/core/services/pantry/pantry.service';
import { TestBed, inject } from '@angular/core/testing';
import { HttpEvent, HttpEventType } from '@angular/common/http';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { PANTRY } from 'src/app/mocks/pantry.mock';
import { PantryComponent } from '../../pantry/pantry.component';


describe('PantryComponent', () => {

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [ HttpClientTestingModule ],
      providers: [ PantryService ]
    });
    cy.mount(PantryComponent);
  });
  it ('displays table', () => {
    cy.get('app-pantry-table');
  });
  it ('displays post button', () =>{
    cy.get('app-new-pantry-item-button');
  });
  it('should get pantry',
  inject(
    [HttpTestingController, PantryService],
    (httpMock: HttpTestingController, pantryService: PantryService) => {
      pantryService.getPantry().subscribe((event: HttpEvent<any>) => {
        switch (event.type) {
          case HttpEventType.Response:
            expect(event.body).equal(PANTRY);
        }
      });

      const mockReq = httpMock.expectOne(pantryService.pantryUrl);

      expect(mockReq.cancelled).to.equal(false);
      expect(mockReq.request.responseType).to.equal('json');
      mockReq.flush(PANTRY);

      httpMock.verify();

    }
  ));


});
