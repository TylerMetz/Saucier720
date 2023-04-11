import { Component, OnInit } from '@angular/core';
import { AuthService } from '../core/services/Auth/auth.service';
import { tap } from 'rxjs/operators';
import { Router } from '@angular/router';
import { CookieService } from 'ngx-cookie-service';
import { HttpClient } from '@angular/common/http';
import { User } from '../core/interfaces/user';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
  
})
export class LoginComponent implements OnInit {

  type: string = "password";
  isText: boolean = false;
  eyeIcon: string = "fa-eye-slash";

  username: string = '';
  sessionID: string = ' ';
  password: string = '';
  errorMessage: string = '';

  constructor(private authService: AuthService, private router: Router, private http: HttpClient,
    private cookieService: CookieService) { 
    }


  ngOnInit(): void {
    
  }
  
  login() {
    const user: User = {
      FirstName: "",
      LastName: "",
      Email: "",
      UserName: this.username,
      Password: this.password,
    }
    console.log(user.UserName);
    console.log(user.Password);
    const body = { username: user.UserName, password: user.Password };
    const options = { withCredentials: true };
    this.authService.login(this.username, this.password)
  .subscribe({
    next: (response) => {
      const sessionID = response.headers.get('Set-Cookie');
      if (sessionID) {
        this.cookieService.set('sessionID', sessionID);
        this.router.navigate(['/']);
      } else {
        console.log(response.body); // Log the plain text response
      }
    },
    error: (error) => {
      this.errorMessage = error.message;
    },
  });
  }
    
  
  hideShowPass(){
    this.isText = !this.isText;
    this.isText ? this.eyeIcon = "fa-eye" : this.eyeIcon = "fa-eye-slash";
    this.isText ? this.type = "text" : this.type = "password";
  }

}
