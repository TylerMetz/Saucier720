import { Component } from '@angular/core';
import { AppComponent } from '../app.component';

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss']
})
export class HomePageComponent {

// used to determine which component to display
  landingPageActive: boolean;
  userDashboardActive: boolean;

  constructor(private appComponent: AppComponent) {
    this.landingPageActive = !appComponent.getAuthService().isLoggedIn();
    this.userDashboardActive = appComponent.getAuthService().isLoggedIn();
  }
}
