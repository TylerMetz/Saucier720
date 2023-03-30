import { DealsService } from 'src/app/core/services/deals/deals.service';
import { TestBed, inject } from '@angular/core/testing';
import { HttpEvent, HttpEventType } from '@angular/common/http';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { DealsComponent } from '../../deals/deals.component';

describe('DealsComponent', () => {
  beforeEach(() => {
    cy.mount(DealsComponent);
  })
  it('displays table', () =>{
    cy.get('app-deals-table');
  });
});
