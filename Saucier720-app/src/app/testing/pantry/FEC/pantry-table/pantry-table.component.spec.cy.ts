import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';

import { PANTRY } from 'src/app/mocks/pantry.mock';
import { PantryTableComponent } from '../../../../pantry/FEC/pantry-table/pantry-table.component';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';

describe('PantryTableComponent', () => {
  let component: PantryTableComponent;
  let fixture: ComponentFixture<PantryTableComponent>;
  let pantryService: PantryService;
  let httpMock: HttpTestingController;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [PantryTableComponent],
      imports: [HttpClientTestingModule],
      providers: [PantryService]
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PantryTableComponent);
    component = fixture.componentInstance;
    pantryService = TestBed.inject(PantryService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  it('should render table with pantry', () => {
    component.pantry = PANTRY;
    fixture.detectChanges();

    const tableRows = fixture.nativeElement.querySelectorAll('.table-responsive tbody tr');
    expect(tableRows.length).equal(PANTRY.length);

    const headerRow = fixture.nativeElement.querySelectorAll('.table-responsive thead tr th');
    expect(headerRow.length).equal(6);
    expect(headerRow[0].textContent).equal('Ingredients');
    expect(headerRow[1].textContent).equal('Cost');
    expect(headerRow[2].textContent).equal('On Sale');
    expect(headerRow[3].textContent).equal('Sale Price');
    expect(headerRow[4].textContent).equal('Sale Info');
    expect(headerRow[5].textContent).equal('Quantity');

    const tableData = fixture.nativeElement.querySelectorAll('.table-responsive tbody td');
    let dataIndex = 0;
    for (const ingredient of PANTRY) {
      expect(tableData[dataIndex++].textContent).equal(ingredient.Name);
      expect(tableData[dataIndex++].textContent).equal(String(ingredient.StoreCost));
      expect(tableData[dataIndex++].textContent).equal(String(ingredient.OnSale));
      expect(tableData[dataIndex++].textContent).equal(
        ingredient.SalePrice === null ? '' : String(ingredient.SalePrice)
      );
      expect(tableData[dataIndex++].textContent).equal(
        ingredient.SaleDetails === null ? '' : ingredient.SaleDetails
      );
      expect(tableData[dataIndex++].textContent).equal(String(ingredient.Quantity));
    }
  });
});
