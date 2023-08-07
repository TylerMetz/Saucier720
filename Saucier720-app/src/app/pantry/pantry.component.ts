import { Component, Input, EventEmitter, ViewChild } from '@angular/core';
import { PantryTableComponent } from './FEC/pantry-table/pantry-table.component';

@Component({
  selector: 'app-pantry',
  templateUrl: './pantry.component.html',
  styleUrls: ['./pantry.component.scss']
})

export class PantryComponent {
  @ViewChild('pantryTable') private pantryTable!: PantryTableComponent;

  handleNewItemAdded(newItem: { name: string, quantity: number }) {
    this.pantryTable.addTempValue(newItem.name, newItem.quantity);
  }
}
