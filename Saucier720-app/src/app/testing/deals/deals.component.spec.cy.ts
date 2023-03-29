import { ComponentFixture, TestBed } from '@angular/core/testing';
import { DealsComponent } from '../../deals/deals.component';

describe('DealsComponent', () => {
  beforeEach(() => {
    cy.mount(DealsComponent);
  })
  it('displays table', () =>{
    cy.get('app-deals-table');
  });
});
