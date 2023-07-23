import { Component } from '@angular/core';
import { AuthService } from 'src/app/core/services/Auth/auth.service';

@Component({
    selector: 'app-user-dashboard',
    templateUrl: './user-dashboard.component.html',
    styleUrls:[ './user-dashboard.component.scss'],
    providers: [AuthService]
  })

  export class UserDashBoardComponent {

  }