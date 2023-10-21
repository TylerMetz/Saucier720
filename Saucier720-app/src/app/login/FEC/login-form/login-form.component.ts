import { Component } from '@angular/core';
import { AuthService } from 'src/app/core/services/Auth/auth.service';
import { Router } from '@angular/router';
import { CookieService } from 'ngx-cookie-service';
import { HttpClient } from '@angular/common/http';
import { User } from 'src/app/core/interfaces/user';
import { lastValueFrom } from 'rxjs';

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.scss']
})
export class LoginFormComponent {

  type: string = "password";
  isText: boolean = false;
  eyeIcon: string = "fa-eye-slash";
  isLoading: boolean = false;
  failedLogin: boolean = false;

  username: string = '';
  sessionID: string = ' ';
  password: string = '';
  errorMessage: string = '';

  constructor(
    private authService: AuthService,
    private router: Router,
    private http: HttpClient,
    private cookieService: CookieService
    ) { }

    async setSessionIDAndNavigateToHome(sessionID: string): Promise<void> {
      await new Promise<void>((resolve) => {
        this.cookieService.set('sessionID', sessionID, 7, '/', 'localhost', false, 'Lax');
        resolve();
      });
      
      // time delay before going home
      setTimeout(() => {
        this.isLoading = false;
        this.router.navigate(['/Home']);
      }, 1000);
    }

    async login() {
      this.failedLogin = false;
      this.isLoading = true;
      const user: User = {
        FirstName: "",
        LastName: "",
        Email: "",
        UserName: this.username,
        Password: this.password,
      };
      console.log(user.UserName);
      console.log(user.Password);
      const body = { UserName: user.UserName, Password: user.Password };
      const options = { withCredentials: true };
      try {
        const response = await lastValueFrom(this.authService.login(this.username, this.password));
        console.log('response', response)
        const sessionID = response.body.value;
        console.log("cookie set ", sessionID);
        this.authService.loggedIn = true;
        this.setSessionIDAndNavigateToHome(sessionID);

        // clear all button/checkbox states from session
        localStorage.removeItem('recipeNavBarButtonState');
        localStorage.removeItem('myRecipesValue');
        localStorage.removeItem('userRecipesValue');
        localStorage.removeItem('mdRecipesValue');

      } catch (error: any) {
        this.failedLogin = true;
        this.errorMessage = error.message;
        this.isLoading = false;
      }
    }



    hideShowPass(){
      this.isText = !this.isText;
      this.isText ? this.eyeIcon = "fa-eye" : this.eyeIcon = "fa-eye-slash";
      this.isText ? this.type = "text" : this.type = "password";
    }

  }
