import { Component, OnInit } from '@angular/core';
import { HttpService } from './core/services/http.service';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title: string = 'Saucier720-app';
  posts: any;
  constructor(private httpService: HttpService) { }

  ngOnInit() {
    this.title = "Saucier720-app"
  }
}
