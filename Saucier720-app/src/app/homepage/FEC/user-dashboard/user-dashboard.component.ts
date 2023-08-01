import { Component, Output, EventEmitter, OnInit } from '@angular/core';
import { AuthService } from 'src/app/core/services/Auth/auth.service';

@Component({
    selector: 'app-user-dashboard',
    templateUrl: './user-dashboard.component.html',
    styleUrls:[ './user-dashboard.component.scss']
  })

  export class UserDashboardComponent implements OnInit{
    
    // used for button generation
    @Output() generationComplete = new EventEmitter<boolean>();

    ngOnInit() {
      setTimeout(() => {
        this.generationComplete.emit(true); // event for button generation
      }, 500); // 500ms temporarily
    }
  }