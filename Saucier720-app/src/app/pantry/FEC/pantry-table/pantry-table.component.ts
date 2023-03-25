import { HttpEvent, HttpEventType } from "@angular/common/http"
import { Component, OnInit } from '@angular/core';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';

@Component({
  selector: 'app-pantry-table',
  providers: [PantryService],
  templateUrl: './pantry-table.component.html',
})
export class PantryTableComponent implements OnInit {

  pantry: any;

  constructor(private pantryService: PantryService) { }

  ngOnInit(){
    this.populatePantry();
  }
  populatePantry(): void {
    this.pantryService.getPantry()
      .subscribe((event: HttpEvent<any>) => {
        switch(event.type) {
          case HttpEventType.Sent:
            console.log('Request sent!');
            break;
          case HttpEventType.ResponseHeader:
            console.log('Response header received!');
            break;
            case HttpEventType.DownloadProgress:
              const kbLoaded = Math.round(event.loaded / 1024);
              console.log(`Download in progress! ${kbLoaded}Kb loaded`);
              break;
            case HttpEventType.Response:
              console.log('Done!', event.body);
              this.pantry = event.body;
        }
      }); //doesnt call until subscribed
  }

}
