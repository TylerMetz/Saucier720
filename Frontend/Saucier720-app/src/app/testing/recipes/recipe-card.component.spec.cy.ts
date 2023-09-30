import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RecipeCardComponent } from '../../recipes/FEC/recipe-card/recipe-card.component';

describe('RecipeCardComponent', () => {
  let component: RecipeCardComponent;
  let fixture: ComponentFixture<RecipeCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RecipeCardComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RecipeCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });
});
