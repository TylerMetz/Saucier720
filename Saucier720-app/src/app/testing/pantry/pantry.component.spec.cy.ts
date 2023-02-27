import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PantryComponent } from '../../pantry/pantry.component';

describe('PantryComponent', () => {
  it('mounts', () => {
    cy.mount(PantryComponent)
  })

  //TODO: example test for getting a html component
  // it('inits with', () => {
  //   cy.mount(PantryComponent)
  //   cy.get('div').should('have.class', 'page')
  // })

});
