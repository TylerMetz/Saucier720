import { Component } from '@angular/core';
import { AuthService } from 'src/app/core/services/Auth/auth.service';

@Component({
    selector: 'app-landing-page',
    templateUrl: './landing-page.component.html',
    styleUrls:[ './landing-page.component.scss'],
    providers: [AuthService]
  })

  export class LandingPageComponent {

  }