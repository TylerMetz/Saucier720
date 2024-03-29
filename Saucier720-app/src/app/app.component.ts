import { Component, OnInit } from '@angular/core';
import { HttpService } from './core/services/http.service';
import { CookieService } from 'ngx-cookie-service';
import { AuthService } from 'src/app/core/services/Auth/auth.service';
import { Router } from '@angular/router';
import { lastValueFrom } from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {
  title: string = 'Saucier720-App';
  posts: any;
  username: string;

  constructor(private cookieService: CookieService, private authService: AuthService, private router: Router) {
    const sessionId = this.cookieService.get('sessionID');
    this.username = sessionId.slice(0, -3);
  }

  ngOnInit() {
    this.title = "Saucier720-App"
  }

  async logout() {
    try {
      const response = await lastValueFrom(this.authService.logout());
      console.log('response', response);

      // clear button/checkbox states from session
      localStorage.removeItem('recipeNavBarButtonState');
      localStorage.removeItem('myRecipesValue');
      localStorage.removeItem('userRecipesValue');
      localStorage.removeItem('mdRecipesValue');

      if (this.router.url === '/Home'){
        window.location.reload();
      } else{
        this.router.navigate(['/Home']);
      }

    } catch (error: any) {
      console.log(error.message);
    }
  }

  getAuthService() {
    const sessionId = this.cookieService.get('sessionID');
    this.username = sessionId.slice(0, -3);
    return this.authService;
  }

}

