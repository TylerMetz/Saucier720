import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PantryTableComponent } from './pantry-table.component';

describe('PantryTableComponent', () => {
  let component: PantryTableComponent;
  let fixture: ComponentFixture<PantryTableComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PantryTableComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PantryTableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
