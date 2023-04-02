import { ComponentFixture, TestBed, inject } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { HttpEvent, HttpEventType } from '@angular/common/http';

import { PANTRY } from 'src/app/mocks/pantry.mock';
import { NewPantryItemButtonComponent } from 'src/app/pantry/FEC/new-pantry-item-button/new-pantry-item-button.component';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';

describe('NewPantryItemButtonComponent', () => {
  let component: NewPantryItemButtonComponent;
  let fixture: ComponentFixture<NewPantryItemButtonComponent>;
  let pantryService: PantryService;
  let httpMock: HttpTestingController;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [NewPantryItemButtonComponent],
      imports: [HttpClientTestingModule],
      providers: [PantryService]
    })
      .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NewPantryItemButtonComponent);
    component = fixture.componentInstance;
    pantryService = TestBed.inject(PantryService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  it('should post pantry item',
    inject(
      [HttpTestingController, PantryService],
      (httpMock: HttpTestingController, pantryService: PantryService) => {
        component.postPantryItem();

        const mockReq = httpMock.expectOne(component.pantryPostUrl);

        expect(mockReq.cancelled).to.be.false;
        expect(mockReq.request.responseType).equal('json');
        mockReq.flush({});

        httpMock.verify();

      }
  ));

  it('should call postPantryItem method on button click', () => {
    const button = fixture.nativeElement.querySelector('button');
    cy.spy(component, 'postPantryItem');
    button.click();
    expect(component.postPantryItem).to.have.been.called;
  });
});
