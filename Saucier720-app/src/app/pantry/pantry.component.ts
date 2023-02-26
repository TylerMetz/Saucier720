import { Component, OnInit } from '@angular/core';
import { Ingredient } from '../core/interfaces/ingredient';
import { PantryService } from '../core/services/pantry/pantry.service';

@Component({
  selector: 'app-pantry',
  templateUrl: './pantry.component.html',
  providers: [PantryService],
  styleUrls: ['./pantry.component.scss']
})

export class PantryComponent implements OnInit {
  pantry: Array<Ingredient> | undefined;

  constructor(private pantryService: PantryService) { }

  ngOnInit(){
    this.getPantry();
  }

  getPantry(): void {
    this.pantryService.getPantry()
      .subscribe(pantry => (this.pantry = pantry)); //doesnt call until subscribed
  }

  showPantry() {
    this.pantryService.getPantry()
      .subscribe((data: Array<Ingredient>) => this.pantry = { ...data });
  }
}
