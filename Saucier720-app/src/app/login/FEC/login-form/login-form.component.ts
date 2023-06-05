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

    async setSessionIDAndNavigateToPantry(sessionID: string): Promise<void> {
      await new Promise<void>((resolve) => {
        this.cookieService.set('sessionID', sessionID, 7, '/', 'localhost', false, 'Lax');
        resolve();
      });
      setTimeout(() => {
        this.router.navigate(['/Pantry']);
      }, 4000);
    }

    async login() {
      const user: User = {
        FirstName: "",
        LastName: "",
        Email: "",
        UserName: this.username,
        Password: this.password,
      };
      console.log(user.UserName);
      console.log(user.Password);
      const body = { username: user.UserName, password: user.Password };
      const options = { withCredentials: true };
      try {
        const response = await lastValueFrom(this.authService.login(this.username, this.password));
        console.log('response', response)
        const sessionID = response.body.value;
        console.log("cookie set ", sessionID);
        this.setSessionIDAndNavigateToPantry(sessionID);
      } catch (error: any) {
        this.errorMessage = error.message;
      }
    }



    hideShowPass(){
      this.isText = !this.isText;
      this.isText ? this.eyeIcon = "fa-eye" : this.eyeIcon = "fa-eye-slash";
      this.isText ? this.type = "text" : this.type = "password";
    }

  }
