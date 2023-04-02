import { ComponentFixture, TestBed, inject } from '@angular/core/testing';

import { DEALS } from 'src/app/mocks/deals.mock';
import { DealsTableComponent } from '../../../../deals/FEC/deals-table/deals-table.component';
import { DealsService } from 'src/app/core/services/deals/deals.service';

describe('DealsTableComponent', () => {
  let component: DealsTableComponent;
  let fixture: ComponentFixture<DealsTableComponent>;
  let dealsService: DealsService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [DealsTableComponent],
      imports: [],
      providers: [DealsService]
    })
      .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(DealsTableComponent);
    component = fixture.componentInstance;
    dealsService = TestBed.inject(DealsService);
  });

  it('should render table with deals',
    () => {
      component.pantry = DEALS;
      fixture.detectChanges();

      const tableRows = fixture.nativeElement.querySelectorAll('.table tbody tr');
      expect(tableRows.length).equal(DEALS.length);

      const headerRow = fixture.nativeElement.querySelectorAll('.table thead tr th');
      expect(headerRow.length).equal(4);

      expect(headerRow[0].textContent).equal('Food Name');
      expect(headerRow[1].textContent).equal('Cost');
      expect(headerRow[2].textContent).equal('Sale Price');
      expect(headerRow[3].textContent).equal('Sale Info');

      const tableData = fixture.nativeElement.querySelectorAll('.table tbody td');
      let dataIndex = 0;
      for (const ingredient of DEALS) {
        expect(tableData[dataIndex++].textContent).equal(ingredient.Name);
        expect(tableData[dataIndex++].textContent).equal(String(ingredient.StoreCost));
        expect(tableData[dataIndex++].textContent).equal(ingredient.SalePrice === null ? '' : String(ingredient.SalePrice));
        expect(tableData[dataIndex++].textContent).equal(ingredient.SaleDetails);
      }
    }
  );
});
