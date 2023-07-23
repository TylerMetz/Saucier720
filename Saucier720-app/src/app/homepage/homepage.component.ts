import { Component } from '@angular/core';

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss']
})

export class HomePageComponent {

    landingPageActive: boolean = true;
    userDashboardActive: boolean = false;
    
}