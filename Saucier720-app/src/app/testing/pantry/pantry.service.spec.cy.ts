import { TestBed, inject } from '@angular/core/testing';
import { HttpEvent, HttpEventType } from '@angular/common/http';
import {HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';

import { PantryService } from '../../core/services/pantry/pantry.service';

describe('PantryService', () => {

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [ HttpClientTestingModule ],
      providers: [ PantryService ]
    });
  });
  it('should get pantry',
  inject(
    [HttpTestingController, PantryService],
    (httpMock: HttpTestingController, pantryService: PantryService) => {
      const mockPantry = [
        {"Name":"peanut butter","StoreCost":369.99,"OnSale":true,"SalePrice":0,"SaleDetails":"BOGO","Quantity":10},
        {"Name":"jelly","StoreCost":1,"OnSale":false,"SalePrice":0,"SaleDetails":"N/A","Quantity":30},
        {"Name":"bread","StoreCost":10.69,"OnSale":true,"SalePrice":0,"SaleDetails":"$2 for 2","Quantity":2}
      ];

      pantryService.getPantry().subscribe((event: HttpEvent<any>) => {
        switch (event.type) {
          case HttpEventType.Response:
            expect(event.body).equal(mockPantry);
        }
      });

      const mockReq = httpMock.expectOne(pantryService.pantryUrl);

      expect(mockReq.cancelled).toBe;

    }
  ))
});
