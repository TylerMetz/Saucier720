import { Component } from '@angular/core';
import { AuthService } from 'src/app/core/services/Auth/auth.service';
import { Router } from '@angular/router';
import { CookieService } from 'ngx-cookie-service';
import { HttpClient } from '@angular/common/http';
import { User } from 'src/app/core/interfaces/user';
import { lastValueFrom } from 'rxjs';
import { LoginRequest } from 'src/app/core/interfaces/types';

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

    async navigateHome() {
      // time delay before going home
      setTimeout(() => {
        this.router.navigate(['/Home']);
      }, 1000);
    }

    login() {
      const request: LoginRequest = {
        UserName: this.username,
        Password: this.password
      };
      this.authService.login(request).subscribe({
        next: (response: any) => {
          this.authService.loggedIn = true; // we really should change this but this is out the login button knows to switch to the logout so i am keeping for now
          console.log(response, 'navigating to home page')
          this.navigateHome();
        },
        error: (err: any) => {
          console.log(err, 'errors')
        }
      })
    }



    hideShowPass(){
      this.isText = !this.isText;
      this.isText ? this.eyeIcon = "fa-eye" : this.eyeIcon = "fa-eye-slash";
      this.isText ? this.type = "text" : this.type = "password";
    }

  }
