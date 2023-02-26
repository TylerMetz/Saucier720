import { Component, OnInit } from '@angular/core';
import { Pantry } from '../core/interfaces/pantry';
import { PantryService } from '../core/services/pantry/pantry.service';

@Component({
  selector: 'app-pantry',
  templateUrl: './pantry.component.html',
  providers: [PantryService],
  styleUrls: ['./pantry.component.scss']
})

export class PantryComponent implements OnInit {
  pantry: Pantry | undefined;

  constructor(private pantryService: PantryService) { }

  ngOnInit(){
    //this.pantry = this.pantryService.getPantry();
  }

  showPantry() {
    this.pantryService.getPantry()
      .subscribe((data: Pantry) => this.pantry = { ...data });
  }
}
