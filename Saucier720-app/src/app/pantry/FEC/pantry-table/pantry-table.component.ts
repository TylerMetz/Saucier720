import { HttpEvent, HttpEventType } from "@angular/common/http"
import { Component, OnInit } from '@angular/core';
import { PantryService } from 'src/app/core/services/pantry/pantry.service';
import { lastValueFrom } from 'rxjs';

@Component({
  selector: 'app-pantry-table',
  providers: [PantryService],
  templateUrl: './pantry-table.component.html',
})
export class PantryTableComponent implements OnInit {

  pantry: any;

  constructor(public pantryService: PantryService) { }

  async ngOnInit(){
    await this.populatePantry();
  }

  public async populatePantry(): Promise<void> {
    try {
      const event: HttpEvent<any> = await lastValueFrom(this.pantryService.getPantry());
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
          break;
      }
    } catch (error) {
      console.error(error);
    }
  }
}
