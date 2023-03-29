import { HttpEvent, HttpEventType } from "@angular/common/http"
import { Component, OnInit } from '@angular/core';
import { DealsService } from 'src/app/core/services/deals/deals.service';

@Component({
  selector: 'app-deals-table',
  providers: [DealsService],
  templateUrl: 'deals-table.component.html'
})
export class DealsTableComponent implements OnInit {

  pantry: any;

  constructor(private dealsService: DealsService) { }

  ngOnInit(){
    this.populateDeals();
  }
  populateDeals(): void {
    this.dealsService.getDeals()
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
