import { Component, OnInit } from '@angular/core';
import { HttpService } from './core/services/http.service';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title: string = 'Saucier720-app';
  posts: any;
  username: string;

  constructor(private cookieService: CookieService) {
    const sessionId = this.cookieService.get('sessionID');
    this.username = sessionId.slice(0, -3);
  }

  ngOnInit() {
    this.title = "Saucier720-app"
  }
}
