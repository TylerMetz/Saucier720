import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NewPantryItemButtonComponent } from './new-pantry-item-button.component';

describe('NewPantryItemButtonComponent', () => {
  let component: NewPantryItemButtonComponent;
  let fixture: ComponentFixture<NewPantryItemButtonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ NewPantryItemButtonComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(NewPantryItemButtonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
