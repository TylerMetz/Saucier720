import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';

import { PantryComponent } from '../../pantry/pantry.component';


describe('PantryComponent', () => {
  it('mounts', () => {
    cy.mount(PantryComponent)
  })

});
