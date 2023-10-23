import { Component, OnInit } from '@angular/core';
import { HttpService } from './core/services/http.service';
import { CookieService } from 'ngx-cookie-service';
import { AuthService } from 'src/app/core/services/Auth/auth.service';
import { Router } from '@angular/router';
import { lastValueFrom } from 'rxjs';
import { LogoutRequest } from './core/interfaces/types';

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
    const request: LogoutRequest = {
      UserName: this.username
    };
    this.authService.logout(request).subscribe({
      next: (response: any) => {
        this.authService.loggedIn = false; // we really should change this but this is out the login button knows to switch to the logout so i am keeping for now
        console.log(response, 'user logged out')
        localStorage.removeItem('recipeNavBarButtonState');
        localStorage.removeItem('myRecipesValue');
        localStorage.removeItem('userRecipesValue');
        localStorage.removeItem('mdRecipesValue');
        if (this.router.url === '/Home'){
          window.location.reload();
        } else{
          this.router.navigate(['/Home']);
        }
      },
      error: (err: any) => {
        console.log(err, 'errors')
      }
    });
  }

  getAuthService() {
    const sessionId = this.cookieService.get('sessionID');
    this.username = sessionId.slice(0, -3);
    return this.authService;
  }

}

