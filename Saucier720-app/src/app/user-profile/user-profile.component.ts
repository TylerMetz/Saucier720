import { Component } from '@angular/core';
import { SettingsComponent } from '../settings/settings.component';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-user-profile',
  templateUrl: './user-profile.component.html',
  styleUrls: ['./user-profile.component.scss']
})
export class UserProfileComponent{

  username: string;


  constructor(private cookieService: CookieService) {
    const sessionId = this.cookieService.get('sessionID');
    this.username = sessionId.slice(0, -3);
  }
}
