import { TestBed, inject } from '@angular/core/testing';
import { HttpEvent, HttpEventType } from '@angular/common/http';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { PantryService } from '../../core/services/pantry/pantry.service';
import { PANTRY } from 'src/app/mocks/pantry.mock';

describe('PantryService', () => {
  
  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [ HttpClientTestingModule ],
      providers: [ PantryService ]
    });
  });

  it('Receives pantry when loading page',
  inject(
    [HttpTestingController, PantryService],
    (httpMock: HttpTestingController, pantryService: PantryService) => {
      pantryService.getPantry().subscribe((event: HttpEvent<any>) => {
        switch (event.type) {
          case HttpEventType.Response:
            expect(event.body).equal(PANTRY);
        }
      });

      // when navigating to the Pantry Page you make a request to
      // localhost:8080/api/Pantry that loads a User's Pantry
      const mockReq = httpMock.expectOne(pantryService.pantryUrl);
      expect(mockReq.cancelled).to.equal(false);
      expect(mockReq.request.responseType).to.equal('json');
      mockReq.flush(PANTRY);

      httpMock.verify();

    }
  ))
  
  it('Receives riley butter when clicking post button',
  inject(
    [HttpTestingController, PantryService],
    (httpMock: HttpTestingController, pantryService: PantryService) => {
      pantryService.getPantry().subscribe((event: HttpEvent<any>) => {
        switch (event.type) {
          case HttpEventType.Response:
            expect(event.body).equal(PANTRY[1]);
        }
      });

      // when navigating to the Pantry Page you make a request to
      // localhost:8080/api/Pantry that loads a User's Pantry
      const mockReq = httpMock.expectOne(pantryService.pantryUrl);
      expect(mockReq.cancelled).to.equal(false);
      expect(mockReq.request.responseType).to.equal('json');
      mockReq.flush(PANTRY[1]);

      httpMock.verify();

    }
  ))
});
