import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SubRecipeComponent } from './sub-recipe.component';

describe('SubRecipeComponent', () => {
  let component: SubRecipeComponent;
  let fixture: ComponentFixture<SubRecipeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SubRecipeComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SubRecipeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
