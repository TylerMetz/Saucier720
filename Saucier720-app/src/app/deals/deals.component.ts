import { HttpEvent, HttpEventType } from "@angular/common/http"
import { Component, OnInit } from '@angular/core';
import { DealsService } from "../core/services/deals/deals.service";


@Component({
  selector: 'app-deals',
  templateUrl: './deals.component.html',
  providers: [DealsService],
  styleUrls: ['./deals.component.scss']
})
export class DealsComponent implements OnInit{
  pantry: any;

  constructor(private dealsService: DealsService) { }

  ngOnInit(){
    this.populatePantry();
  }

  populatePantry(): void {
    this.dealsService.getPantry()
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
