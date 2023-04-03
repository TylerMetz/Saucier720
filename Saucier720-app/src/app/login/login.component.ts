import { Component, OnInit } from '@angular/core';
import { AuthService } from '../core/services/Auth/auth.service';
import { tap } from 'rxjs/operators';
import { Router } from '@angular/router';

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
  password: string = '';
  errorMessage: string = '';

  constructor(private authService: AuthService, private router: Router) { }

  ngOnInit(): void {
  }

  login(): void {
    this.authService.login(this.username, this.password)
      .pipe(
        tap(response => {
          // Cookies are now stored in the browser
          // Redirect to the home page or another protected page
          this.router.navigate(['/']);
        })
      )
      .subscribe({
        error: error => {
          this.errorMessage = error.message;
        }
      });
  }

  hideShowPass(){
    this.isText = !this.isText;
    this.isText ? this.eyeIcon = "fa-eye" : this.eyeIcon = "fa-eye-slash";
    this.isText ? this.type = "text" : this.type = "password";
  }

}
